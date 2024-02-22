package scrap

import (
	"strings"
	"time"

	"github.com/aglide100/ebs-radio-downloader/pkg/dir"
	"github.com/aglide100/ebs-radio-downloader/pkg/model"
	"github.com/tidwall/gjson"
)

func GetCurrentProgram() (*model.Program, error) {
	body, err := CreateHttpReq(currentProgramURL)
	if err != nil {
		return nil, err
	}
	currentJSON := string(body)

	timeFormat := "15:04"

	title := dir.PreProcessing(strings.ReplaceAll(gjson.Get(currentJSON, "nowProgram.title").String(), "/", ""))
	subTitle := dir.PreProcessing(strings.ReplaceAll(gjson.Get(currentJSON, "nowProgram.subTitle").String(), "/", ""))

	startAt := strings.ReplaceAll(gjson.Get(currentJSON, "nowProgram.start").String(), "24:", "00:")
	endAt := strings.ReplaceAll(gjson.Get(currentJSON, "nowProgram.end").String(), "24:", "00:")
	startAt = strings.ReplaceAll(startAt, "25:", "01:")
	endAt = strings.ReplaceAll(endAt, "25:", "01:")

	path, err := dir.CreateProgramDir(title, "outputs")
	if err != nil  {
		return nil, err
	}

	now := time.Now().Local()
	var startAtParsed, endAtParsed time.Time

	if len(startAt) <= 1 || len(endAt) <=1 {
		//pass
	} else {
		startAtParsed, err = time.Parse(timeFormat, startAt)
		if err != nil {
			return nil, err
		}
		startAtParsed = time.Date(now.Year(), now.Month(), now.Day(), startAtParsed.Hour(), startAtParsed.Minute(), 0, 0, now.Location())

		endAtParsed, err = time.Parse(timeFormat, endAt)
		if err != nil {
			return nil, err
		}
		endAtParsed = time.Date(now.Year(), now.Month(), now.Day(), endAtParsed.Hour(), endAtParsed.Minute(), 0, 0, now.Location())
	}

	program := &model.Program{
		Title: title,
		SubTitle: subTitle,
		Path: path,
		StartAt: startAtParsed,
		EndAt: endAtParsed,
	}

	return program, nil
}