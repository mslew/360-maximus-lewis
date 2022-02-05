package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func (r root) CharacterScan() {
	var files []string
	var text []string
	var length int
	lineCount := 0
	dec := true
	fmt.Println("Running Character scanner...")
	err := filepath.Walk(string(r), func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, fileInFiles := range files {
		fileExtension := filepath.Ext(fileInFiles)
		if fileExtension == ".go" {
			fmt.Println("Opening " + fileInFiles)
			file, err := os.Open(string(r))
			if err != nil {
				log.Fatal(err)
			}
			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanLines)
			for scanner.Scan() {
				text = append(text, scanner.Text())
			}
			file.Close()
			for _, line := range text {
				lineCount++
				length = len(line)
				if length > 99 {
					fmt.Println("Count exceeds 100 characters on line" + string(rune(lineCount)))
					dec = false
				} else {
					continue
				}
			}
		}
		lineCount = 0
	}
	if dec == false {
		fmt.Println("Character count error")
	} else {
		fmt.Println("No character count errors")
	}
}
