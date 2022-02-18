package main

import (
	"fmt"
	"os"
	"path/filepath"
)

//scans the file system for a README.md file
func (r root) ReadMeScan() bool {
	var files []string
	dec := false
	doesItRun := true
	fmt.Println("Running README.md scanner...")
	err := filepath.Walk(string(r), func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		doesItRun = false
		panic(err)
	}
	for _, file := range files {
		if file == string(r)+"/README.md" {
			dec = true
			break
		}
	}
	if dec == true {
		fmt.Println("README.md found.")
	} else {
		fmt.Println("README.md not found.")
	}
	return doesItRun
}
