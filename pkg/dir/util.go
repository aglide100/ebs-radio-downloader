package dir

import (
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/aglide100/ebs-radio-downloader/pkg/logger"
	"github.com/aglide100/ebs-radio-downloader/pkg/word"
	"go.uber.org/zap"
)

func AddEscapePath(path string) string {
    escapeChars := []rune{' ', '(', ')', '!', '&', '<', '>', '[', ']', '{', '}', ';', '*', '^', '$', '~', '\'', '"'}

    res := PreProcessing(path)

    for _, char := range escapeChars {
        res = strings.ReplaceAll(res, string(char), "\\"+string(char))
    }

    return res
}

func PreProcessing(text string) string {
    res := strings.TrimLeft(text, ".")
	res = strings.ReplaceAll(res, ":", "")
	res = strings.ReplaceAll(res, "|", "")
	res = strings.ReplaceAll(res, "\"", "")
    res = strings.ReplaceAll(res, "\t", "")
    res = strings.ReplaceAll(res, "\n", "")
	res = strings.ReplaceAll(res, "\\", "")

    for strings.Contains(res, "  ") {
        res = strings.ReplaceAll(res, "  ", " ")
    }

    res = strings.TrimSpace(res)

	return res
}

func CountDirs(path string) (int, error) {
    var folderCount int

    err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if info.IsDir() {
            folderCount++
        }
        return nil
    })

    if err != nil {
        return 0, err
    }

    return folderCount, nil
}

func SortDateDir(path string) ([]fs.DirEntry, error) {
    entries, err := os.ReadDir(path)
    if err != nil {
        logger.Error(err.Error())
        return nil, err
    }
    
    sort.Slice(entries, func(i, j int) bool {
        return extractDate(entries[i].Name()) > extractDate(entries[j].Name())
    })

    return entries, nil 
}

func IsDownloaded(path, subtitle string) (bool) {
    // logger.Info("path", zap.Any("path", path))

    found := false

    entries, err := os.ReadDir(path)
    if err != nil {
        logger.Info("", zap.Any("path", path), zap.Any("subtitle", subtitle))
        logger.Error(err.Error())
        return false
    }

    for _, e := range entries {
        if strings.Contains(PreProcessing(e.Name()), PreProcessing(subtitle)) {
            distance := word.EditDistance(e.Name(), subtitle)

            if distance <= 5 {
                subdir := filepath.Join(path, e.Name())
                
                err := filepath.WalkDir(subdir, func(subpath string, d fs.DirEntry, err error) error {
                    if err != nil {
                        logger.Error(err.Error())
                        return err
                    }

                    if strings.Contains(d.Name(), ".mp3") {
                        found = true
                        logger.Info("found", zap.Any("filename", d.Name()))
                        return filepath.SkipDir
                    }

                    return nil
                })

                if err != nil {
                    logger.Error(err.Error())
                    break
                }

                break
            }
        }
    }

    return found
} 

func extractDate(name string) int {
    dateStr := strings.ReplaceAll(name, "-", "")

    res, err := strconv.Atoi(dateStr)
    if err != nil {
        return 0
    }

    return res
}
