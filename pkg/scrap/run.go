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

	isExist := false

	prev := current
	if dir.SubDirIsExist(current.Path, current.SubTitle) {
		logger.Info("subtitle is exist", zap.Any("current", current))
		isExist = true
	}

	idx := 0
	defer ticker.Stop()
	for range ticker.C {
		remain := time.Until(prev.EndAt)

		idx++
		isChange := false

		if (remain <= time.Second * 30) {
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
				isExist = true
			} else {
				isExist = false
			}

			if prev.Title != current.Title || prev.SubTitle != current.SubTitle {
				isChange = true
				if dir.SubDirIsExist(prev.Path, prev.SubTitle) {
					logger.Info("already exist", zap.Any("prev", prev))
					prev = current 
					continue
				}

				if (prev.Title != "" && !Contains(exclusive, prev.Title)) {
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
		}
		
		if (Contains(exclusive, current.Title)) {
			waitUntilDone(current)
			continue
		}

		if !all && !Contains(wanted, current.Title) {
			waitUntilDone(current)
			continue
		}

		if (!isExist) {
			err = DownloadChunk(m3u8URL, baseURL, isChange, current)
			if err != nil {
				return err
			}
		} else {
			logger.Debug("skip download")
		}
	}

	return nil
}
