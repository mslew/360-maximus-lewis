package readmechecker

import (
	"fmt"
	"io/ioutil"
	"os"
)

func scan() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		os.Exit(0)
	}
	for _, f := range files {
		if f.Name() != "README" {
			fmt.Println("You need a README file in your project.")
		} else {
			fmt.Println("README checker success.")
		}
	}
}
