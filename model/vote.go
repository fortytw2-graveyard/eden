package model

import "github.com/satori/go.uuid"

// A Vote is a single vote on a post or comment
type Vote struct {
	UUID     string
	VoteType string
	ObjectID int
	UserID   int
	Up       bool // if false, down
}

// NewVote constructs a new Vote from the given params
func NewVote(voteType string, objectID, user int, up bool) *Vote {
	if voteType != "post" && voteType != "comment" {
		return nil
	}

	return &Vote{
		UUID:     uuid.NewV4().String(),
		VoteType: voteType,
		ObjectID: objectID,
		UserID:   user,
		Up:       up,
	}
}
