package main

import (
	"fmt"
	"os"
)

//Interface for validate
type Validate interface {
	ReadMeScan() bool
	LicenseScan() bool
	CharacterScan() bool
	TabScan() bool
}

//root string
type root string

//find the root where the program was executed
func RootDir() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return dir
}

//welcome header
func welcome() {
	fmt.Println("*********************")
	fmt.Println("      VALIDATOR      ")
	fmt.Println("*********************")
}

//footer
func footer() {
	fmt.Println("*********************")
	fmt.Println(" Thank you for using ")
	fmt.Println("*********************")
}

//break used to breakup the different processes
func codeBreak() {
	fmt.Println("***********************************************************************")
}

//main function to execute the program
func main() {
	var v Validate
	r := root(RootDir())
	v = r
	welcome()
	v.ReadMeScan()
	codeBreak()
	v.LicenseScan()
	codeBreak()
	v.CharacterScan()
	codeBreak()
	v.TabScan()
	footer()
}
