package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
	"time"
)

type postMessage struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Create string `json:"create"`
	Due    string `json:"due"`
}

type noteMessage struct {
	Title  string
	Body   string
	Create time.Time
	Due    time.Time
	Id     uuid.UUID
}

type ResponseStruct struct {
	Msg string `json:"msg"`
}

var db *sql.DB

func main() {
	initDB()
	defer db.Close()
	router := mux.NewRouter()
	router.HandleFunc("/note", addNote).Methods("POST")
	router.HandleFunc("/note", getNotes).Methods("GET")
	router.HandleFunc("/note/{id}", deleteNote).Methods("DELETE")
	router.HandleFunc("/note/{id}", updateNote).Methods("PUT")
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

func timeFormat(t string) (time.Time, error) {
	const layout = "2006/01/02 15:04"
	tim, err := time.Parse(layout, t)
	if err != nil {
		fmt.Println(err.Error())
		return time.Time{}, err
	}
	return tim, nil
}

func addNote(res http.ResponseWriter, req *http.Request) {
	var recivedNote postMessage
	var noteToSave noteMessage
	err := json.NewDecoder(req.Body).Decode(&recivedNote)
	if err != nil {
		fmt.Println(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(fmt.Sprintf("something went wrong with json parsing")))
		return
	}
	noteToSave.Title = recivedNote.Title
	noteToSave.Body = recivedNote.Body
	noteToSave.Create, err = timeFormat(recivedNote.Create)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(fmt.Sprintf("something went wrong with time parsing")))
		return
	}
	noteToSave.Due, err = timeFormat(recivedNote.Due)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(fmt.Sprintf("something went wrong with time parsing")))
		return
	}
	noteToSave.Id = uuid.New()

	inSt := `INSERT INTO notes(Title, Body, CreateAt, Due, Id) values($1, $2, $3, $4, $5)`
	_, err = db.Exec(inSt, noteToSave.Title, noteToSave.Body, noteToSave.Create, noteToSave.Due, noteToSave.Id)
	if err != nil {
		fmt.Println(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(fmt.Sprintf("something went wrong with saving to db")))
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	resp := ResponseStruct{
		Msg: fmt.Sprintf("added new note"),
	}
	json.NewEncoder(res).Encode(resp)
}

func getNotes(res http.ResponseWriter, req *http.Request) {
	fmt.Println("get see notes")
	rows, err := db.Query("SELECT * FROM notes")
	if err != nil {
		fmt.Println(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(fmt.Sprintf("something went wrong with reading from db")))
	}
	defer rows.Close()
	var notes []noteMessage
	var rs noteMessage
	for rows.Next() {
		var t string
		var b string
		var c time.Time
		var d time.Time
		var i uuid.UUID
		rows.Scan(&t, &b, &c, &d, &i)
		rs.Title = t
		rs.Body = b
		rs.Create = c
		rs.Due = d
		rs.Id = i
		notes = append(notes, rs)
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	val := map[string][]noteMessage{"msg": notes}
	jsonVal, _ := json.Marshal(val)
	res.Write(jsonVal)

}

func deleteNote(res http.ResponseWriter, req *http.Request) {
	fmt.Println("delete note")
	fmt.Fprintf(res, "delet note")

}

func updateNote(res http.ResponseWriter, req *http.Request) {
	fmt.Println("put update note")
	fmt.Fprintf(res, "update note")

}
