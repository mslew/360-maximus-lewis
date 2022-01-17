//Author: Max Lewis
//Making a matchmaker

package main

import (
	"fmt"
	"math"
)

//header to start the program
func header() {
	fmt.Printf("***********************\n")
	fmt.Printf("      Matchmaker!      \n")
	fmt.Printf("***********************\n")
	fmt.Printf("Enter a number between 1 and 5. 1 being you disagree and 5 being you agree\n")
}

//creating a question object
type question struct {
	question    string
	answer      int
	idealAnswer float64
}

//question 1
func questionOne() int {
	var intAnswer float64
	var strAnswer string
	questionOne := question{question: "Is Steely Dan the best band ever (1-5)?", answer: 0, idealAnswer: 5}
	fmt.Printf(questionOne.question)
	fmt.Sscanln(strAnswer, &intAnswer)
	verdict := math.Abs(questionOne.idealAnswer - intAnswer)
	return int(verdict)
}

//question 2
func questionTwo() int {
	var answer string
	questionTwo := question{question: "Is broccoli the best vegetable (1-5)?", answer: 0, idealAnswer: 4}
	fmt.Printf(questionTwo.question)
	fmt.Scanln(&answer)
	return 0
}

//question 3
func questionThree() int {
	var answer string
	questionThree := question{question: "Are you a fan of candy corn (1-5)?", answer: 0, idealAnswer: 1}
	fmt.Printf(questionThree.question)
	fmt.Scanln(&answer)
	return 0
}

//question 4
func questionFour() int {
	var answer string
	questionFour := question{question: "Hiking is the best activity (1-5)?", answer: 0, idealAnswer: 3}
	fmt.Printf(questionFour.question)
	fmt.Scanln(&answer)
	return 0
}

//question 5
func questionFive() int {
	var answer string
	questionFive := question{question: "Milk chocolate is the best type of chocolate {1-5)?", answer: 0, idealAnswer: 2}
	fmt.Printf(questionFive.question)
	fmt.Scanln(&answer)
	return 0
}

//display answer
func printVerdict() {

}

//driver code
func main() {
	header()
	questionOne()
	questionTwo()
	questionThree()
	questionFour()
	questionFive()
}
