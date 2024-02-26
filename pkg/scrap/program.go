package scrap

import (
	"net/url"
	"path/filepath"
	"strings"
	"time"

	"github.com/aglide100/ebs-radio-downloader/pkg/dir"
	"github.com/aglide100/ebs-radio-downloader/pkg/logger"
	"github.com/aglide100/ebs-radio-downloader/pkg/model"
	"github.com/tidwall/gjson"
)

const (
	timeFormat = "15:04"
)

func GetCurrentProgram(currentProgramURL, imgDomain string, exclusive []string) (*model.Program, error) {
	body, err := CreateHttpReq(currentProgramURL)
	if err != nil {
		return nil, err
	}
	currentJSON := string(body)

	title := cleanStr(gjson.Get(currentJSON, "nowProgram.title").String())
	subTitle := cleanStr(gjson.Get(currentJSON, "nowProgram.subTitle").String())
	summary := gjson.Get(currentJSON, "nowProgram.summary").String()
	imgURL := gjson.Get(currentJSON, "nowProgram.programThumbnail").String()

	startAt, endAt, err := normalizeTime(gjson.Get(currentJSON, "nowProgram.start").String(), gjson.Get(currentJSON, "nowProgram.end").String())
	if err != nil {
		return nil, err
	}

	imgUrl, err := url.JoinPath(imgDomain, imgURL)
	if err != nil {
		return nil, err
	}

	program := &model.Program{
		Title:          title,
		SubTitle:       subTitle,
		StartAt:        startAt,
		EndAt:          endAt,
		Summary:        summary,
		ImgPath:        imgUrl,
	}

	if (!Contains(exclusive, title)) {
		program, err = DownloadThumb(program)
		if err != nil {
			logger.Error(err.Error())
			return nil, err
		}
	}

	return program, nil
}

func DownloadThumb(current *model.Program) (*model.Program, error) {
	path, err := dir.CreateProgramDir(current.Title, "outputs")
	if err != nil {
		return nil, err
	}

	imgPath, err := GetThumbnail(current.ImgPath, filepath.Join("outputs", current.Title))
	if err != nil {
		return nil, err
	}

	current.ImgPath = imgPath
	current.Path = path

	return current, nil 
}

func cleanStr(str string) string {
	return strings.TrimRight(strings.ReplaceAll(str, "/", ""), ". ")
}

func normalizeTime(startAt, endAt string) (time.Time, time.Time, error) {
	now := time.Now().Local()
	var startAtParsed, endAtParsed time.Time

	if len(startAt) <= 1 || len(endAt) <= 1 {
		return startAtParsed, endAtParsed, nil
	}

	startAtParsed, err := time.Parse(timeFormat, fixTime(startAt))
	if err != nil {
		return startAtParsed, endAtParsed, err
	}
	startAtParsed = time.Date(now.Year(), now.Month(), now.Day(), startAtParsed.Hour(), startAtParsed.Minute(), 0, 0, now.Location())

	endAtParsed, err = time.Parse(timeFormat, fixTime(endAt))
	if err != nil {
		return startAtParsed, endAtParsed, err
	}
	endAtParsed = time.Date(now.Year(), now.Month(), now.Day(), endAtParsed.Hour(), endAtParsed.Minute(), 0, 0, now.Location())

	return startAtParsed, endAtParsed, nil
}

func fixTime(timeStr string) string {
	timeStr = strings.ReplaceAll(timeStr, "24:", "00:")
	timeStr = strings.ReplaceAll(timeStr, "25:", "01:")
	timeStr = strings.ReplaceAll(timeStr, "26:", "02:")
	timeStr = strings.ReplaceAll(timeStr, "27:", "03:")
	timeStr = strings.ReplaceAll(timeStr, "28:", "04:")
	
	return timeStr
}
