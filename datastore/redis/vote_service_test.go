package redis

import (
	"testing"

	"github.com/fortytw2/eden/model"
)

func TestVoteService(t *testing.T) {
	vs := NewVoteService(GetRedisPool())

	votes := []*model.Vote{
		model.NewVote("post", 1, 1, true),
		model.NewVote("post", 1, 1, true),
		model.NewVote("post", 1, 2, true),
		model.NewVote("post", 1, 3, true),
		model.NewVote("comment", 2, 1, false),
		model.NewVote("comment", 2, 2, false),
		model.NewVote("comment", 2, 3, false),
		model.NewVote("comment", 2, 4, false),
		model.NewVote("comment", 2, 5, false),
	}

	for _, vote := range votes {
		err := vs.SaveVote(vote)
		if err != nil {
			t.Errorf("error saving vote: %s", err)
		}
	}

	commentTotal, err := vs.GetRealCommentVotes(2)
	if err != nil {
		t.Errorf("error getting comment votes, %s", err)
	}

	postTotal, err := vs.GetRealPostVotes(1)
	if err != nil {
		t.Errorf("error getting post votes, %s", err)
	}

	if commentTotal != -5 {
		t.Errorf("commentTotal is not the exact number in GetRealCommentVotes(), %d", commentTotal)
	}

	if postTotal != 4 {
		t.Errorf("postTotal is not the exact number in GetRealPostVotes(), %d", postTotal)
	}

}
