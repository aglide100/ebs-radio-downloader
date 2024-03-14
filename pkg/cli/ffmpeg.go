package cli

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/aglide100/ebs-radio-downloader/pkg/logger"
	"github.com/aglide100/ebs-radio-downloader/pkg/model"
)

func RunConvertTsToMP3(inputPath, outputPath string, target *model.Program) error {
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

    err = AddMetadata(absOutputPath, target)
	if err != nil {
		logger.Error(err.Error())
        return err
    }
    
    return nil
}

func AddMetadata(absOutputPath string, target *model.Program) error {
    titleMetadata := "title=\"" + target.SubTitle + "\""
    artistMetadata := "artist=\"" + target.Title + "\""
    albumMetadata := "album=\"" + target.Title + "\""

    dirPath := filepath.Dir(absOutputPath)
    // logger.Info(dirPath)
    cmd := exec.Command("ffmpeg", "-i", absOutputPath, "-i", target.ImgPath, "-map", "0:0", "-map", "1:0", "-c", "copy", "-id3v2_version", "3", "-metadata", titleMetadata, "-metadata", artistMetadata, "-metadata", albumMetadata, dirPath +"/out.mp3")
    
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

    err = os.Remove(absOutputPath)
	if err != nil {
		logger.Error(err.Error())
        return err
    }
    
    err = os.Rename(dirPath + "/out.mp3", absOutputPath)
    if err != nil {
		logger.Error(err.Error())
        return err
    }

    return nil
}