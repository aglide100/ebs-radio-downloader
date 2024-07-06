package checker

import (
	"path/filepath"
	"strconv"
	"time"
)

func TimeToPath(time time.Time) (string, error) {
	year := time.Year()
	month := int(time.Month())

	path := filepath.Join(strconv.Itoa(year), strconv.Itoa(month))

	return path, nil
}
