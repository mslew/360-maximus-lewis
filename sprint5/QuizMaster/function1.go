package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//this reads the json data from the data.json file
//it will store it in an object of an array of database objects
func GetBankInfo(w http.ResponseWriter, r *http.Request) {
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
	w.Write(byteValue)

	returnFunction1 := ReturnFunction1{}

	for i := 0; i < len(databaseArray.Databases); i++ {
		returnFunction1.ID = append(returnFunction1.ID, databaseArray.Databases[i].ID)
		returnFunction1.name = append(returnFunction1.name, databaseArray.Databases[i].Name)
		returnFunction1.numberOfQuestions = append(returnFunction1.numberOfQuestions, databaseArray.Databases[i].NumberOfQuestions)
	}
}
