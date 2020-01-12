package service

import (
	"../../comment"
	"../../entity"
)

// CommentService implements menu.CommentService interface
type CommentService struct {
	commentRepo comment.CommentRepository
}

// NewCommentService returns a new CommentService object
func NewCommentService(commRepo comment.CommentRepository) *CommentService {
	return &CommentService{commentRepo: commRepo}
}

// Comments returns all stored comments
func (cs *CommentService) Comments() ([]entity.Comment, error) {
	cmnts, err := cs.commentRepo.Comments()
	if err != nil {
		return nil, err
	}
	return cmnts, nil
}

// Comment retrieves stored comment by its id
func (cs *CommentService) Comment(id uint) (entity.Comment, error) {
	cmnt, errs := cs.commentRepo.Comment(id)
	if errs != nil {
		return cmnt, errs
	}
	return cmnt, nil
}

// UpdateComment updates a given comment
func (cs *CommentService) UpdateComment(comment entity.Comment) error {
	errs := cs.commentRepo.UpdateComment(comment)
	if errs != nil {
		return errs
	}
	return nil
}

// DeleteComment deletes a given comment
func (cs *CommentService) DeleteComment(id uint) error {
	errs := cs.commentRepo.DeleteComment(id)
	if errs != nil {
		return errs
	}
	return nil
}

// StoreComment stores a given comment
func (cs *CommentService) StoreComment(comment entity.Comment) error {
	errs := cs.commentRepo.StoreComment(comment)
	if errs != nil {
		return errs
	}
	return nil
}
