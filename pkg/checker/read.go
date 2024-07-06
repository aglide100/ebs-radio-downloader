package checker

import (
	"fmt"
	"os"
	"path/filepath"
)

// json read
func ReadDirInJsons(programPath string) error {
	entries, err := os.ReadDir(programPath)
	if err != nil {
		return err
	}

	for _, e := range entries {
		if filepath.Ext(e.Name()) == ".json" {
			fmt.Println(e.Name())
		}
	}

	return nil
}
