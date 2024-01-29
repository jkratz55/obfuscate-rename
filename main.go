package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	flag "github.com/spf13/pflag"
)

var (
	logger *log.Logger
)

func main() {
	var (
		directory = flag.StringP("directory", "d", "", "Directory to search for files and rename")
		recursive = flag.BoolP("recursive", "r", false, "Recursively search for files in the directory")
		extension = flag.String("extension", "", "File extension to search for")
	)
	flag.Parse()

	logger = log.New(os.Stdout, "[obfuscate-rename] ", log.LstdFlags)

	if strings.TrimSpace(*directory) == "" {
		logger.Fatal("directory is required")
	}

	err := renameFiles(*directory, *recursive, *extension)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Println("Done!")
}

func renameFiles(path string, recursive bool, extension string) error {
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, file := range files {
		currentPath := filepath.Join(path, file.Name())
		if file.IsDir() && recursive {
			err := renameFiles(currentPath, recursive, extension)
			if err != nil {
				return err
			}
		}

		if !file.IsDir() {
			if extension == "" || filepath.Ext(currentPath) == extension {
				newPath := strings.ReplaceAll(currentPath, filepath.Base(currentPath), uuid.New().String()) + filepath.Ext(currentPath)
				err := os.Rename(currentPath, newPath)
				if err != nil {
					logger.Printf("failed to rename file %s to %s: %v\n", currentPath, newPath, err)
				} else {
					logger.Printf("renamed file %s to %s\n", currentPath, newPath)
				}
			}
		}
	}

	return nil
}
