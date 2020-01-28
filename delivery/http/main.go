package main

import (
	//"../../entity"
	"../../cartserv/crtrepository"
	"../../cartserv/crtservice"
	"../../advertisement/arepository"
	"../../advertisement/aservice"
	"../../balance/brepository"
	"../../balance/bservice"
	"../../comment/crepository"
	"../../comment/cservice"
	"../../product/prepository"
	"../../product/pservice"
	"../../SoldProduct/sprepository"
	"../../SoldProduct/spservice"
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
var productService *pservice.ProductService
var sproductService *spservice.ProductService
var commentService *cservice.CommentService
var advertService *aservice.AdvertService
var cartSerivice *crtservice.CartService
var sessionService *service.SessionService

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

	commentRepo := crepository.NewCommPostRepo(db)
	commentServ := cservice.NewCommentService(commentRepo)

	advertRepo := arepository.NewPostAdvertRepo(db)
	advertServ := aservice.NewAdvertService(advertRepo)

	productRep := prepository.NewPostProductRepo(db)
	productService = pservice.NewProductService(productRep)

	sproductRep := sprepository.NewPostProductRepo(db)
	sproductService = spservice.NewProductService(sproductRep)

	cartRepo := crtrepository.NewPostCartRepo(db)
	cartServ := crtservice.NewCartService(cartRepo)

/*	sessionRepo := repository.NewSessionRepo(db)
	sessionServ := service.NewSessionService(sessionRepo)*/

	adminUserHandler := handler.NewAdminUserHandler(tmpl, userService, balanceService, cartServ, sproductService)
	adminProductHandler := handler.NewAdminSellerHandler(tmpl, productService, commentServ, advertServ)
	adminCommentHandler := handler.NewCommentHandler(tmpl, productService, commentServ, userService)
	adminCartHandler := handler.NewCartHandler(tmpl, cartServ, productService)
	adminOrderHandler := handler.NewAdminOrderHandler(tmpl, productService, cartServ, sproductService, balanceService)

	fs := http.FileServer(http.Dir("../../ui/assets"))

	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/signuppage", adminUserHandler.Signuppage)
	http.HandleFunc("/error", adminUserHandler.ErrorPage)
	http.HandleFunc("/", adminProductHandler.Index_handler)
	http.HandleFunc("/signinpage", adminUserHandler.Login)
	http.HandleFunc("/signup", adminUserHandler.AdminRegistration)
	http.HandleFunc("/signin", adminUserHandler.AdminLogin)
	http.HandleFunc("/profile", adminUserHandler.ProfileHandler)
	http.HandleFunc("/mySales", adminUserHandler.MySalesHandler)
	http.HandleFunc("/newProduct", adminProductHandler.NewSellerProducts)
	http.HandleFunc("/product", adminProductHandler.ProductPage)
	http.HandleFunc("/search", adminProductHandler.SearchProducts)
	http.HandleFunc("/house", adminProductHandler.SearchProducts)
	http.HandleFunc("/electronics", adminProductHandler.SearchProducts)
	http.HandleFunc("/cars", adminProductHandler.SearchProducts)
	http.HandleFunc("/goods", adminProductHandler.SearchProducts)
	http.HandleFunc("/addComment", adminCommentHandler.AddComment)
	http.HandleFunc("/addCart", adminCartHandler.AddCart)
	http.HandleFunc("/order/add", adminOrderHandler.AddOrder)
	http.HandleFunc("/signout", adminUserHandler.Logout)
	http.ListenAndServe(":8080", nil)

}
