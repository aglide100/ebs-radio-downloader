package cli

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/aglide100/ebs-radio-downloader/pkg/logger"
)

func RunConvertTsToMP3(inputPath, outputPath string) error {
	absInputPath, err := filepath.Abs(inputPath)
    if err != nil {
        logger.Error(err.Error())
        return err
    }

    absOutputPath, err := filepath.Abs(outputPath)
    if err != nil {
        logger.Error(err.Error())
        return err
    }

    cmd := exec.Command("ffmpeg", "-i", absInputPath, "-vn", "-acodec", "libmp3lame", absOutputPath)
    
	var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
	
    err = cmd.Run()
    if err != nil {
		log.Printf("Command failed with error: %v", err)
        log.Printf("Command stderr: %s", stderr.String())
		log.Printf("cmd: %s", cmd.String())

		logger.Error(err.Error())
        return err
    }

	err = os.Remove(inputPath)
	if err != nil {
		logger.Error(err.Error())
        return err
    }
    
    return nil
}