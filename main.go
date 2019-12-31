package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB
var temple = template.Must(template.ParseFiles("index.html", "portfolio.html", "signup.html"))

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "dagim@123"
	dbname   = "shopusersdb"
)

func dbConn() (db *sql.DB) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	//defer db.Close()
	fmt.Println("DB Connected sucessfully !")
	return db

}
func init() {
	db = dbConn()
}

func index_handler(w http.ResponseWriter, r *http.Request) {
	temple.ExecuteTemplate(w, "index.html", nil)

}
func signup_handler(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("first") != "" {
		first := r.FormValue("first")
		last := r.FormValue("last")
		email := r.FormValue("email")
		password := r.FormValue("password")
		second := r.FormValue("secondpassword")
		if second == password {

			person := Person{first, last, email, password}
			Insert(db, person)
			//temple.ExecuteTemplate(w, "signup.html", nil)
			http.Redirect(w, r, "/", 303)
		} else {
			http.Redirect(w, r, "/signup", 303)
		}
		return
	}

	temple.ExecuteTemplate(w, "signup.html", nil)

}
func port_handler(w http.ResponseWriter, r *http.Request) {

	email := r.FormValue("email")
	password := r.FormValue("password")
	var i = Read(db, email, password)
	// var i = ReadData(email, password)
	if i == "0" {
		http.Redirect(w, r, "/", 303)
	}
	temple.ExecuteTemplate(w, "portfolio.html", i)
}
func main() {

	fz := http.FileServer(http.Dir("assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fz))
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/signup", signup_handler)
	http.HandleFunc("/portfolio", port_handler)

	http.ListenAndServe(":8080", nil)
}

type Person struct {
	FirstName string
	lastName  string
	email     string
	password  string
}

func Insert(db *sql.DB, person Person) {
	var stat = "INSERT INTO users(firstname,lastname,email,password)VALUES($1,$2,$3,$4)"
	_, err := db.Exec(stat, person.FirstName, person.lastName, person.email, person.password)
	if err != nil {
		panic(err)
	}
}
func Read(db *sql.DB, email string, password string) string {
	var stat = "SELECT firstname,lastname FROM users WHERE email=$1 AND password=$2"
	row := db.QueryRow(stat, email, password)
	var first string
	var last string
	switch err := row.Scan(&first, &last); err {
	case sql.ErrNoRows:
		return "0"
	case nil:
		return first + " " + last
	default:
		panic(err)
	}
}
