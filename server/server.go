package main

import (
	"database/sql"
	"fmt"
	"github.com/altereitay/todolistapp/routes"
	_ "github.com/altereitay/todolistapp/types"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
)

var db *sql.DB

func main() {
	initDB()
	defer db.Close()
	router := mux.NewRouter()
	router.HandleFunc("/note", func(w http.ResponseWriter, r *http.Request) {
		routes.AddNote(w, r, db)
	}).Methods("POST")
	router.HandleFunc("/note", func(w http.ResponseWriter, r *http.Request) {
		routes.GetNotes(w, r, db)
	}).Methods("GET")
	router.HandleFunc("/note/{id}", func(w http.ResponseWriter, r *http.Request) {
		routes.DeleteNote(w, r, db)
	}).Methods("DELETE")
	router.HandleFunc("/note/{id}", func(w http.ResponseWriter, r *http.Request) {
		routes.UpdateNote(w, r, db)
	}).Methods("PUT")
	fmt.Println("server running on port 8080")
	http.ListenAndServe(":8080", router)

}

func initDB() {
	connStr := "postgres://postgres:E!446380@localhost/todolistDB?sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
}
