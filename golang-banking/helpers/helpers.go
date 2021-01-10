package helpers

import (
	"regexp"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
	"hasan.com/go-bank-backend/interfaces"
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

// Validation is a function to validate values such as username and email
func Validation(values []interfaces.Validation) bool {
	username := regexp.MustCompile("^([A-Za-z0-9]{5,})+$")
	email := regexp.MustCompile("^[A-Za-z0-9]+[@]+[A-Za-z0-9]+[.]+[A-Za-z]+$")

	for i := 0; i < len(values); i++ {
		switch values[i].Valid {
		case "username":
			if !username.MatchString(values[i].Value) {
				return false
			}
		case "email":
			if !email.MatchString(values[i].Value) {
				return false
			}
		case "password":
			if len(values[i].Value) < 5 {
				return false
			}
		}
	}
	return true
}
