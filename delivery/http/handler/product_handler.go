package handler

import (
	"../../../advertisement"
	"../../../comment"
	"../../../entity"
	"../../../product"
	"html/template"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type AdminProductHandler struct {
	tmpl       *template.Template
	productSrv product.ProductService
	commentSrv comment.CommentService
	advertSrv  advertisement.AdvertService
}

func NewAdminSellerHandler(t *template.Template, s product.ProductService, c comment.CommentService, a advertisement.AdvertService) *AdminProductHandler {
	return &AdminProductHandler{tmpl: t, productSrv: s, commentSrv: c, advertSrv: a}
}

func (productSrv *AdminProductHandler) ProductPage(w http.ResponseWriter, req *http.Request) {

	type productDisplay struct {
		AverageStars int
		Product      entity.Product
		Comments     []entity.Comment
	}
	if req.Method == http.MethodGet {
		id := req.URL.Query().Get("id")
		if len(id) < 1 {
			http.Redirect(w, req, "/error", 404)
			return
		}
		rID, err := strconv.Atoi(id)
		if err != nil {
			panic(err)
		}

		product, err := productSrv.productSrv.Product(rID)
		if err != nil {
			panic(err)
		}
		comments, err := productSrv.commentSrv.ProductComment(rID)
		if err != nil {
			panic(err)
		}
		average := 0
		if len(comments) > 0 {
			sum := 0
			for com := range comments {
				sum += int(comments[com].Rating)
			}
			log.Println(sum)
			average = sum / len(comments)
		} else {
			average = 0
		}
		data := productDisplay{average, product,
			comments}
		productSrv.tmpl.ExecuteTemplate(w, "newProduct.html", data)
	} else {
		http.Redirect(w, req, "/error", 404)
	}

}

func (productSrv *AdminProductHandler) Index_handler(w http.ResponseWriter, req *http.Request) {

	products, err := productSrv.productSrv.Products()
	if err != nil {
		panic(err)
	}
	adverts, err := productSrv.advertSrv.Adverts()
	if err != nil {
		panic(err)
	}
	type ProductData struct {
		Image string
		Link  string
		Price string
	}
	productHouse := []ProductData{}
	productCars := []ProductData{}
	productElect := []ProductData{}
	productGoods := []ProductData{}
	house := true
	cars := true
	elect := true
	goods := true

	for i := 0; i < len(products); i = i + 1 {

		if products[i].Ptype == "House" && house {
			productHouse = append(productHouse, ProductData{
				products[i].Image,
				strconv.Itoa(int(products[i].ID)),
				strconv.FormatFloat(products[i].Price, 'f', 2, 64)})
		} else if products[i].Ptype == "Cars" && cars {
			productCars = append(productCars, ProductData{
				products[i].Image,
				strconv.Itoa(int(products[i].ID)),
				strconv.FormatFloat(products[i].Price, 'f', 2, 64)})
		} else if products[i].Ptype == "Electronics" && elect {
			productElect = append(productElect, ProductData{
				products[i].Image,
				strconv.Itoa(int(products[i].ID)),
				strconv.FormatFloat(products[i].Price, 'f', 2, 64)})
		} else if goods {
			productGoods = append(productGoods, ProductData{
				products[i].Image,
				strconv.Itoa(int(products[i].ID)),
				strconv.FormatFloat(products[i].Price, 'f', 2, 64)})
		}

		if len(productHouse) > 4 {
			house = false
		} else if len(productElect) > 4 {
			elect = false
		} else if len(productCars) > 4 {
			cars = false
		} else if len(productGoods) > 4 {
			goods = false
		}
	}
	//HOTFIX FOR CAROUSEL
	adsProduct := []entity.Product{}
	for i := 0; i < len(adverts); i = i + 1 {
		advert, err := productSrv.productSrv.Product(int(adverts[i].ProductID))
		if err != nil {
			panic(err)
		}
		adsProduct = append(adsProduct, advert)
	}
	if len(adsProduct) < 3 {
		adsProduct = append(adsProduct, entity.Product{Image: "img1.jpg"})
		if len(adsProduct) < 3 {
			adsProduct = append(adsProduct, entity.Product{Image: "img1.jpg"})
			if len(adsProduct) < 3 {
				adsProduct = append(adsProduct, entity.Product{Image: "img1.jpg"})
			}
		}
	}

	ads := struct {
		FirstProduct  entity.Product
		SecondProduct entity.Product
		ThirdProduct  entity.Product
	}{
		adsProduct[0],
		adsProduct[1],
		adsProduct[2]}
	//---------------------
	data := struct {
		SliderAdvert struct {
			FirstProduct  entity.Product
			SecondProduct entity.Product
			ThirdProduct  entity.Product
		}
		HouseData []ProductData
		ElectData []ProductData
		CarsData  []ProductData
		GoodsData []ProductData
	}{
		ads,
		productHouse,
		productElect,
		productCars,
		productGoods}

	err = productSrv.tmpl.ExecuteTemplate(w, "index.html", data)
	if err != nil{
		panic(err)
	}
}

func writeFile(mf *multipart.File, fname string) {

	wd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	path := filepath.Join(wd, "../", "../", "ui", "assets", "images", fname)
	image, err := os.Create(path)

	if err != nil {
		panic(err)
	}
	defer image.Close()
	io.Copy(image, *mf)
}

func (pserv *AdminProductHandler) SellerProducts(w http.ResponseWriter, r *http.Request) {
	products, errs := pserv.productSrv.Products()
	if errs != nil {
		panic(errs)
	}
	err := pserv.tmpl.ExecuteTemplate(w, "#", products)
	if err != nil {
		panic(err.Error())
	}
}

func (pserv *AdminProductHandler) NewSellerProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		prod := entity.Product{}
		prod.Name = r.FormValue("name")
		prod.Ptype = r.FormValue("type")
		prod.Price, _ = strconv.ParseFloat(r.FormValue("price"), 64)
		prod.Description = r.FormValue("description")

		mf, fh, err := r.FormFile("img")
		if err != nil {
			panic(err)
		}
		defer mf.Close()

		prod.Image = fh.Filename
		prod.UserID = OldSession.user.ID

		writeFile(&mf, fh.Filename)
		err = pserv.productSrv.StoreP(prod)
		if err != nil {
			panic(err)
		}
		http.Redirect(w, r, "/mySales", http.StatusSeeOther)

	} else {
		http.Redirect(w, r, "/mySales", http.StatusSeeOther)
	}
}

func (pserv *AdminProductHandler) UpdateSellerProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ids := r.URL.Query().Get("id")
		id, err := strconv.Atoi(ids)

		if err != nil {
			panic(err)
		}
		pro, err := pserv.productSrv.Product(id)

		if err != nil {
			panic(err)
		}
		pserv.tmpl.ExecuteTemplate(w, "#", pro)
	} else if r.Method == http.MethodPost {
		prod := entity.Product{}
		prod.Name = r.FormValue("name")
		prod.Ptype = r.FormValue("type")
		prod.Price, _ = strconv.ParseFloat(r.FormValue("price"), 64)
		prod.Description = r.FormValue("description")
		prod.Image = r.FormValue("image")
		mf, _, err := r.FormFile("img")

		if err != nil {
			panic(err)
		}
		defer mf.Close()
		writeFile(&mf, prod.Image)
		err = pserv.productSrv.UpdateP(prod)
		if err != nil {
			panic(err.Error())
		}
		http.Redirect(w, r, "#", http.StatusSeeOther)

	} else {
		http.Redirect(w, r, "#", http.StatusSeeOther)
	}
}
func (pserv *AdminProductHandler) DeleteSellerProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ids := r.URL.Query().Get("id")
		id, err := strconv.Atoi(ids)
		if err != nil {
			panic(err)
		}
		err = pserv.productSrv.DeleteP(id)
		if err != nil {
			panic(err.Error())
		}
	}
	http.Redirect(w, r, "#", http.StatusSeeOther)
}
/*func (pserv *AdminProductHandler) BrowseHouse(w http.ResponseWriter, r *http.Request) {
	pserv.productSrv.Bytype("")
	pserv.tmpl.ExecuteTemplate(w, "houses.html", data)

}*/
func (pserv *AdminProductHandler) SearchProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ids := r.URL.Query().Get("search")
		if len(ids) == 0 {
			http.Redirect(w, r, "/error", 404)
		}
		results, err := pserv.productSrv.SearchProduct(ids)
		if err != nil {
			panic(err)
		}
		pserv.tmpl.ExecuteTemplate(w, "search.html", results)
	} else {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
	}
}

func (pserv *AdminProductHandler) SearchProduct(prod string) ([]entity.Product, error) {
	product, err := pserv.productSrv.SearchProduct(prod)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (pserv *AdminProductHandler) ProductDetail(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		ids := r.URL.Query().Get("id")
		id, err := strconv.Atoi(ids)

		if err != nil {
			panic(err)
		}

		pro, err := pserv.productSrv.Product(id)

		if err != nil {
			panic(err.Error)
		}

		_ = pserv.tmpl.ExecuteTemplate(w, "productdetail.layout", pro)
	}
}
