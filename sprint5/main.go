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

func getBankHandler(w http.ResponseWriter, r *http.Request) {
	databaseArray := DatabaseArray{}
	content, _ := os.Open("data.json")

	defer content.Close()

	byteValue, _ := ioutil.ReadAll(content)
	err2 := json.Unmarshal(byteValue, &databaseArray)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	w.Write(byteValue)

	returnFunction1 := ReturnFunction1{}

	for i := 0; i < len(databaseArray.Databases); i++ {
		returnFunction1.ID = append(returnFunction1.ID, databaseArray.Databases[i].ID)
		returnFunction1.name = append(returnFunction1.name, databaseArray.Databases[i].Name)
		returnFunction1.numberOfQuestions = append(returnFunction1.numberOfQuestions, databaseArray.Databases[i].NumberOfQuestions)
	}
}
func showQHandler(w http.ResponseWriter, r *http.Request) {
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
	strArray := StrArray{}
	strArray.ID = []string{"1", "2", "2", "1"} // how to take in input and not have constant json??????
	//err3 := json.NewDecoder(r.Body).Decode(&strArray)
	//if err3 != nil {
	//fmt.Println(err3.Error())
	//	return
	//}
	var intArray = []int{0}

	for _, i := range strArray.ID {
		j, err3 := strconv.Atoi(i)
		if err3 != nil {
			panic(err)
		}
		intArray = append(intArray, j)
	}
	returnFunction2 := ReturnFunction2{}
	for dCount := 0; dCount < len(databaseArray.Databases); dCount++ { // 1 - 2 databases
		for inCount := 0; inCount < len(strArray.ID); inCount += 2 { // looping trhough input array
			if strArray.ID[inCount] == databaseArray.Databases[dCount].ID { //matching id's

				if inCount+1 <= len(intArray) { //check proper num Q
					for qCount := 0; qCount < intArray[inCount+1]; qCount++ { //handle num Q
						returnFunction2.Questions = append(returnFunction2.Questions, databaseArray.Databases[dCount].Questions[qCount])
						returnFunction2.Options = append(returnFunction2.Options, databaseArray.Databases[dCount].Options[qCount])
					}
				}
			}
		}
	}
	//Attempting to write returnFunction2 json struct to w response Writer

	//jsonRF2, err := json.Marshal(returnFunction2)
	//if err != nil {
	//	fmt.Println(err3.Error())
	//}
	//w.Header().Set("Content-Type", "application/json")
	//w.Write(jsonRF2)

	json.NewEncoder(w).Encode(returnFunction2) // attempting to write json to http

}

func main() {
	router := mux.NewRouter()
	//testing
	//var strArray = []string{"1", "2", "2", "1"}
	strArray := StrArray{}
	strArray.ID = []string{"1", "2", "2", "1"}
	//strArray.ID = {"1","2","2","1"}

	for i := 0; i < 10; i++ {
		go router.HandleFunc("/getBank", getBankHandler) //10 multi-threaded requests
		go router.HandleFunc("/showQ", showQHandler)     //10 sort of working multi-threaded requests
	}
	//router.HandleFunc("/getBank", getBankHandler) //go to localhost:8080/getBank for testing
	//router.HandleFunc("/showQ", showQHandler) //localhost:8080/showQ
	http.ListenAndServe(":8080", router)
}
