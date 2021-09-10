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

func SelectDb() {
	db, err := sql.Open("postgres", config.ConnStr)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, last_name FROM test ORDER BY id ASC")
	if err != nil {
		panic(err)
	}
	sendUsers := []structApi.SendUser{}

	for rows.Next() {
		p := structApi.SendUser{}
		err := rows.Scan(&p.Id, &p.LastName)
		if err != nil {
			fmt.Println(err)
			continue
		}
		sendUsers = append(sendUsers, p)
	}

	for _, p := range sendUsers {
		fmt.Println(p.Id, p.LastName)
	}
}
