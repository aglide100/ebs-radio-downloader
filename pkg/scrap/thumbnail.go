package scrap

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func GetThumbnail(imgURL, path string) (string, error) {	
	imgName := filepath.Base(imgURL)
	imgPath := filepath.Join(path, imgName)
	
	if _, err := os.Stat(imgPath); !os.IsNotExist(err) {
		// fmt.Println("File is exist:", imgName)
		return imgPath, nil
	}

	resp, err := http.Get(imgURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()


	file, err := os.Create(imgPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", err
	}

	return imgPath, nil
}