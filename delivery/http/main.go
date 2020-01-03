package main

import (
	"../../entity"
	"../../users/repository"
	"../../users/service"
	"database/sql"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
)

var db *sql.DB
var tmpl = template.Must(template.ParseGlob("../../ui/*.html"))
var userService *service.UserService

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "dagim@123"
	dbname   = "onlinedb"
)

func signuppage(w http.ResponseWriter, req *http.Request) {
	_ = tmpl.ExecuteTemplate(w, "signup.html", nil)
}
func index_handler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		panic(err.Error())
	}

}

func login(w http.ResponseWriter, req *http.Request) {
	err := tmpl.ExecuteTemplate(w, "signIn.html", nil)
	if err != nil {
		panic(err.Error())
	}

}

func Registration(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Redirect(w, req, "/registration", http.StatusSeeOther)
		return
	}
	usr := entity.User{}
	usr.FirstName = req.FormValue("first")
	usr.LastName = req.FormValue("last")
	usr.Email = req.FormValue("email")
	usr.Password = req.FormValue("password")
	// a := req.FormValue("secondpassword")
	// if a == usr.Password{

	// }

	hashedpass, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)

	usr.Password = string(hashedpass)

	err = userService.StoreUser(usr)
	if err != nil {
		panic(err.Error())
	}
	http.Redirect(w, req, "/signinpage", 303)
	if err != nil {
		panic(err.Error())
	}

}
func Login(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Redirect(w, req, "/signinpage", http.StatusSeeOther)
		return
	}
	email := req.FormValue("email")
	password := req.FormValue("password")

	log.Println(email)
	usr, err := userService.User(email)

	if err != nil {
		log.Println("Username or Password is incorrect")
		http.Redirect(w, req, "/signinpage", 303)
		return
	}

	if email == usr.Email {
		err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(password))
		log.Println("Reached here")
		if err != nil {
			panic(err.Error())
		}
		user, err := userService.UserwithID(int(usr.ID))
		if err != nil {
			panic(err.Error())
		}
		err = tmpl.ExecuteTemplate(w, "profile.html", user)
	}
	http.Redirect(w, req, "/signinpage", 303)

	//case sql.ErrNoRows
	//}
	//err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(password))
	if err != nil {
		panic(err.Error())
	}

}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	// defer db.Close()
	fmt.Println("DB connected sucessfully")
	if err = db.Ping(); err != nil {
		panic(err)
	}

	usrRepo := repository.NewUserPostRepo(db)
	userService = service.NewUserService(usrRepo)

	fs := http.FileServer(http.Dir("../../ui/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/signuppage", signuppage)
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/signinpage", login)
	http.HandleFunc("/signup", Registration)
	http.HandleFunc("/signin", Login)

	http.ListenAndServe(":8080", nil)

}
