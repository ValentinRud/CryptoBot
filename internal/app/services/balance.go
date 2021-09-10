package services

import "Project/internal/app/gateways"

type BalanceService struct {
	coinmarketcapGateway *gateways.CoinMarketCapGateway
}

func NewBalanceService(coinmarketcapGateway *gateways.CoinMarketCapGateway) *BalanceService {
	return &BalanceService{
		coinmarketcapGateway: coinmarketcapGateway,
	}
}

func (s *BalanceService) GetBalance() (interface{}, error) {
	return s.coinmarketcapGateway.GetBalance()
}
