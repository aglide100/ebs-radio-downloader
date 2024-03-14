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
	imgBaseURL = "https://static.ebs.co.kr/images"
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

	currentProgram, err := GetCurrentProgram(currentProgramURL, imgBaseURL, exclusive)
	if err != nil {
		return err
	}

	isExist := false

	// TODO; 중복 관련해서 로직을 다시 작성해야함
	// 재방송의 경우 동일날에도 여러번 반복되지만 그전날 방송또한 재방을 함
	// 게다가 시간을 토대로 자르기때문에 앞뒤 부분에 일부 내용이 들어가거나 짤릴 수 있음
	prevProgram := currentProgram
	if dir.IsDownloaded(currentProgram.Path, currentProgram.SubTitle) {
		logger.Info("subtitle is exist", zap.Any("current", currentProgram))
		isExist = true
	}

	idx := 0
	defer ticker.Stop()
	for range ticker.C {
		remain := time.Until(prevProgram.EndAt)

		idx++
		isChange := false

		if (remain <= time.Second * 30) {
			idx = 6
		}
		
		if (idx == 6) {
			idx = 0
			currentProgram, err = GetCurrentProgram(currentProgramURL, imgBaseURL, exclusive)
			if err != nil {
				logger.Error("Can't get current program")
				return err
			}

			if !Contains(exclusive, currentProgram.Title) && dir.IsDownloaded(currentProgram.Path, currentProgram.SubTitle) {
				prevProgram = currentProgram
				logger.Info("already exist", zap.Any("title", currentProgram.Title), zap.Any("subTitle", currentProgram.SubTitle))
				isExist = true
			} else {
				isExist = false
			}

			if prevProgram.Title != currentProgram.Title || prevProgram.SubTitle != currentProgram.SubTitle {
				isChange = true
				if !Contains(exclusive, prevProgram.Title) && dir.IsDownloaded(prevProgram.Path, prevProgram.SubTitle) {
					logger.Info("already exist", zap.Any("prevProgram title", prevProgram.Title), zap.Any("prevProgram subtitle", prevProgram.SubTitle))
					prevProgram = currentProgram 
					continue
				}

				if (prevProgram.Title != "" && !Contains(exclusive, prevProgram.Title)) {
					go func(target *model.Program) {
						// logger.Info("prevProgram", zap.Any("target", target))
						subDirPath, filename, err := CombinedTS(target.Path, target.SubTitle)
						if err != nil {
							logger.Error(err.Error())
							return
						}
		
						err = cli.RunConvertTsToMP3(filepath.Join(subDirPath, filename), filepath.Join(subDirPath, target.SubTitle+".mp3"), target)
						if err != nil {
							logger.Error(err.Error())
							return
						}
					}(prevProgram)
				}
			}
			
			prevProgram = currentProgram 

			if currentProgram.Title == "" {
				now := time.Now()
				currentProgram.EndAt = time.Date(now.Year(), now.Month(), now.Day(), 4, 55, 0, 0, now.Location())
				// currentProgram.Title = " "
				logger.Info("current broadcast is done, until wait ", zap.Any("time", now))
			}	
		}
		
		if (Contains(exclusive, currentProgram.Title)) {
			waitUntilDone(currentProgram)
			continue
		}

		if !all && !Contains(wanted, currentProgram.Title) {
			waitUntilDone(currentProgram)
			continue
		}

		if (!isExist) {
			err = DownloadChunk(m3u8URL, baseURL, isChange, currentProgram)
			if err != nil {
				return err
			}
		} else {
			logger.Debug("skip download")
		}
	}

	return nil
}
