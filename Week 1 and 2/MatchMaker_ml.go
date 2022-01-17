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
	fmt.Printf("Enter a number between 1 and 5. 1 being you disagree and 5 being you agree.\n")
}

//footer
func footer() {
	fmt.Printf("Thank you for playing!\n")
}

//creating a question object
type question struct {
	question         string
	answer           float64
	idealAnswer      float64
	answerDifference float64
}

//question 1
func questionOne() question {
	var answer float64
	question := question{question: "Is Steely Dan the best band ever (1-5)?\n", answer: 0, idealAnswer: 5, answerDifference: 0}
	fmt.Printf(question.question)
	fmt.Scanf("%f\n", &answer)
	question.answerDifference = math.Abs(question.idealAnswer - answer)
	return question
}

//question 2
func questionTwo() question {
	var answer float64
	question := question{question: "Is broccoli the best vegetable (1-5)?\n", answer: 0, idealAnswer: 4, answerDifference: 0}
	fmt.Printf(question.question)
	fmt.Scanf("%f\n", &answer)
	question.answerDifference = math.Abs(question.idealAnswer - answer)
	return question
}

//question 3
func questionThree() question {
	var answer float64
	question := question{question: "Are you a fan of candy corn (1-5)?\n", answer: 0, idealAnswer: 1, answerDifference: 0}
	fmt.Printf(question.question)
	fmt.Scanf("%f\n", &answer)
	question.answerDifference = math.Abs(question.idealAnswer - answer)
	return question
}

//question 4
func questionFour() question {
	var answer float64
	question := question{question: "Hiking is the best activity (1-5)?\n", answer: 0, idealAnswer: 3, answerDifference: 0}
	fmt.Printf(question.question)
	fmt.Scanf("%f\n", &answer)
	question.answerDifference = math.Abs(question.idealAnswer - answer)
	return question
}

//question 5
func questionFive() question {
	var answer float64
	question := question{question: "Milk chocolate is the best type of chocolate (1-5)?\n", answer: 0, idealAnswer: 2, answerDifference: 0}
	fmt.Printf(question.question)
	fmt.Scanf("%f\n", &answer)
	question.answerDifference = math.Abs(question.idealAnswer - answer)
	return question
}

//display answer also do some math and fancy printing techniques
func printVerdict(questionOne question, questionTwo question, questionThree question, questionFour question, questionFive question) {
	var verdict float64 = questionOne.answerDifference + questionTwo.answerDifference + questionThree.answerDifference + questionFour.answerDifference + questionFive.answerDifference
	list := [5]question{questionOne, questionTwo, questionThree, questionFour, questionFive}
	fmt.Printf("**************************************\n")
	for i := 0; i < 5; i++ {
		fmt.Printf("The ideal answer for %s was %.1f making the difference %.1f.\n", list[i].question, list[i].idealAnswer, list[i].answerDifference)
	}
	fmt.Printf("**************************************\n")
	fmt.Printf("Your final score was %f.\n", verdict)

	if 0 >= verdict || verdict <= 5 {
		fmt.Printf("Let's get married!\n")
	} else if 6 >= verdict || verdict <= 10 {
		fmt.Printf("Let's go on a date!\n")
	} else if verdict >= 11 {
		fmt.Printf("You're not the one for me.\n")
	}
}

//driver code
func main() {
	header()
	var questionOne question = questionOne()
	var questionTwo question = questionTwo()
	var questionThree question = questionThree()
	var questionFour question = questionFour()
	var questionFive question = questionFive()
	printVerdict(questionOne, questionTwo, questionThree, questionFour, questionFive)
	footer()
}
