package users

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"hasan.com/go-bank-backend/helpers"
	"hasan.com/go-bank-backend/interfaces"
)

// Login is a function that helps user login, it returs a map<string, interface>
func Login(username string, pass string) map[string]interface{} {

	// Add validation to login
	valid := helpers.Validation(
		[]interfaces.Validation{
			{Value: username, Valid: "username"},
			{Value: pass, Valid: "password"},
		})

	if !valid {
		return map[string]interface{}{"message": "not valid values"}
	}

	db := helpers.ConnectDB()
	defer db.Close()

	user := &interfaces.User{}

	if db.Where("username = ? ", username).First(&user).RecordNotFound() {
		return map[string]interface{}{"message": "User not found"}
	}

	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))

	if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
		return map[string]interface{}{"message": "Wrong password"}
	}

	accounts := []interfaces.ResponseAccount{}
	db.Table("accounts").Select("id, name, balance").Where("user_id = ? ", user.ID).Scan(&accounts)

	response := prepareResponse(user, accounts)

	return response
}

// Register is a function to register a user
func Register(username string, email string, pass string) map[string]interface{} {
	// Add validation to registration
	valid := helpers.Validation(
		[]interfaces.Validation{
			{Value: username, Valid: "username"},
			{Value: email, Valid: "email"},
			{Value: pass, Valid: "password"},
		})

	if !valid {
		return map[string]interface{}{"message": "not valid values"}
	}

	db := helpers.ConnectDB()
	defer db.Close()

	generatedPassword := helpers.HashAndSalt([]byte(pass))
	user := &interfaces.User{Username: username, Email: email, Password: generatedPassword}
	db.Create(&user)

	account := &interfaces.Account{Type: "Daily Account", Name: string(username + "'s" + " account"), Balance: 0, UserID: user.ID}
	db.Create(&account)

	accounts := []interfaces.ResponseAccount{}
	respAccount := interfaces.ResponseAccount{ID: account.ID, Name: account.Name, Balance: int(account.Balance)}
	accounts = append(accounts, respAccount)
	var response = prepareResponse(user, accounts)

	return response
}

func prepareToken(user *interfaces.User) string {
	tokenContent := jwt.MapClaims{
		"user_id": user.ID,
		"expiry":  time.Now().Add(time.Minute * 60).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("TokenPassword"))
	helpers.HandleErr(err)

	return token
}

func prepareResponse(user *interfaces.User, accounts []interfaces.ResponseAccount) map[string]interface{} {
	responseUser := &interfaces.ResponseUser{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Accounts: accounts,
	}

	var token = prepareToken(user)
	var response = map[string]interface{}{"message": "all is fine"}
	response["jwt"] = token
	response["data"] = responseUser

	return response
}
