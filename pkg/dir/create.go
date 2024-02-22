package dir

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func CreatePath(path string) error {
	err := os.MkdirAll(strings.ReplaceAll(path, "\\", ""), 0755)
	if err != nil && !os.IsExist(err) {
		return err
	}

	return nil
}

func CreateTodayDir(base string) (string, error) {
	now := time.Now()

	today := now.Format("2006-01-02")

	dir := fmt.Sprintf("/%s", today)

	err := CreatePath(base + dir)
	if err != nil {
		return "", err
	}

	return base + dir, nil
}

func CreateProgramDir(title, output string) (string, error) {
	path := filepath.Join(output, title)

	err := CreatePath(path)
	if err != nil {
		return "", err
	}

	pathWithDate, err := CreateTodayDir(path)
	if err != nil {
		return "", err
	}

	return pathWithDate, nil
}