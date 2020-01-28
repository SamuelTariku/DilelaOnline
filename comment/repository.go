package comment

import "../entity"

// CommentRepository specifies customer comment related database operations
type CommentRepository interface {
	Comments() ([]entity.Comment, error)
	Comment(id int) (entity.Comment, error)
	ProductComment(productid int) ([]entity.Comment, error)
	UpdateComment(comment entity.Comment) error
	DeleteComment(id int) error
	StoreComment(comment entity.Comment) error
}
