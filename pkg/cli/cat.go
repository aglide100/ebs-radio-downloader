package cli

import (
	"bytes"
	"log"
	"os/exec"
	"strings"

	"github.com/aglide100/ebs-radio-downloader/pkg/dir"
	"github.com/aglide100/ebs-radio-downloader/pkg/logger"
)

func RunCombineTs(inputPath, outputPath string) error {
	// yap, this code should be works only on linux...
	cmd := exec.Command("bash", "-c", "cat " + strings.TrimRight(dir.AddEscapePath(inputPath), "/") + "/*.ts" + " > " + dir.AddEscapePath(outputPath) + "/all.ts")
	
	var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
        log.Printf("Command failed with error: %v", err)
        log.Printf("Command stderr: %s", stderr.String())
        log.Printf("cmd: %s", cmd.String())

		logger.Error(err.Error())
		return err
	}

	return nil
}