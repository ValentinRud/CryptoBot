package um

import (
	"Project/config"
	structApi "Project/struct"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/lib/pq"
)

var ag structApi.User

func Js() {

	file, err := os.Open("data1.json")
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	jerr := json.Unmarshal(data, &ag)
	if jerr != nil {
		fmt.Println(jerr)
	}

	fmt.Println(ag)

}

func InsertDb() {
	db, err := sql.Open("postgres", config.ConnStr)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	result, err := db.Query("insert into test (id, first_name, last_name, age, status) values ($1,$2,$3,$4,$5)", ag.Id, ag.FirstName, ag.LastName, ag.Age, ag.Status)
	if err != nil {
		panic(err)
	}
	defer result.Close()

}
