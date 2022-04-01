package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	//"net/http"
	"os"
	"strconv"
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
type ReturnFunction2 struct {
	Questions []string
	Options   []string
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
	returnFunction := ReadFromJson()
	questionOptions := ShowQuestions()

	fmt.Println(returnFunction)

	fmt.Println(questionOptions.Questions)
	fmt.Println(questionOptions.Options)
	//fmt.Println(randArray(questionOptions.Options))

	//fmt.Println("\n", questionOptions)
}

/*func randArray(src []string) []string {
	final := make([]string, len(src))
	rand.Seed(time.Now().UnixNano())
	perm := rand.Perm(len(src))

	for i, v := range perm {
		final[v] = src[i]
	}
	return final
}
*/

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
						for oCount := 0; oCount <= len(databaseArray.Databases[dCount].Options); oCount++ {
							returnFunction2.Options = append(returnFunction2.Options, databaseArray.Databases[dCount].Options[qCount][oCount])
						}
					}
				}
			}
		}
	}
	return returnFunction2

}
