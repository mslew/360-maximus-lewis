package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

//this will scan files in the directory
//finds .go files and opens them
//goes line by line and counts characters for each line
//if the characters exceed 100 it will report to user.
func (r root) CharacterScan() bool {
	var files []string
	var text []string
	var length int
	lineCount := 0
	dec := true
	doesItRun := true
	fmt.Println("Running Character scanner...")
	err := filepath.Walk(string(r), func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		doesItRun = false
		panic(err)
	}
	for _, fileInFiles := range files {
		fileExtension := filepath.Ext(fileInFiles)
		if fileExtension == ".go" {
			fmt.Println("Opening " + fileInFiles)
			file, err := os.Open(string(r))
			if err != nil {
				doesItRun = false
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
	return doesItRun
}
