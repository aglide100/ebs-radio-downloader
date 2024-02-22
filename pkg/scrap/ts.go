package scrap

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/aglide100/ebs-radio-downloader/pkg/cli"
	"github.com/aglide100/ebs-radio-downloader/pkg/dir"
	"github.com/aglide100/ebs-radio-downloader/pkg/logger"
	"github.com/aglide100/ebs-radio-downloader/pkg/model"
	"go.uber.org/zap"
)

func DownloadTSFile(url, fileName string) error {
	if _, err := os.Stat(fileName); !os.IsNotExist(err) {
		// fmt.Println("File is exist:", fileName)
		return nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("downloading...", fileName)
	return nil
}

func GetTsList(host, body string) (tsList []model.TsInfo) {
	lines := strings.Split(body, "\n")
	var ts model.TsInfo
	var duration float64
	
	for _, line := range lines {
		/*
		#EXTINF:9.984,
		media_4716901.ts
		*/
		if strings.HasPrefix(line, "#EXTINF:") {
			durationStr := strings.TrimPrefix(line, "#EXTINF:")
			duration, _ = strconv.ParseFloat(strings.Split(durationStr, ",")[0], 64)
		} else if !strings.HasPrefix(line, "#") && line != "" && strings.Contains(line, ".ts") {
			ts = model.TsInfo{
				Name:     line,
				Url:      host + line,
				Duration: duration,
			}
			tsList = append(tsList, ts)
			
			duration = 0
		}
	}

	return
}

func CreateSubDir(path, title string) (string, error) {
	dirCnt, err := dir.CountDirs(path)
	if err != nil {
		logger.Error("can't count folders", zap.Any("path", path), zap.Error(err))
		return "", err
	}

	dirNum := fmt.Sprintf("%02d", dirCnt)

	subPath := filepath.Join(path, fmt.Sprintf("%s.%s", dirNum, title)) 

	err = dir.CreatePath(subPath)
	if err != nil {
		logger.Error("can't create path", zap.Any("path", subPath))
		return "", err
	} 

	return subPath, nil
}

func CombinedTS(path, title string) (string, string, error) {
	newTitle := dir.PreProcessing(title)

	basePath := strings.TrimRight(path, "/") + "/"

	subPath, err := CreateSubDir(basePath, newTitle)
	if err != nil {
		logger.Error("can't create sub folders", zap.Any("path", basePath), zap.Error(err))
		return "","", err
	}

	err = cli.RunCombineTs(basePath, subPath)
	if err != nil {
		logger.Error("can't RunCombineTS", zap.Any("basePath", basePath), zap.Any("subPath", dir.AddEscapePath(subPath)))
		return "","", err
	}
	
	files, err := os.ReadDir(strings.ReplaceAll(basePath, "\\", ""))
	if err != nil {
		logger.Error("can't open path", zap.Any("path", basePath))
		return "","", err
	}
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".ts") && file.Name() != dir.AddEscapePath(newTitle)+".ts" {
			err := os.Remove(strings.ReplaceAll(basePath, "\\", "") + file.Name())
			if err != nil {
				return "","", err
			}
		}
	}

	return subPath, "/all.ts", nil
}	

