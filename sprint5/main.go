package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

var gotBank = false

//object to store the json data
type DatabaseArray struct {
	Databases []struct {
		ID                string   `json:"id"`
		Name              string   `json:"name"`
		NumberOfQuestions string   `json:"number_of_questions"`
		Questions         []string `json:"questions"`
		Options           []string `json:"options"`
		Answer            []string `json:"answer"`
	} `json:"databases"`
}

//object to store the data from the first return function
type ReturnFunction1 struct {
	ID                []string
	name              []string
	numberOfQuestions []string
}

//object to store the data from the second return function
type ReturnFunction2 struct {
	Questions []string `json:"question"`
	Options   []string `json:"options"`
}
type StrArray struct {
	ID []string `json:"id"`
}

func Welcome() {
	fmt.Println("*********************************************************************")
	fmt.Println("                      Welcome to the QuizMaster                      ")
	fmt.Println("Follow directions or type 'help' for instructions, enter to continue.")
	fmt.Println("*********************************************************************")
}

func getBankHandler(w http.ResponseWriter, r *http.Request) {
	databaseArray := DatabaseArray{}
	content, _ := os.Open("data.json")

	defer content.Close()

	byteValue, _ := ioutil.ReadAll(content)
	err2 := json.Unmarshal(byteValue, &databaseArray)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	//w.Write(byteValue) all Bank info
	returnFunction1 := ReturnFunction1{}
	for i := 0; i < len(databaseArray.Databases); i++ {
		returnFunction1.ID = append(returnFunction1.ID, databaseArray.Databases[i].ID)
		returnFunction1.name = append(returnFunction1.name, databaseArray.Databases[i].Name)
		returnFunction1.numberOfQuestions = append(returnFunction1.numberOfQuestions, databaseArray.Databases[i].NumberOfQuestions)
	}
	gotBank = true
	fmt.Println(returnFunction1.name)
	fmt.Println(returnFunction1.ID)
	fmt.Println(returnFunction1.numberOfQuestions)
	json.NewEncoder(w).Encode(returnFunction1.name)
	json.NewEncoder(w).Encode(returnFunction1.ID)
	json.NewEncoder(w).Encode(returnFunction1.numberOfQuestions)

}
func (s *StrArray) showQHandler(w http.ResponseWriter, r *http.Request) {
	databaseArray := DatabaseArray{}
	content, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer content.Close()
	byteValue, _ := ioutil.ReadAll(content)
	err2 := json.Unmarshal(byteValue, &databaseArray)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	//s.ID = []string{"1", "2", "2", "1"}

	var intArray = []int{0}

	for _, i := range s.ID {
		j, err3 := strconv.Atoi(i)
		if err3 != nil {
			panic(err)
		}
		intArray = append(intArray, j)
	}
	returnFunction2 := ReturnFunction2{}
	for dCount := 0; dCount < len(databaseArray.Databases); dCount++ { // 1 - 2 databases
		for inCount := 0; inCount < len(s.ID); inCount += 2 { // looping trhough input array
			if s.ID[inCount] == databaseArray.Databases[dCount].ID { //matching id's

				if inCount+1 <= len(intArray) { //check proper num Q
					for qCount := 0; qCount < intArray[inCount+1]; qCount++ { //handle num Q
						returnFunction2.Questions = append(returnFunction2.Questions, databaseArray.Databases[dCount].Questions[qCount])
						returnFunction2.Options = append(returnFunction2.Options, databaseArray.Databases[dCount].Options[qCount])
					}
				}
			}
		}
	}
	fmt.Println(returnFunction2.Questions)
	fmt.Println(returnFunction2.Options)
	json.NewEncoder(w).Encode(returnFunction2) // write json to http

}

func main() {
	router := mux.NewRouter()
	var s StrArray
	var n int
	var help string
	Welcome()
	fmt.Scanln(&help)
	if help == "help" {
		fmt.Println("The quizmaster is a comprehensive Quizing program that offers users quiz questions.")
		fmt.Println("You will be offered two banks of questions to choose from.")
		fmt.Println("Enter the information for the bank and you will be shown questions.")
		//print instructions here
	}
	fmt.Print("Enter desired number of Question Banks:")
	fmt.Scanln(&n)
	s.ID = make([]string, 2*n)
	for i := 0; i < 2*n; i++ {
		if i%2 == 0 {
			fmt.Print("Enter an id: ")
		} else {
			fmt.Print("Enter number of questions: ")
		}
		fmt.Scanf("%s\n", &s.ID[i])
	}
	for i := 0; i < 10; i++ {
		go router.HandleFunc("/getBank", getBankHandler) //localhost:8080/getBank

		go router.HandleFunc("/showQ", s.showQHandler) //localhost:8080/showQ

	}

	http.ListenAndServe(":8080", router)

}
