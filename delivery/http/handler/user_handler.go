package handler

import (
	"../../../form"
	"../../../SoldProduct"
	"../../../balance"
	"../../../cartserv"
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
	cartSrv cartserv.CartService
	sproductSrv SoldProduct.ProductService
}



func NewAdminUserHandler(t *template.Template, ur users.UserService, b balance.BalanceService, crt cartserv.CartService, sp SoldProduct.ProductService) *AdminUserHandler {
	return &AdminUserHandler{tmpl: t, userSrv: ur, balsrv: b, cartSrv:crt, sproductSrv:sp}
}

func (userService *AdminUserHandler) MySalesHandler(w http.ResponseWriter, req *http.Request) {
	/*
		if req.Method == http.MethodPost {
			addName := req.URL.Query().Get("addName")
			addPrice := req.URL.Query().Get("addPrice")

		}
	*/

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
		ShoppingCart    []entity.Cart
		ShoppingHistory []entity.Product
		Token			string
	}

	var token string

	if req.Method == http.MethodGet{
		token = req.URL.Query().Get("token")
	} else {
		token = " "
	}




	if OldSession.active {

		balanceAmmount, err := userService.balsrv.Balance(int(OldSession.user.ID))
		if err != nil {
			panic(err)
		}

		shoppingCartData, err := userService.cartSrv.UserCart(int(OldSession.user.ID))
		if err != nil {
			panic(err)
		}

		shoppingHistoryData,err := userService.sproductSrv.SoldP(int(OldSession.user.ID))
		if err != nil {
			panic(err)
		}



		data := profileData{OldSession.user.FirstName,
			balanceAmmount.YourBalance, shoppingCartData,
			shoppingHistoryData, token}

		err = userService.tmpl.ExecuteTemplate(w, "profile.html", data)
		if err != nil {
			panic(err)
		}

	} else {
		http.Redirect(w, req, "/", 303)
	}
}

func (userService *AdminUserHandler) Login(w http.ResponseWriter, req *http.Request) {

	validator := form.Input{Values:req.PostForm, VErrors: form.ValidationErrors{}}
	userService.tmpl.ExecuteTemplate(w, "signIn.html", validator)
}

func (userService *AdminUserHandler) Logout(w http.ResponseWriter, req *http.Request) {
	SessionStop()
	http.Redirect(w, req, "/", 303)
}
func (userService *AdminUserHandler) ErrorPage(w http.ResponseWriter, req *http.Request) {
	i := "<!DOCTYPE html><html><head></head><body>ERROR 404</body>"
	t := template.New("")
	t, _ = t.Parse(i)
	t.Execute(w, nil)
}
func (userService *AdminUserHandler) Signuppage(w http.ResponseWriter, req *http.Request) {
	validator := form.Input{Values:req.PostForm, VErrors: form.ValidationErrors{}}
	userService.tmpl.ExecuteTemplate(w, "signup.html", validator)
}

func (userService *AdminUserHandler) AdminRegistration(w http.ResponseWriter, req *http.Request) {
	if OldSession.active {
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

	validator_signin := form.Input{Values:req.PostForm, VErrors:form.ValidationErrors{}}
	validator_signin.Required("first", "last", "email", "password", "secondpassword")
	validator_signin.MatchesPattern("email", form.EmailRX)
	validator_signin.MinLength("password", 8)
	validator_signin.PasswordMatches("password", "secondpassword")

	if !validator_signin.Valid() {
		userService.tmpl.ExecuteTemplate(w, "signup.html", validator_signin)
	}
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
	if OldSession.active {
		http.Redirect(w, req, "/profile", http.StatusSeeOther)
	}
	if req.Method != "POST" {
		http.Redirect(w, req, "/signinpage", http.StatusSeeOther)
		return
	}

	validator := form.Input{Values:req.PostForm, VErrors: form.ValidationErrors{}}
	email := req.FormValue("email")
	password := req.FormValue("password")

	log.Println(email)
	usr, err := userService.userSrv.User(email)

	if err != nil {
		validator.VErrors.Add("Login fail", "Username or password incorrect")
		userService.tmpl.ExecuteTemplate(w, "signIn.html", validator)
		return
	}


	if email == usr.Email {
		err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(password))
		log.Println("Reached here")
		if err != nil {
			if err != nil {
				validator.VErrors.Add("Login fail", "Username or password incorrect")
				err = userService.tmpl.ExecuteTemplate(w, "signIn.html", validator)
				if err != nil{
					panic(nil)
				}
				return
			}
		}
		user, err := userService.userSrv.UserwithID(int(usr.ID))
		if err != nil {
			panic(err.Error())
		}
		//TEMPORARY
		//WILL SET UP SESSIONS
		//Profile generated a random token and stores it in database
		SessionStart(user)
		//Redirect to profile with encrypted token
		http.Redirect(w, req, "profile.html", http.StatusSeeOther)
	}
	http.Redirect(w, req, "/signinpage", 303)
	if err != nil {
		panic(err.Error())
	}
}
