package datastore

import "github.com/fortytw2/eden/model"

// A Datastore holds all storage services -
// abstracting like this allows us to not be tied to any single database backend
// no authorization/authentication is handled by the datastore
type Datastore struct {
	UserService
	BoardService
	PostService
	CommentService
}

// UserService provides a wrapper around user persistance functions
type UserService interface {
	CreateUser(*model.User) error
	UpdateUser(*model.User) error
	GetUserByID(id int) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
}

// BoardService handles CRUD of boards
type BoardService interface {
	CreateBoard(*model.Board) error
	UpdateBoard(*model.Board) error

	GetBoardByName(boardName string) (*model.Board, error)
	GetBoardByID(boardID int) (*model.Board, error)

	DeleteBoard(boardID int) error
}

// PostService abstracts away the complexities of dealing with posts
type PostService interface {
	CreatePost(*model.Post) error

	GetPost(postID int) (*model.Post, error)

	GetBoardPostsByName(boardName string, postFilter *model.PostFilter) ([]model.Post, error)
	GetBoardPostsByID(boardID string, postFilter *model.PostFilter) ([]model.Post, error)
	GetUserPosts(userID int) ([]model.Post, error)

	DeletePost(postID int) error
}

// A CommentService provides access to comments for posts and creation of posts
type CommentService interface {
	CreateRootComment(postID int, newComment *model.Comment) error
	CreateChildComment(commentID int, newComment *model.Comment) error

	GetComment(commentID int) (*model.Comment, error)

	GetPostComments(postID int) ([]model.Comment, error)
	GetUserComments(userID int) ([]model.Comment, error)

	DeleteComment(commentID int) error
}

// GetPostWithComments returns everything needed to present a post to a user
func (d *Datastore) GetPostWithComments(postID int) (post *model.Post, comments []model.Comment, err error) {
	post, err = d.GetPost(postID)
	if err != nil {
		return
	}

	comments, err = d.GetPostComments(postID)
	if err != nil {
		return
	}

	return
}

// GetCommentWithContext returns the specific context of a single comment
func (d *Datastore) GetCommentWithContext(commentID int) (post *model.Post, comment *model.Comment, err error) {
	comment, err = d.GetComment(commentID)
	if err != nil {
		return
	}

	post, err = d.GetPost(comment.PostID)
	if err != nil {
		return
	}

	return
}
