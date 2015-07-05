package model

import "time"

// Comment is either a root or child comment on a post
type Comment struct {
	ID        int        `json:"id"`
	UserID    int        `json:"-" db:"op_id"`
	PostID    int        `json:"-" db:"post_id"`
	CommentID int        `json:"-" db:"comment_id"`
	Votes     int        `json:"votes" db:"-"`
	Body      string     `json:"body"`
	Children  []*Comment `json:"children,omitempty"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
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
