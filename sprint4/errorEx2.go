package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
* A function that ensures reading an integer variable within a given range
* inputs: range lower and upper limits
* output: an integer value within the given range
* pre-condition: upper and lower limits are integers, lower<=upper
 */
func getIntegerV1(lower int, upper int) (int, error) {
	var value int
	var err error
	for stop := false; stop != true; {
		fmt.Printf("Enter an integer value in the range [%d, %d]:", lower, upper)
		_, err := fmt.Scanf("%d\n", &value)
		if err != nil {
			fmt.Println("Error, please try again...")

		} else if value >= lower && value <= upper {
			stop = true
		}

	}

	return value, err
}

/*
* A function that ensures reading an integer variable within a given range
* inputs: range lower and upper limits
* output: an integer value within the given range
* pre-condition: upper and lower limits are integers, return error otherwise
 */

func getIntegerV2(lower int64, upper int64) (int64, error) {
	var err error
	var value int64
	if lower > upper {
		return 0, errors.New("Error: Upper limit should be >= lower")
	}
	reader := bufio.NewReader(os.Stdin)
	for stop := false; stop != true; {
		fmt.Printf("Enter an integer value in the range [%d, %d]:", lower, upper)
		strValue, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error, please try again...")

		} else {
			fValue, err2 := strconv.ParseFloat(strings.TrimSpace(strValue), 64)
			value = int64(fValue)
			if err2 != nil {
				fmt.Println("Error, please try again...")

			} else if value >= lower && value <= upper {
				stop = true
			}
		}
	}
	return value, err
}

func main() {
	result, err := getIntegerV2(10, 15)

	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println(result)
	}
}
