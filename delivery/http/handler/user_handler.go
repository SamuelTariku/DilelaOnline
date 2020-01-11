package handler

import (
	"../../../balance"
	"../../../entity"
	"../../../users"
	//"../../../users/brepository"
	//"../../../users/bservice"
	//"database/sql"
	//"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
)

type AdminUserHandler struct {
	tmpl    *template.Template
	userSrv users.UserService
	balsrv  balance.BalanceService
}

type SessionHandler struct {
	active bool
	user   entity.User
}

var session SessionHandler

func NewAdminUserHandler(t *template.Template, ur users.UserService, b balance.BalanceService) *AdminUserHandler {
	return &AdminUserHandler{tmpl: t, userSrv: ur, balsrv: b}
}

func (userService *AdminUserHandler) Index_handler(w http.ResponseWriter, req *http.Request) {
	userService.tmpl.ExecuteTemplate(w, "index.html", nil)
}
func (userService *AdminUserHandler) MySalesHandler(w http.ResponseWriter, req *http.Request) {
	userService.tmpl.ExecuteTemplate(w, "mySales.html", nil)
}

func (userService *AdminUserHandler) ProfileHandler(w http.ResponseWriter, req *http.Request) {
	type displayItem struct {
		ItemName  string
		ItemPrice float64
	}
	type profileData struct {
		Username        string
		AccountNo       float64
		ShoppingCart    []displayItem
		ShoppingHistory []displayItem
	}
	if session.active {

		balanceAmmount, err := userService.balsrv.Balance(int(session.user.ID))
		if err != nil {
			panic(err)
		}
		//Add shopping cart table

		shoppingCartData := []displayItem{
			displayItem{"Item 1", 200.5},
			displayItem{"Item 2", 1450.6},
		}

		shoppingHistoryData := []displayItem{
			displayItem{"Item 1", 200.5},
			displayItem{"Item 2", 1450.6},
		}

		data := profileData{session.user.FirstName,
			balanceAmmount.YourBalance, shoppingCartData,
			shoppingHistoryData}
		err = userService.tmpl.ExecuteTemplate(w, "profile.html", data)
		if err != nil {
			panic(err)
		}

	} else {
		http.Redirect(w, req, "/", 303)
	}
}

func (userService *AdminUserHandler) Login(w http.ResponseWriter, req *http.Request) {
	userService.tmpl.ExecuteTemplate(w, "signIn.html", nil)
}

func (userService *AdminUserHandler) Signuppage(w http.ResponseWriter, req *http.Request) {
	userService.tmpl.ExecuteTemplate(w, "signup.html", nil)
}

func (userService *AdminUserHandler) AdminRegistration(w http.ResponseWriter, req *http.Request) {
	if session.active {
		http.Redirect(w, req, "/profile", http.StatusSeeOther)
	}
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
	use, err := userService.userSrv.User(usr.Email)
	err = userService.balsrv.StoreId(use.ID)
	balanceEmpty := entity.Balance{}
	balanceEmpty.ID = use.ID
	balanceEmpty.YourBalance = 0
	userService.balsrv.Storeb(int(balanceEmpty.ID), balanceEmpty)
	if err != nil {
		panic(err.Error())
	}
	http.Redirect(w, req, "/signinpage", 303)
	if err != nil {
		panic(err.Error())
	}
}

func (userService *AdminUserHandler) AdminLogin(w http.ResponseWriter, req *http.Request) {
	if session.active {
		http.Redirect(w, req, "/profile", http.StatusSeeOther)
	}
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
		//TEMPORARY
		//WILL SET UP SESSIONS
		session = SessionHandler{true, user}
		http.Redirect(w, req, "/profile", http.StatusSeeOther)
	}
	http.Redirect(w, req, "/signinpage", 303)
	if err != nil {
		panic(err.Error())
	}
}
