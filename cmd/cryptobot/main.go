package main

import (
	"Project/internal/app/gateways"
	"Project/internal/app/handlers"
	"Project/internal/app/services"
)

func main() {

	//Ошибка с полями в структуре относящимуся к json полученного от api
	//Как вызвать поля структуры из package um в package db (что бы сработало я перенес функцию в um
	//	api.ApiInfo()
	//	um.Js()
	//	um.InsertDb()
	//	bot.Telebot()

	coinmarketcapGateway := gateways.NewCoinMarketCapGateway()
	balanceService := services.NewBalanceService(coinmarketcapGateway)
	telegramHandler := handlers.NewTelegramHandler(balanceService)

	telegramHandler.Telebot()
}
