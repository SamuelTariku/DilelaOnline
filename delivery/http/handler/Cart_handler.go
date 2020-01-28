package handler

import (
	"../../../cartserv"
	"../../../entity"
	"../../../product"
	"html/template"
	"net/http"
	"strconv"
)

type CartHandler struct{
	tmpl *template.Template
	car  cartserv.CartService
	productSrv product.ProductService
}

func NewCartHandler(t *template.Template, c cartserv.CartService, p product.ProductService) *CartHandler{
	return &CartHandler{tmpl: t, car:c, productSrv:p}
}


func (c *CartHandler) GetCart(w http.ResponseWriter, r *http.Request)([]entity.Cart){
	if r.Method == http.MethodGet{
		cid := r.URL.Query().Get("userid")
		userid, err := strconv.Atoi(cid)
		if err != nil{
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
		cart, err := c.car.UserCart(userid)
		if err != nil{
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return cart
	} else {
		http.Redirect(w, r, "/error", 404)
	}
	return nil
}

func (c *CartHandler) AddCart(w http.ResponseWriter, r *http.Request){
	if !OldSession.active {
		http.Redirect(w, r, "/signin", http.StatusSeeOther)
		return
	}
	if r.Method == http.MethodGet{
		pid := r.URL.Query().Get("proid")
		pnm := r.URL.Query().Get("productname")
		userid := OldSession.user.ID
		cart := entity.Cart{}
		proid, err := strconv.Atoi(pid)
		if err != nil{
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		cartProduct, err := c.productSrv.Product(proid)

		if err != nil {
			panic(err)
		}

		cart.ProductID = uint(proid)
		cart.UserID = userid
		cart.Price = cartProduct.Price
		cart.ProductName = pnm

		err = c.car.StoreC(cart)
		if err != nil{
			panic(err)
		}
	}
	http.Redirect(w, r, "/profile", http.StatusSeeOther)
}