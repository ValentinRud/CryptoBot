package gateways

import (
	"Project/config"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type CoinMarketCapGateway struct {
}

func NewCoinMarketCapGateway() *CoinMarketCapGateway {
	return &CoinMarketCapGateway{}
}

func (s *CoinMarketCapGateway) GetBalance() (interface{}, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", config.ServerUrl, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := url.Values{}
	q.Add("start", "1")
	q.Add("limit", "1")
	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", config.PgvTestBotToken)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to server")
		os.Exit(1)
	}

	fmt.Println(resp.Status)

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(respBody))

	return string(respBody), nil
}
