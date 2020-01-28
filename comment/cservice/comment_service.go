package cservice

import (
	"../../comment"
	"../../entity"
)

// CommentService implements menu.CommentService interface
type CommentService struct {
	commentRepo comment.CommentRepository
}

// NewCommentService returns a new CommentService object
func NewCommentService(commRepo comment.CommentRepository) comment.CommentService {
	return &CommentService{commentRepo: commRepo}
}

// Comments returns all stored comments
func (cs *CommentService) Comments() ([]entity.Comment, error) {
	cmnts, errs := cs.commentRepo.Comments()
	if errs != nil {
		return nil, errs
	}
	return cmnts, errs
}

// Comment retrieves stored comment by its id
func (cs *CommentService) Comment(id int) (entity.Comment, error) {
	cmnt, errs := cs.commentRepo.Comment(id)
	if errs != nil {
		return cmnt, errs
	}
	return cmnt, errs
}

func (cs *CommentService) ProductComment(productid int) ([]entity.Comment, error) {
	cmnts, errs := cs.commentRepo.ProductComment(productid)
	if errs != nil {
		return nil, errs
	}
	return cmnts, errs
}

// UpdateComment updates a given comment
func (cs *CommentService) UpdateComment(comment entity.Comment) error {
	errs := cs.commentRepo.UpdateComment(comment)
	return errs
}

// DeleteComment deletes a given comment
func (cs *CommentService) DeleteComment(id int) error {
	errs := cs.commentRepo.DeleteComment(id)
	return errs
}

// StoreComment stores a given comment
func (cs *CommentService) StoreComment(comment entity.Comment) error {
	errs := cs.commentRepo.StoreComment(comment)
	return errs
}
