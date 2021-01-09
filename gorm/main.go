package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const dsn = "user=postgres password=password dbname=testgorm sslmode=disable"

// User struct defines a data field for a user
type User struct {
	gorm.Model
	Name  string
	Email string
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic("failed to connect to the database")
	}
	defer db.Close()

	var users []User
	db.Find(&users)
	fmt.Println("{}", users)

	json.NewEncoder(w).Encode(users)
}

func newUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic("failed to connect to the database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	db.Create(&User{Name: name, Email: email})
	fmt.Fprintf(w, "New user succesfully created")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic("failed to connect to the database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)

	fmt.Fprintf(w, "Successfully deleted user")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic("failed to connect to the database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name, email := vars["name"], vars["email"]

	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)

	user.Email = email

	fmt.Fprintf(w, "Successfully updated user")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", allUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}", deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", updateUser).Methods("PUT")
	myRouter.HandleFunc("/user/{name}/{email}", newUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func initialMigration() {
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{})
}

func main() {
	fmt.Println("Go ORM Tutorial")

	// Add the call to our new initialMigration function
	initialMigration()

	// Handle Subsequent requests
	handleRequests()
}
