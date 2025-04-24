package tools

import (
	"time"
)

type mockDB struct{}

// GetUserCoins implements DatabaseInterface.
func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	coinData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &coinData
}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username:  "alex",
	},
	"jason": {
		AuthToken: "456DEF",
		Username:  "jason",
	},
	"marie": {
		AuthToken: "789GHI",
		Username:  "marie",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    100,
		Username: "alex",
	},
	"jason": {
		Coins:    200,
		Username: "jason",
	},
	"marie": {
		Coins:    300,
		Username: "marie",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	//simulateDB CALL
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

// setup database
func (d *mockDB) SetupDatabase() error {
	return nil
}
