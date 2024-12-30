package tools

import (
	"time"
)

type mockDB struct {}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123abc",
		Username: "alex",
	},
	"jason": {
		AuthToken: "456def",
		Username: "jason",
	},
	"maria": {
		AuthToken: "789ghi",
		Username: "maria",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins: 100,
		Username: "alex",
	},
	"jason": {
		Coins: 200,
		Username: "jason",
	},
	"maria": {
		Coins: 300,
		Username: "maria",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// simulate a database query
	time.Sleep(1 * time.Second)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// simulate a database query
	time.Sleep(1 * time.Second)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}