package main

import (
	"fmt"
	"os"
)

type Validate interface {
	ReadMeScan()
	LicenseScan()
	CharacterScan()
	TabScan()
}

type root string

func RootDir() string {
	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return mydir
}

func welcome() {
	fmt.Println("*********************")
	fmt.Println("      VALIDATOR      ")
	fmt.Println("*********************")
}

func main() {
	root := root(RootDir())
	var v Validate
	welcome()
	v = root
	v.ReadMeScan()
	v.LicenseScan()
	v.CharacterScan()
	v.TabScan()
}
