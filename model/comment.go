package model

// Comment is either a root or child comment on a post
type Comment struct {
	ID        int
	PostID    int
	CommentID int
	Votes     int `json:"votes" db:"-"`
	Body      string
	children  []*Comment
}

// NewComment generates a comment on a comment
func NewComment(postID int, commentID int, body string) *Comment {
	return nil
}
