package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//this function scans the folder for files.
//this will open files with .go extensions.
//line by line and will see if there are tabs or not
func (r root) TabScan() bool {
	var files []string
	dec := true
	var badFile string
	doesItRun := true
	fmt.Println("Running tab scanner...")
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
			defer file.Close()
			reader := bufio.NewReader(file)
			for {
				line, err := reader.ReadString('\n')
				if err != nil {
					break
				}
				if strings.HasPrefix(line, "	") {
					badFile = string(fileInFiles)
					dec = false
					break
				}
			}
		}
	}
	if dec == false {
		fmt.Println("Tab Error at " + badFile)
	} else {
		fmt.Println("No tab errors")
	}
	return doesItRun
}
