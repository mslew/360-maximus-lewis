package licensechecker

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
		if f.Name() != "LICENSE" {
			fmt.Println("You need a LICENSE file in your project.")
		} else {
			fmt.Println("LICENSE checker success.")
		}
	}
}
