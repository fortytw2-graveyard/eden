package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/fortytw2/eden/api/util"
	"github.com/fortytw2/eden/datastore"
	"github.com/fortytw2/eden/model"
	"github.com/julienschmidt/httprouter"
)

type newCommentData struct {
	ParentID int `json:"parent_id"`
	Body     string
}

// CreateComment creates a new post
func CreateComment(ds *datastore.Datastore) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		user, err := util.Authenticate(r, ds)
		if err != nil {
			util.JSONError(w, err, http.StatusUnauthorized)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			util.JSONError(w, err, http.StatusBadRequest)
			return
		}

		var nbd newCommentData
		err = json.Unmarshal(body, &nbd)
		if err != nil {
			util.JSONError(w, err, http.StatusBadRequest)
			return
		}

		num, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			util.JSONError(w, err, http.StatusBadRequest)
			return
		}

		p := &model.Comment{
			UserID:    user.ID,
			PostID:    num,
			CommentID: nbd.ParentID,
			Body:      nbd.Body,
		}

		err = ds.CreateComment(p)
		if err != nil {
			util.JSONError(w, err, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"success": "comment submitted"})
	}
}
