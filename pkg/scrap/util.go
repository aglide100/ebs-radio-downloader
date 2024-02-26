package scrap

import (
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/aglide100/ebs-radio-downloader/pkg/logger"
	"github.com/aglide100/ebs-radio-downloader/pkg/model"
	"github.com/aglide100/ebs-radio-downloader/pkg/word"
	"go.uber.org/zap"
)

var (
	distanceThreshold = 3
)

func CreateHttpRes(url string) (*http.Response, error) {
	res, err := http.Get(url)
  	
	if err != nil {
		return nil, err
  	}
  	defer res.Body.Close()

  	if res.StatusCode != 200 {
		return nil, err
	}

	return res, nil
}

func CreateHttpReq(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil) 
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, HandleHttpStatusErr(res)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}


func HandleHttpStatusErr(res *http.Response) (error) {
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return  err
	}

	return errors.New("status code error! "+ string(data) )
}

func Contains(s []string, substr string) bool {
    for _, v := range s {
		distance := word.EditDistance(v, substr)

		if v == substr {
            return true
        }

        if distance <= distanceThreshold {
			logger.Info("should be this one,", zap.Any("s1", v), zap.Any("s2", substr))
			return true
		}
    }

    return false
}

func waitUntilDone(current *model.Program) {
	if waitDuration := time.Until(current.EndAt); waitDuration > 0 {
		logger.Info("waiting until program done", zap.Any("duration", waitDuration), zap.Any("current", current))
		time.Sleep(waitDuration)
	}
}