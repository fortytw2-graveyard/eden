package redis

import (
	"errors"
	"strconv"

	"github.com/fortytw2/eden/model"
	redigo "github.com/garyburd/redigo/redis"
)

var (
	ErrUserDidNotVote = errors.New("user did not vote")
)

// VoteService handles setting and getting of votes
type VoteService struct {
	pool *redigo.Pool
}

// NewVoteService returns a VoteService set to use the internal redis pool
func NewVoteService(rpl *redigo.Pool) *VoteService {
	return &VoteService{
		pool: rpl,
	}
}

// GetRealPostVotes returns the total real votes on a post
func (vs *VoteService) GetRealPostVotes(postID int) (total int, err error) {
	total, err = redigo.Int(vs.pool.Get().Do("GET", "post:"+strconv.Itoa(postID)))
	if err != nil {
		return 0, nil
	}
	return
}

// GetRealCommentVotes returns the total real votes on a comment
func (vs *VoteService) GetRealCommentVotes(commentID int) (total int, err error) {
	total, err = redigo.Int(vs.pool.Get().Do("GET", "comment:"+strconv.Itoa(commentID)))
	if err != nil {
		return 0, nil
	}
	return
}

// CheckUserPostVote returns whether or not a user has voted on a post - and the details if they have
func (vs *VoteService) CheckUserPostVote(userID, postID int) (vote *model.Vote, err error) {
	return vs.checkVoteStatus("post", userID, postID)
}

// CheckUserCommentVote returns whether or not a user has voted on a comment - and the details if they have
func (vs *VoteService) CheckUserCommentVote(userID, postID int) (vote *model.Vote, err error) {
	return vs.checkVoteStatus("comment", userID, postID)
}

// SaveVote commits a vote + updates totals based off the internals of the vote
func (vs *VoteService) SaveVote(v *model.Vote) (err error) {
	c := vs.pool.Get()
	if v.Up {
		_, err = c.Do("INCR", v.VoteType+":"+strconv.Itoa(v.ObjectID))
		if err != nil {
			return
		}
		_, err = c.Do("SET", v.VoteType+":"+strconv.Itoa(v.ObjectID)+":"+strconv.Itoa(v.UserID), 1)
	} else {
		_, err = vs.pool.Get().Do("DECR", v.VoteType+":"+strconv.Itoa(v.ObjectID))
		if err != nil {
			return
		}
		_, err = c.Do("SET", v.VoteType+":"+strconv.Itoa(v.ObjectID)+":"+strconv.Itoa(v.UserID), -1)
	}
	return
}

// UpdateVote allows a user to change their vote on a post/comment
func (vs *VoteService) UpdateVote(v *model.Vote) (err error) {
	c := vs.pool.Get()
	var status int
	status, err = redigo.Int(c.Do("GET", v.VoteType+":"+strconv.Itoa(v.ObjectID)+":"+strconv.Itoa(v.UserID)))
	if err != nil {
		return
	}

	if status == -1 && v.Up {
		err = vs.SaveVote(v)
	} else if status == 1 && !v.Up {
		err = vs.SaveVote(v)
	}
	return
}

func (vs *VoteService) checkVoteStatus(voteType string, userID, postID int) (vote *model.Vote, err error) {
	c := vs.pool.Get()
	var status int
	status, err = redigo.Int(c.Do("GET", voteType+":"+strconv.Itoa(postID)+":"+strconv.Itoa(userID)))
	if err != nil {
		return
	}

	if status == 1 {
		vote = model.NewVote(voteType, postID, userID, true)
		return
	}

	if status == 0 {
		vote = nil
		err = ErrUserDidNotVote
		return
	}

	if status == -1 {
		vote = model.NewVote(voteType, postID, userID, false)
		return
	}
	return
}
