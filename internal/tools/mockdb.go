package tools

import (
	"time"
)

type LoginDetails struct {
	AuthToken string
	Username  string
	CreatedAt time.Time
}

type CoinDetails struct {
	Coins    int64
	Username string
	Balance  int64
}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetCoinBalance(username string) *CoinDetails
}

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "token_alex_123",
		Username:  "alex",
		CreatedAt: time.Now().Add(-24 * time.Hour),
	},
	"jordan": {
		AuthToken: "token_jordan_456",
		Username:  "jordan",
	},
	"sam": {
		AuthToken: "token_sam_789",
		Username:  "sam",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    1500,
		Username: "alex",
		Balance:  1500,
	},
	"jordan": {
		Coins:    2500,
		Username: "jordan",
		Balance:  2500,
	},
	"sam": {
		Coins:    3500,
		Username: "sam",
		Balance:  3500,
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(50 * time.Millisecond) // Simulate database lateny
	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) GetCoinBalance(username string) *CoinDetails {
	time.Sleep(50 * time.Millisecond) // Simulate database lateny
	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	// No setup needed for mock database
	return nil
}

func NewDatabase() (DatabaseInterface, error) {
	return &mockDB{}, nil
}

func GetDatabaseInstance() (DatabaseInterface, error) {
	return &mockDB{}, nil
}
