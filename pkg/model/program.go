package model

import "time"

type Program struct {
	Title string
	SubTitle  string
	StartAt time.Time
	EndAt time.Time
	Path string
}