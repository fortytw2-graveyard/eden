package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fortytw2/eden/api/util"
	"github.com/fortytw2/eden/datastore"
	"github.com/fortytw2/eden/model"
	"github.com/julienschmidt/httprouter"
)

// VoteComment lets a user upvote a comment
func VoteComment(ds *datastore.Datastore) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		user, err := util.Authenticate(r, ds)
		if err != nil {
			util.JSONError(w, err, http.StatusUnauthorized)
			return
		}

		postID, err := strconv.Atoi(ps.ByName("post_id"))
		if err != nil {
			util.JSONError(w, err, http.StatusBadRequest)
			return
		}

		commentID, err := strconv.Atoi(ps.ByName("comment_id"))
		if err != nil {
			util.JSONError(w, err, http.StatusBadRequest)
			return
		}

		voteStatus, err := strconv.ParseBool(ps.ByName("vote"))
		if err != nil {
			util.JSONError(w, err, http.StatusBadRequest)
			return
		}

		comment, err := ds.GetComment(commentID)
		if err != nil {
			util.JSONError(w, err, http.StatusInternalServerError)
			return
		}

		if postID != comment.PostID {
			util.JSONError(w, err, http.StatusBadRequest)
			return
		}

		v, err := ds.CheckUserCommentVote(user.ID, commentID)
		if v == nil {
			err = ds.SaveVote(model.NewVote("comment", commentID, user.ID, voteStatus))
			if err != nil {
				util.JSONError(w, err, http.StatusInternalServerError)
				return
			}
			return
		}

		v.Up = voteStatus
		err = ds.UpdateVote(v)
		if err != nil {
			util.JSONError(w, err, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"success": "upvoted comment"})
	}
}

// UpvotePost makes a new upvote on a post
func UpvotePost(ds *datastore.Datastore) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		posts, err := ds.GetBoardPostsByName(ps.ByName("board"), 0)
		if err != nil {
			util.JSONError(w, err, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(posts)
	}
}
