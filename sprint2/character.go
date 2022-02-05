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
	var count int
	var lineCount int
	var chars []rune
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
			file, err := os.Open(string(r))
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			reader := bufio.NewReader(file)
			for {
				lineCount++
				line, err := reader.ReadString('\n')
				if err != nil {
					break
				}
				chars = []rune(line)
				for i := 0; i < len(chars); i++ {
					count++
				}
				if count > 100 {
					fmt.Println("Greater than 100 characters on line" + string(lineCount))
					dec = false
				} else {
					continue
				}
				count = 0
			}
		}
	}
	if dec == false {
		fmt.Println("Character count exceeds 100 characters on specific lines.")
	} else {
		fmt.Println("No character count errors.")
	}
}
