package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func (r root) ReadMeScan() {
	var files []string
	dec := false
	fmt.Println("Running README.md scanner...")
	err := filepath.Walk(string(r), func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
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
}
