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
	question         string  //the question
	answer           float64 //answer, from the user
	idealAnswer      float64 //my ideal answer
	answerDifference float64 //the difference between mine and user's answer
	weight           float64 //how important this is to me
}

const (
	QUESTION_ONE_IDEAL  float64 = 5
	QUESTION_ONE_WEIGHT float64 = 3
	QUESTION_ONE_RANGE  float64 = 0

	QUESTION_TWO_IDEAL  float64 = 4
	QUESTION_TWO_WEIGHT float64 = 2
	QUESTION_TWO_RANGE  float64 = 0

	QUESTION_THREE_IDEAL  float64 = 1
	QUESTION_THREE_WEIGHT float64 = 3
	QUESTION_THREE_RANGE  float64 = 0

	QUESTION_FOUR_IDEAL  float64 = 3
	QUESTION_FOUR_WEIGHT float64 = 2
	QUESTION_FOUR_RANGE  float64 = 0

	QUESTION_FIVE_IDEAL  float64 = 2
	QUESTION_FIVE_WEIGHT float64 = 1
	QUESTION_FIVE_RANGE  float64 = 0
)

func validate(question question) question {
	var answer float64
	for {
		fmt.Printf(question.question)
		a, err := fmt.Scanf("%f", &answer)
		if a == 0 || answer > 5 || answer < 1 || err != nil {
			fmt.Printf("Please enter a valid number between 1 and 5.\n")
		} else {
			question.answer = answer
			return question
		}
	}
}

//question 1
func questionOne() question {
	question := question{question: "Is Steely Dan the best band ever (1-5)?\n",
		answer: 0, idealAnswer: QUESTION_ONE_IDEAL, answerDifference: QUESTION_ONE_RANGE,
		weight: QUESTION_ONE_WEIGHT}
	validate(question)
	question.answerDifference = math.Abs(question.idealAnswer - question.answer)
	return question
}

//question 2
func questionTwo() question {
	question := question{question: "Is broccoli the best vegetable (1-5)?\n",
		answer: 0, idealAnswer: QUESTION_TWO_IDEAL, answerDifference: QUESTION_TWO_RANGE,
		weight: QUESTION_TWO_WEIGHT}
	validate(question)
	question.answerDifference = math.Abs(question.idealAnswer - question.answer)
	return question
}

//question 3
func questionThree() question {
	question := question{question: "Are you a fan of candy corn (1-5)?\n",
		answer: 0, idealAnswer: QUESTION_THREE_IDEAL, answerDifference: QUESTION_THREE_RANGE,
		weight: QUESTION_THREE_WEIGHT}
	validate(question)
	question.answerDifference = math.Abs(question.idealAnswer - question.answer)
	return question
}

//question 4
func questionFour() question {
	question := question{question: "Hiking is the best activity (1-5)?\n",
		answer: 0, idealAnswer: QUESTION_FOUR_IDEAL, answerDifference: QUESTION_FOUR_RANGE,
		weight: QUESTION_FOUR_WEIGHT}
	validate(question)
	question.answerDifference = math.Abs(question.idealAnswer - question.answer)
	return question
}

//question 5
func questionFive() question {
	question := question{question: "Milk chocolate is the best type of chocolate (1-5)?\n",
		answer: 0, idealAnswer: QUESTION_FIVE_IDEAL, answerDifference: QUESTION_FIVE_RANGE,
		weight: QUESTION_FIVE_WEIGHT}
	validate(question)
	question.answerDifference = math.Abs(question.idealAnswer - question.answer)
	return question
}

//display answer also do some math and fancy printing techniques
func printVerdict(questionOne question, questionTwo question, questionThree question, questionFour question, questionFive question) {
	var verdict float64 = (questionOne.answerDifference * questionOne.weight) +
		(questionTwo.answerDifference * questionTwo.weight) +
		(questionThree.answerDifference * questionThree.weight) +
		(questionFour.answerDifference * questionFour.weight) +
		(questionFive.answerDifference * questionFive.weight)
	score := 100 - verdict
	list := [5]question{questionOne, questionTwo, questionThree, questionFour, questionFive}
	for i := 0; i < 5; i++ {
		fmt.Printf("************************************\n")
		fmt.Printf("The compatibility score for \n%s was %.2f\n", list[i].question, list[i].answerDifference)
		fmt.Printf("The weighted score is: %.2f\n", list[i].answerDifference*list[i].weight)
	}
	fmt.Printf("**************************************\n")
	fmt.Printf("Your final score was %.2f.\n", score)

	if score >= 90 && score <= 100 {
		fmt.Printf("Let's get married!\n")
	} else if score >= 70 && score <= 89 {
		fmt.Printf("Let's go on a date!\n")
	} else if score <= 69 {
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
