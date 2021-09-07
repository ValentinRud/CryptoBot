package main

import (
	"Project/um"
)

func main() {

	//Ошибка с полями в структуре относящимуся к json полученного от api
	//Как вызвать поля структуры из package um в package db (что бы сработало я перенес функцию в um)
	//bot.Tbot()
	//api.ApiInfo()
	um.Js()
	um.InsertDb()

}
