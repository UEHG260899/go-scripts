package main

import (
	"fmt"
	"log"
	"io"
	"path"
	"os"
)


func main() {
	directoriesForFiles := []string{"Pdfs", "Images"}
	home := os.Getenv("HOME")
	err := os.Chdir(home + "/Descargas")

	if err != nil {
		log.Fatal(err)
	}

	for _, directory := range directoriesForFiles {
		err := os.Mkdir(directory, 0775)
		if err != nil && !os.IsExist(err) {
			log.Fatal(err)
		}
	}


	files, err := os.ReadDir(".")
	
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fileName := file.Name()

		switch path.Ext(fileName) {
		case ".pdf":
			MoveFile(fileName, "Pdfs/" + fileName)
		case ".jpg", ".png", ".jpeg":
			MoveFile(fileName, "Images/" + fileName)
		}

	}
	
}

func MoveFile(sourcePath, destinationPath string) error {
	inputFile, err := os.Open(sourcePath)

	if err != nil {
		return fmt.Errorf("Couldn´t open source file")
	}
	defer inputFile.Close()

	outputFile, err := os.Create(destinationPath)

	if err != nil {
		return fmt.Errorf("Couldn´t open destination path")
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, inputFile)

	if err != nil {
		return fmt.Errorf("Couldn´t copy to destination")
	}

	// If running windows
	// inputFile.Close()

	err = os.Remove(sourcePath)

	if err != nil {
		return fmt.Errorf("Couldn´t remove source file")
	}
	
	return nil
}
