package model

// A Vote is a single vote on a post or comment
type Vote struct {
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
		VoteType: voteType,
		ObjectID: objectID,
		UserID:   user,
		Up:       up,
	}
}
