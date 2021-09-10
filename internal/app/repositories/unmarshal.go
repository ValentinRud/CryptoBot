package um

import (
	"Project/config"
	"Project/internal/app/models"
	structApi "Project/struct"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/lib/pq"
)

// 1
// Сделать структуру репозитория с двумя методами
// Create (insert) func .. Create(..) error
// List (select списка) func ... List(..)
// Открыть соединение с базой один раз
// Пробросить сущность базы аргументом в конструкторе репозитория

// 2
// Получить json в сообщение из телеграма и сохранить его в базу

// 3
// Если ты получишь команду /getUser, то нужно вывести пользователя из базы
// Пример сообщения, которое ты должен получить на команду /getUser:
/*
Идентификатор: 10
Имя: Имя
Фамилия: Фамилия
Возраст: 18
Статус: студент
*/

var ag models.User

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
