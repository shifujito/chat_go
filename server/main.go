package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	db, err := sql.Open("postgres", "user=postgres dbname=chat password=chat")
	if err != nil {
		log.Fatalln("not connect", err)
	}
	defer db.Close()

	// sqlStatement := `INSERT INTO account (id) values (1)`
	// _, err = db.Exec(sqlStatement)
	// if err != nil {
	// 	log.Fatalln("sql syntax error", err)
	// }
}

func defalutFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/", defalutFunc)
	server.ListenAndServe()
}
