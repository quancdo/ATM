package common

import (
	"log"

	"github.com/ATM/pkg/database"
	"github.com/ATM/pkg/models"
)

func GetAccounts() ([]models.Account, error) {
	db := database.GetDBInstance()
	accounts := []models.Account{}

	resp := db.Find(&accounts)

	if resp.RecordNotFound() {
		log.Print("Account not found")
		return nil, resp.Error
	}

	if resp.Error != nil {
		log.Print("Internal error")
		return nil, resp.Error
	}

	return accounts, nil
}

func GetAccountByAPIkey(api_key string) (*models.Account, error) {
	db := database.GetDBInstance()
	var account = &models.Account{}

	resp := db.Where("api_key = ?", api_key).First(account)
	if resp.RecordNotFound() {
		log.Print("Account not found")
		return nil, resp.Error
	}

	if resp.Error != nil {
		log.Print("Internal error")
		return nil, resp.Error
	}
	return account, nil
}
