package scrap

import (
	"flag"
	"path/filepath"
	"time"

	"github.com/aglide100/ebs-radio-downloader/pkg/cli"
	"github.com/aglide100/ebs-radio-downloader/pkg/dir"
	"github.com/aglide100/ebs-radio-downloader/pkg/logger"
	"github.com/aglide100/ebs-radio-downloader/pkg/model"
	"go.uber.org/zap"
)

var (
	outputs = flag.String("outputs", "outputs", "output dir")
	baseURL = "https://ebsonair.ebs.co.kr/cloud/iradio/"
	currentProgramURL = "https://www.ebs.co.kr/onair/cururentOnair.json?channelCd=IRADIO"
	m3u8URL = "https://ebsonair.ebs.co.kr/cloud/iradio/chunklist.m3u8"
)

func Init() error {
	flag.Parse()

	err := dir.CreatePath(*outputs)
	if err != nil  {
		return err
	}

	return nil
}

func Run(wanted, exclusive []string) error {
	duration, _ := time.ParseDuration("10s")
	ticker := time.NewTicker(duration)

	all := false
	if (len(wanted) == 1 && wanted[0] == "*") {
		all = true
	}

	current, err := GetCurrentProgram()
	if err != nil {
		return err
	}

	prev := current

	idx := 0
	defer ticker.Stop()
	for range ticker.C {
		remain := time.Until(prev.EndAt)

		idx++
		isChange := false
		if (remain <= time.Second * 30) {
			// logger.Info("remain", zap.Any("remain", remain), zap.Any("prev", prev.EndAt))
			idx = 6
		}
		
		if (idx == 6) {
			idx = 0
			current, err = GetCurrentProgram()
			if err != nil {
				logger.Error("Can't get current program")
				return err
			}

			if dir.SubDirIsExist(current.Path, current.SubTitle) {
				prev = current
				logger.Info("subtitle is exist", zap.Any("current", current))
				continue
			}

			if prev.Title != current.Title || prev.SubTitle != current.SubTitle {
				isChange = true
				if (prev.Title != "" && !Contains(exclusive, prev.Title)) {
					logger.Info("prev", zap.Any("prev", prev))
			
					go func(target *model.Program) {
						logger.Info("prev", zap.Any("target", target))
						path, filename, err := CombinedTS(target.Path, target.SubTitle)
						if err != nil {
							logger.Error(err.Error())
							return
						}
		
						err = cli.RunConvertTsToMP3(filepath.Join(path, filename),  filepath.Join(path, target.SubTitle+".mp3" ))
						if err != nil {
							logger.Error(err.Error())
							return
						}
					}(prev)
				}
			}
			
			prev = current 

			if current.Title == "" {
				now := time.Now()
				current.EndAt = time.Date(now.Year(), now.Month(), now.Day(), 4, 55, 0, 0, now.Location())
				current.Title = " "
				logger.Info("current broadcast is done, until wait ", zap.Any("time", now))
			}

			// logger.Info("program", zap.Any("current", current))		
		}

		
		if (Contains(exclusive, current.Title)) {
			waitUntilDone(current)
			continue
		}

		if !all && !Contains(wanted, current.Title) {
			waitUntilDone(current)
			continue
		}

		body, err := CreateHttpReq(m3u8URL)
		if err != nil {
			return err
		}
		
		list := GetTsList(baseURL, string(body))
		
		if isChange {
			err = DownloadTSFile(list[len(list)-1].Url, filepath.Join(current.Path, list[len(list)-1].Name))
			if err != nil {
				logger.Error("Can't download ts file")
				return err
			}
		} else {
			for _, val := range list {
				err = DownloadTSFile(val.Url, filepath.Join(current.Path, val.Name))
				if err != nil {
					logger.Error("Can't download ts file")
					return err
				}

				if (val.Duration > 10) {
					time.Sleep(time.Second * time.Duration((val.Duration - 10)))
				}
			}
		}
	}

	return nil
}
