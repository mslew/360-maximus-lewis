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
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return dir
}

func welcome() {
	fmt.Println("*********************")
	fmt.Println("      VALIDATOR      ")
	fmt.Println("*********************")
}

func footer() {
	fmt.Println("*********************")
	fmt.Println(" Thank you for using ")
	fmt.Println("*********************")
}

func codeBreak() {
	fmt.Println("***********************************************************************")
}

func main() {
	root := root(RootDir())
	var v Validate
	welcome()
	v = root
	v.ReadMeScan()
	codeBreak()
	v.LicenseScan()
	codeBreak()
	v.CharacterScan()
	codeBreak()
	v.TabScan()
	footer()
}
