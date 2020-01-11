package main

import (
	//"../../entity"
	"../../balance/brepository"
	"../../balance/bservice"
	"../../users/repository"
	"../../users/service"
	"../http/handler"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"html/template"
	"net/http"
)

var db *sql.DB

// var tmpl = template.Must(template.ParseGlob("../../ui/*.html"))

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "onlinedb"
)

var userService *service.UserService
var balanceService *bservice.BalanceService

func main() {
	/* Database connection */
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	fmt.Println("DB connected sucessfully")
	if err = db.Ping(); err != nil {
		panic(err)
	}

	/* ---------------------- */

	tmpl := template.Must(template.ParseGlob("../../ui/*.html"))

	usrRepo := repository.NewUserPostRepo(db)
	userService = service.NewUserService(usrRepo)

	br := brepository.NewBalanceRepo(db)
	balanceService = bservice.NewBalanceService(br)

	adminUserHandler := handler.NewAdminUserHandler(tmpl, userService, balanceService)

	fs := http.FileServer(http.Dir("../../ui/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/signuppage", adminUserHandler.Signuppage)
	http.HandleFunc("/", adminUserHandler.Index_handler)
	http.HandleFunc("/signinpage", adminUserHandler.Login)
	http.HandleFunc("/signup", adminUserHandler.AdminRegistration)
	http.HandleFunc("/signin", adminUserHandler.AdminLogin)
	http.HandleFunc("/profile", adminUserHandler.ProfileHandler)
	http.ListenAndServe(":8080", nil)

}
