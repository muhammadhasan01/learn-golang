package migrations

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"hasan.com/go-bank-backend/helpers"
	"hasan.com/go-bank-backend/interfaces"
)

func createAccounts() {
	db := helpers.ConnectDB()

	defer db.Close()

	users := &[2]interfaces.User{
		{Username: "Hasan", Email: "hasan@hasan.com"},
		{Username: "Husen", Email: "husen@husen.com"},
	}

	for i := range users {
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := &interfaces.User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
		db.Create(&user)

		account := &interfaces.Account{Type: "Daily Account", Name: string(users[i].Username + "'s" + " account"), Balance: uint(10000 * int(i+1)), UserID: user.ID}
		db.Create(&account)
	}
}

// Migrate is a function used to migrate data from golang to the prefered databse
func Migrate() {
	User := &interfaces.User{}
	Account := &interfaces.Account{}
	db := helpers.ConnectDB()
	db.AutoMigrate(&User, &Account)
	defer db.Close()

	createAccounts()
}
