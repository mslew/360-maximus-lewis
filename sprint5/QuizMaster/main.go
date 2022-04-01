package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"

	//"net/http"
	"os"
	"strconv"
)

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

type ReturnFunction1 struct {
	ID                []string
	name              []string
	numberOfQuestions []string
}
type ReturnFunction2 struct {
	Questions []string
	Options   []string
}

func ReadFromJson() ReturnFunction1 {
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

	returnFunction1 := ReturnFunction1{}

	for i := 0; i < len(databaseArray.Databases); i++ {
		returnFunction1.ID = append(returnFunction1.ID, databaseArray.Databases[i].ID)
		returnFunction1.name = append(returnFunction1.name, databaseArray.Databases[i].Name)
		returnFunction1.numberOfQuestions = append(returnFunction1.numberOfQuestions, databaseArray.Databases[i].NumberOfQuestions)
	}
	return returnFunction1
}

func main() {
	returnFunction := ReadFromJson()
	questionOptions := ShowQuestions()

	fmt.Println(returnFunction)

	r := rand.New(rand.NewSource(time.Now().Unix()))
	for _, i := range r.Perm(len(questionOptions.Questions)) {
		fmt.Println(questionOptions.Questions[i])
		fmt.Println(questionOptions.Options[i])
	}

	//fmt.Println(questionOptions.Questions)
	//fmt.Println(questionOptions.Options)

}

//w http.ResponseWriter, r *http.Request
func ShowQuestions() ReturnFunction2 {
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
	returnFunction2 := ReturnFunction2{}

	var strArray = []string{"1", "2", "2", "1"}
	var intArray = []int{0}

	for _, i := range strArray {
		j, err3 := strconv.Atoi(i)
		if err3 != nil {
			panic(err)
		}
		intArray = append(intArray, j)
	}
	for dCount := 0; dCount < len(databaseArray.Databases); dCount++ { // 1 - 2 databases
		for inCount := 0; inCount < len(strArray); inCount += 2 { // looping trhough input array
			if strArray[inCount] == databaseArray.Databases[dCount].ID { //matching id's

				if inCount+1 <= len(intArray) { //check proper num Q
					for qCount := 0; qCount < intArray[inCount+1]; qCount++ { //handle num Q
						returnFunction2.Questions = append(returnFunction2.Questions, databaseArray.Databases[dCount].Questions[qCount])
						returnFunction2.Options = append(returnFunction2.Options, databaseArray.Databases[dCount].Options[qCount])
					}
				}
			}
		}
	}
	return returnFunction2

}
