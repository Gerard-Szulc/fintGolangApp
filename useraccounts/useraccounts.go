package useraccounts

import (
	"fintGolangApp/helpers"
	"fintGolangApp/interfaces"
)

func updateAccount(id uint, amount int) {
	db := helpers.ConnectDB()
	db.Model(&interfaces.Account{}).Where("id = ?", id).Update("balance", amount)
	defer db.Close()
}
