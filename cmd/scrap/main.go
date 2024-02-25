package main

import (
	"flag"
	"strings"

	"github.com/aglide100/ebs-radio-downloader/pkg/logger"
	"github.com/aglide100/ebs-radio-downloader/pkg/scrap"
	"go.uber.org/zap"
)

var (
	// typo program titles
	// ex) EASY ENGLISH(영어회화 레벨2),EASY ENGLISH(영어회화 레벨1),EASY ENGLISH(영어회화 레벨0)
	wanted = flag.String("wanted", "*", "")
	exclusive = flag.String("exclusive", "라디오 캠페인", "")
)

func main() {
	err := scrap.Init()
	if err != nil {
		logger.Error(err.Error())
		return
	}

	wanted_list := strings.Split(*wanted, ",")
	exclusive_list := strings.Split(*exclusive, ",")
	exclusive_list = append(exclusive_list, " ")

	logger.Info("", zap.Any("wanted_list", wanted_list))
	logger.Info("", zap.Any("exclusive", exclusive))

	err = scrap.Run(wanted_list, exclusive_list)
	if err != nil {
		logger.Error(err.Error())
		return
	}
}
