package checker

import (
	"fmt"
	"strings"
	"time"

	"github.com/aglide100/ebs-radio-downloader/pkg/model"
)

type BroadCast struct {
	Program     model.Program `json:Program`
	Episodes    []Episode     `json:Episodes`
	FileStartAt time.Time     `json:FileStartAt`
	FileEndAt   time.Time     `json:FileStartAt`
}

type Episode struct {
	Title string
	ID    string
}

type Configure struct {
}

type Constraint func() error

// Validator ...
type Validator struct {
	constraints []Constraint
}

// 뭐가 좀 필요한거 같은데
// TODO; add type check
func (validator *Validator) Validate() error {
	errors := []string{}

	for _, constraint := range validator.constraints {
		err := constraint()

		if err != nil {
			errors = append(errors, err.Error())
		}
	}

	if len(errors) == 0 {
		return nil
	}

	return fmt.Errorf(strings.Join(errors, ", "))
}
