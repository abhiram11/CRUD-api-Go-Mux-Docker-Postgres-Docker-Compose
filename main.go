package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq" // postgres driver for database/sql package

)

type User struct {
	ID       int     `json:"id"`
	Name	string	`json:"name"`
	Email	string `json:"email`
}

func main() {
	//connect to database
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	//connection successful or no
	if err != nil {
		log.Fatal(err)
	}

	// this is called at the end of the main function, as a closing/cleanup function
	defer db.Close()

	//creater router to handle different requests
	router := mux.NewRouter() //.StrictSlash(true)
	router.HandleFunc("/users", getUsers(db)).Methods("GET")
	router.HandleFunc("/users/{id}", getUser(db)).Methods("GET")
	router.HandleFunc("/users", createUser(db)).Methods("POST")
	router.HandleFunc("/users/{id}", updateUser(db)).Methods("PUT")
	router.HandleFunc("/users/{id}", deleteUser(db)).Methods("DELETE")

	// start server
	log.Fatal(http.ListenAndServe(":8000", jsonContentTypeMiddleware(router) )) // hardcoded for now

}

func jsonContentTypeMiddleware ( next http.Handler) http.Handler {
	return  http.HandlerFunc( func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	}
}

// defining functions/ controllers as fdefined in router.handleFuync