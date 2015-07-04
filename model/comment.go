package model

// Comment is either a root or child comment on a post
type Comment struct {
	ID        int
	UserID    int `db:"op_id"`
	PostID    int `db:"post_id"`
	CommentID int `db:"comment_id"`
	Votes     int `json:"votes" db:"-"`
	Body      string
	Children  []*Comment
}

// NewComment generates a comment on a comment
func NewComment(postID, commentID, userID int, body string) *Comment {
	return &Comment{
		UserID:    userID,
		PostID:    postID,
		CommentID: commentID,
		Body:      body,
	}
}
