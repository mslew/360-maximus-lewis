package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type DatabaseArray struct {
	Databases []struct {
		ID                string     `json:"id"`
		Name              string     `json:"name"`
		NumberOfQuestions string     `json:"number_of_questions"`
		Questions         []string   `json:"questions"`
		Options           [][]string `json:"options"`
		Answer            []string   `json:"answer"`
	} `json:"databases"`
}

type ReturnFunction1 struct {
	ID                []string
	name              []string
	numberOfQuestions []string
}

/*// let's declare a global Articles array
// that we can then populate in our main function
// to simulate a database
var QDB []DB

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
func returnAllDBs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnDBs")
	json.NewEncoder(w).Encode(QDB)
}
func returnSingleDB(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	//fmt.Fprintf(w, "Key: " + key)

	for _, article := range QDB {
		if article.id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllDBs)

	myRouter.HandleFunc("/DB/{id}", returnSingleDB)

	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}*/

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
	/*QDB = []DB{
		// DB{Id: "123", Name: "DB_H", Desc: "Article Description", Content: "Article Content"},
		DB{id: "123", name: "DB_H", number_of_questions: "4", options: "Article Content", answer: "{a,b,c,d,e}"},
		DB{id: "456", name: "DB_S", number_of_questions: "6", options: "FILLER", answer: "{b,x,d,x}"},
		//DB{Id: "456", Name: "DB_S ", Desc: "Article Description", Content: "Article Content"},
	}*/

	returnFunction := ReadFromJson()

	fmt.Println(returnFunction)
}
