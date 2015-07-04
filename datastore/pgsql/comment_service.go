package pgsql

import (
	"github.com/fortytw2/eden/model"
	"github.com/jmoiron/sqlx"
)

// CommentService is a CommentService backed by Postgres
type CommentService struct {
	db *sqlx.DB
}

// NewCommentService returns a new CommentService from the db handle
func NewCommentService(db *sqlx.DB) *CommentService {
	return &CommentService{
		db: db,
	}
}

// CreateRootComment creates a top level comment on a post
func (cs *CommentService) CreateRootComment(postID int, newComment *model.Comment) (err error) {
	return
}

// CreateChildComment creates a comment on another comment
func (cs *CommentService) CreateChildComment(commentID int, newComment *model.Comment) (err error) {
	return
}

// GetComment returns a comment along with its full tree
func (cs *CommentService) GetComment(commentID int) (c *model.Comment, err error) {
	return
}

// GetPostComments returns all root level comments along with their children for a Post`
func (cs *CommentService) GetPostComments(postID int) (comments []*model.Comment, err error) {
	return
}

// GetUserComments returns all of a users past comments
func (cs *CommentService) GetUserComments(userID int) (comments []*model.Comment, err error) {
	return
}

// DeleteComment removes a comment
func (cs *CommentService) DeleteComment(commentID int) (err error) {
	return
}
