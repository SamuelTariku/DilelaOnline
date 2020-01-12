package comment

import "../entity"

// CommentRepository specifies customer comment related database operations
type CommentRepository interface {
	Comments() ([]entity.Comment, error)
	Comment(id uint) (entity.Comment, error)
	UpdateComment(comment entity.Comment) error
	DeleteComment(id uint) error
	StoreComment(comment entity.Comment) error
}
