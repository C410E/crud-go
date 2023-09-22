package main

//"text/template"
//"github.com/go-sql-driver/mysql"

import (
	"database/sql"
	"fmt"
	"net/http"
)

type Names struct {
	Id    int
	Name  string
	Email string
}

func dbconn(db *sql.DB) {
	dbDriver := "mysql"
	dbUser := ""
	dbPass := ""
	dbName := ""

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"+ @/"+dbName)
	if err != nil {
		panic(err.Error())
	}

	return db

}

func main() {
	fmt.Println("Server iniciado na porta http://localhost:8000")

	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)

	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)

	http.ListenAndServe(":8000", nil)
}
