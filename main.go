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