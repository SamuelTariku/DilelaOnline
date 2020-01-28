package handler

import (
	"../../../comment"
	"../../../entity"
	"../../../product"
	"../../../users"
	"html/template"
	"net/http"
	"strconv"
)

type AdminCommentHandler struct {
	tmpl       *template.Template
	productSrv product.ProductService
	commentSrv comment.CommentService
	userSrv    users.UserService
}

func NewCommentHandler(t *template.Template, s product.ProductService, c comment.CommentService, u users.UserService) *AdminCommentHandler {
	return &AdminCommentHandler{tmpl: t, productSrv: s, commentSrv: c, userSrv: u}
}

func (commentSrv *AdminCommentHandler) AddComment(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		ids, err := strconv.Atoi(req.URL.Query().Get("product"))
		if err != nil {
			panic(err)
		}
		newComment := entity.Comment{}
		newComment.ProductID = uint(ids)
		newComment.Message = req.FormValue("cmessage")
		rat, err := strconv.Atoi(req.FormValue("crating"))
		if err != nil {
			panic(err)
		}
		newComment.Rating = uint(rat)

		if(OldSession.active){
			newComment.Name = OldSession.user.FirstName  //Anonymous if not logged in
			newComment.UserID = OldSession.user.ID
			newComment.Email = OldSession.user.Email
		} else {
			newComment.Name = "Anonymous"  //Anonymous if not logged in
			newComment.Email = " "
		}

		err = commentSrv.commentSrv.StoreComment(newComment)
		if err != nil {
			panic(err)
		}
		urlMod := "/product?id=" + strconv.Itoa(ids)
		http.Redirect(w, req, urlMod, http.StatusSeeOther)
	}
	return
}
