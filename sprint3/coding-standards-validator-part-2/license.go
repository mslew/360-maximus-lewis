package main

import (
	"fmt"
	"os"
	"path/filepath"
)

//this scans the file system for LICENSE.md file
func (r root) LicenseScan() bool {
	var files []string
	dec := false
	doesItRun := true
	fmt.Println("Running LICENSE.md scanner...")
	err := filepath.Walk(string(r), func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		doesItRun = false
		panic(err)
	}
	for _, file := range files {
		if file == string(r)+"/LICENSE.md" {
			dec = true
			break
		}
	}
	if dec == true {
		fmt.Println("LICENSE.md found.")
	} else {
		fmt.Println("LICENSE.md not found.")
	}
	return doesItRun
}
