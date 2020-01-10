package handler

import (
	"../../../entity"
	"../../../users"
	//"../../../users/repository"
	//"../../../users/service"
	//"database/sql"
	//"fmt"
	//_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
)

type AdminUserHandler struct {
	tmpl    *template.Template
	userSrv users.UserService
}

func NewAdminUserHandler(t *template.Template, ur users.UserService) *AdminUserHandler {
	return &AdminUserHandler{tmpl: t, userSrv: ur}
}

func (userService *AdminUserHandler) Index_handler(w http.ResponseWriter, req *http.Request) {
	userService.tmpl.ExecuteTemplate(w, "index.html", nil)
}

func (userService *AdminUserHandler) Login(w http.ResponseWriter, req *http.Request) {
	userService.tmpl.ExecuteTemplate(w, "signIn.html", nil)
}

func (userService *AdminUserHandler) Signuppage(w http.ResponseWriter, req *http.Request) {
	userService.tmpl.ExecuteTemplate(w, "signup.html", nil)
}

func (userService *AdminUserHandler) AdminRegistration(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Redirect(w, req, "/registration", http.StatusSeeOther)
		return
	}
	usr := entity.User{}
	usr.FirstName = req.FormValue("first")
	usr.LastName = req.FormValue("last")
	usr.Email = req.FormValue("email")
	usr.Password = req.FormValue("password")

	hashedpass, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)

	usr.Password = string(hashedpass)

	err = userService.userSrv.StoreUser(usr)
	if err != nil {
		panic(err.Error())
	}
	http.Redirect(w, req, "/signinpage", 303)
	if err != nil {
		panic(err.Error())
	}
}

func (userService *AdminUserHandler) AdminLogin(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Redirect(w, req, "/signinpage", http.StatusSeeOther)
		return
	}
	email := req.FormValue("email")
	password := req.FormValue("password")

	log.Println(email)
	usr, err := userService.userSrv.User(email)

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
		user, err := userService.userSrv.UserwithID(int(usr.ID))
		if err != nil {
			panic(err.Error())
		}
		err = userService.tmpl.ExecuteTemplate(w, "profile.html", user)
	}
	http.Redirect(w, req, "/signinpage", 303)
	if err != nil {
		panic(err.Error())
	}
}
