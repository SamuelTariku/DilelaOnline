package handler

import (
	"../../../entity"
	"../../../product"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type AdminProductHandler struct {
	tmpl       *template.Template
	productSrv product.ProductService
}

func NewAdminSellerHandler(t *template.Template, s product.ProductService) *AdminProductHandler {
	return &AdminProductHandler{tmpl: t, productSrv: s}
}

func writeFile(mf *multipart.File, fname string) {

	wd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	path := filepath.Join(wd, "../", "../", "../", "ui", "assets", "img", fname)
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

		writeFile(&mf, fh.Filename)
		if err != nil {
			panic(err)
		}
		http.Redirect(w, r, "#", http.StatusSeeOther)

	} else {
		err := pserv.tmpl.ExecuteTemplate(w, "#", nil)
		if err != nil {
			panic(err.Error())
		}

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

func (pserv *AdminProductHandler) SearchProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ids := r.URL.Query().Get("search")
		if len(ids) == 0 {
			http.Redirect(w, r, "#", 303)
		}
		results, err := pserv.productSrv.SearchProduct(ids)

		if err != nil {
			panic(err)
		}
		pserv.tmpl.ExecuteTemplate(w, "#", results)
	} else {
		http.Redirect(w, r, "#", http.StatusSeeOther)
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
