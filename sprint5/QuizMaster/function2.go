package s

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"net/http"
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
type ReturnFunction2 struct {
	Questions []string
	Options   []string
}

type StrArray struct {
	ID []string `json:"id"`
}

func ShowQ(w http.ResponseWriter, r *http.Request) {
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
	//return databaseArray

	//var idArray []string

	strArray := StrArray{}
	err3 := json.NewDecoder(r.Body).Decode(&strArray)
	if err3 != nil {
		fmt.Println(err3.Error())
		return
	}
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
	b, _ := json.Marshal(returnFunction2)
	w.Write(b)

}

/*
  for _, v := range r {
      if databaseArray.ID == v {
          fmt.Println(databaseArray)
        //fmt.Printf(json.NewEncoder(w).Encode(databaseArray))
      }

  }
*/
