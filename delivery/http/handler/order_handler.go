package handler

import (
	"../../../SoldProduct"
	"../../../balance"
	"../../../cartserv"
	"../../../product"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type AdminOrderHandler struct {
	tmpl       *template.Template
	productSrv product.ProductService
	cartSrv		cartserv.CartService
	sproductSrv SoldProduct.ProductService
	balanceSrv	balance.BalanceService

}

func (h AdminOrderHandler) AddOrder(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet{
		cartID,err := strconv.Atoi(request.URL.Query().Get("id"))
		if err != nil{
			panic(err)
		}
		cart, err := h.cartSrv.Cart(cartID)
		if err != nil{
			panic(err)
		}
		product, err := h.productSrv.Product(int(cart.ProductID))
		if err != nil{
			panic(err)
		}
		err = h.sproductSrv.StoreP(product)
		if err != nil{
			panic(err)
		}

		//CALC
		price := product.Price
		log.Println("PRICE ", price)
		var comm = price * 0.15
		log.Println("CALC comm", comm)
		sellerp := price - comm

		//BUYER BALANCE

		buyer := OldSession.user.ID
		buyerbalance, err := h.balanceSrv.Balance(int(buyer))
		if err != nil{
			panic(err)
		}
		log.Println("BuyerBalance Before:", buyerbalance.YourBalance)
		buyerbalance.YourBalance = buyerbalance.YourBalance - price
		log.Println("BuyerBalance After:", buyerbalance.YourBalance)

		err = h.balanceSrv.Storeb(int(buyer), buyerbalance)
		if err != nil{
			panic(err)
		}

		//SELLER BALANCE
		seller := product.UserID
		sellerbal, err := h.balanceSrv.Balance(int(seller))
		if err != nil{
			panic(err)
		}
		sellerbal.YourBalance = sellerbal.YourBalance + sellerp
		log.Println("SellerBalance After:", sellerbal.YourBalance)
		err = h.balanceSrv.Storeb(int(seller), sellerbal)
		if err != nil{
			panic(err)
		}

		//ADMIN BALANCE
		if err != nil{
			panic(err)
		}
		adminb,err := h.balanceSrv.Balance(int(OldSession.user.ID))
		if err != nil{
			panic(err)
		}
		adminb.YourBalance = adminb.YourBalance + comm
		log.Println("AdminBalance", adminb.YourBalance)
		err = h.balanceSrv.Storeb(int(OldSession.user.ID),adminb)
		if err != nil{
			panic(err)
		}

		//REMOVE PRODUCT FROM CART
		err = h.cartSrv.DeleteC(cartID)
		err = h.productSrv.DeleteP(int(product.ID))
		if err != nil{
			panic(err)
		}

		http.Redirect(writer, request, "/profile", http.StatusSeeOther)

	}
}

func NewAdminOrderHandler(t *template.Template, p product.ProductService, c cartserv.CartService, sp SoldProduct.ProductService, bl balance.BalanceService) *AdminOrderHandler {
	return &AdminOrderHandler{tmpl: t, productSrv: p, cartSrv: c, sproductSrv:sp, balanceSrv:bl}
}

