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
	var recivedNote PostMessage
	var noteToSave NoteMessage
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
	res.WriteHeader(http.StatusOK)
	resp := map[string]string{"msg": "added new note"}
	respJson, _ := json.Marshal(resp)
	res.Write(respJson)
}

func getNotes(res http.ResponseWriter, req *http.Request) {
	rows, err := db.Query("SELECT * FROM notes")
	if err != nil {
		fmt.Println(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(fmt.Sprintf("something went wrong with reading from db")))
	}
	defer rows.Close()
	var notes []NoteMessage
	var nm NoteMessage
	for rows.Next() {
		rows.Scan(&nm.Title, &nm.Body, &nm.Create, &nm.Due, &nm.Id)
		notes = append(notes, nm)
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	val := map[string][]NoteMessage{"msg": notes}
	jsonVal, _ := json.Marshal(val)
	res.Write(jsonVal)

}

func deleteNote(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	fmt.Println(id)
	_, err := db.Exec(`DELETE FROM notes WHERE id=$1`, id)
	if err != nil {
		fmt.Println(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(fmt.Sprintf("something went wrong with deleting from db")))
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	val := map[string]string{"msg": "deleted note"}
	jsonVal, _ := json.Marshal(val)
	res.Write(jsonVal)

}

func updateNote(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	rows, err := db.Query(`SELECT * FROM notes WHERE id=$1`, id)
	if err != nil {
		fmt.Println(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(fmt.Sprintf("something went wrong with reading from db")))
	}
	var oldNote NoteMessage
	for rows.Next() {
		rows.Scan(&oldNote.Title, &oldNote.Body, &oldNote.Create, &oldNote.Due, &oldNote.Id)
	}
	var recivedNote PostMessage
	var noteToSave NoteMessage
	err = json.NewDecoder(req.Body).Decode(&recivedNote)
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
	noteToSave.Id = oldNote.Id
	db.Exec(`UPDATE notes SET title=$1, body=$2, createat=$3, due=$4 WHERE id=$5`, noteToSave.Title, noteToSave.Body, noteToSave.Create, noteToSave.Due, oldNote.Id)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	val := map[string]string{"msg": "updated note"}
	jsonVal, _ := json.Marshal(val)
	res.Write(jsonVal)

}