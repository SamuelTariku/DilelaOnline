package comment

import "../entity"

// CommentService specifies customer comment related service
type CommentService interface {
	Comments() ([]entity.Comment, error)
	Comment(id uint) (entity.Comment, error)
	UpdateComment(comment entity.Comment) error
	DeleteComment(id uint) error
	StoreComment(comment entity.Comment) error
}
