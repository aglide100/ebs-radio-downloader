package checker

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const fileName = "log.json"

func WriteConfigure(broadcast *BroadCast, path string) error {
	doc, err := json.Marshal(broadcast)
	if err != nil {
		return err
	}

	resultPath := filepath.Join(broadcast.Program.Path, path, fileName)

	err = os.WriteFile(path, doc, os.FileMode(0644))
	if err != nil {
		return err
	}

	return nil
}

func ReadConfigure(path string) error {
	_, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return nil
}
