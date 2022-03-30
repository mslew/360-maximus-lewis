package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type DatabaseArray struct {
	databases []Database `json:"database`
}

type Database struct {
	id                  []string `json:"id"`
	name                string   `json:"name"`
	number_of_questions string   `json:"number_of_questions"`
	questions           []string `json:"questions"`
	options             []string `json:"options"`
	answer              []string `json:"answer"`
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

func ReadFromJson() DatabaseArray {
	var databaseArray DatabaseArray
	content, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer content.Close()
	jsonParser := json.NewDecoder(content)
	err = jsonParser.Decode(&databaseArray)
	return databaseArray
}

func main() {
	/*QDB = []DB{
		// DB{Id: "123", Name: "DB_H", Desc: "Article Description", Content: "Article Content"},
		DB{id: "123", name: "DB_H", number_of_questions: "4", options: "Article Content", answer: "{a,b,c,d,e}"},
		DB{id: "456", name: "DB_S", number_of_questions: "6", options: "FILLER", answer: "{b,x,d,x}"},
		//DB{Id: "456", Name: "DB_S ", Desc: "Article Description", Content: "Article Content"},
	}*/

	databaseArray := ReadFromJson()

	for i := 0; i < len(databaseArray.databases); i++ {
		fmt.Printf("%v, %s, %s, %v, %v, %v", databaseArray.databases[i].id,
			databaseArray.databases[i].name,
			databaseArray.databases[i].number_of_questions,
			databaseArray.databases[i].questions,
			databaseArray.databases[i].options,
			databaseArray.databases[i].answer)
	}
}
