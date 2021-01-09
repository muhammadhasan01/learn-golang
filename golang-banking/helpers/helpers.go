package helpers

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

// HandleErr is a function to handle error, it will panic an error message
func HandleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// HashAndSalt is a function to hash password and retrieve a hashed password
func HashAndSalt(pass []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	HandleErr(err)

	return string(hashed)
}

// ConnectDB is a function to connect to the postgre database
func ConnectDB() *gorm.DB {
	const dsn = "user=postgres password=password dbname=testgorm sslmode=disable"
	db, err := gorm.Open("postgres", dsn)
	HandleErr(err)
	return db
}
