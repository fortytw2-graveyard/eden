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

// GetPost returns a single post and its comments
func GetPost(ds *datastore.Datastore) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		num, err := strconv.Atoi(ps.ByName("post_id"))
		if err != nil {
			util.JSONError(w, err, http.StatusInternalServerError)
			return
		}

		post, comments, err := ds.GetPostWithComments(num)
		if err != nil {
			util.JSONError(w, err, http.StatusInternalServerError)
			return
		}

		encoded := struct {
			Post     *model.Post      `json:"post"`
			Comments []*model.Comment `json:"comments,omitempty"`
		}{
			Post:     post,
			Comments: comments,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(encoded)
	}
}

type newPostData struct {
	Title string
	Link  string
	Body  string
}

// CreatePost creates a new post
func CreatePost(ds *datastore.Datastore) httprouter.Handle {
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

		var nbd newPostData
		err = json.Unmarshal(body, &nbd)
		if err != nil {
			util.JSONError(w, err, http.StatusBadRequest)
			return
		}

		board, err := ds.GetBoardByName(ps.ByName("board"))
		if err != nil || board == nil {
			util.JSONError(w, err, http.StatusBadRequest)
			return
		}

		p := &model.Post{
			OpID:  user.ID,
			Board: board.ID,
			Title: nbd.Title,
			Link:  nbd.Link,
			Body:  nbd.Body,
		}

		err = ds.CreatePost(p)
		if err != nil {
			util.JSONError(w, err, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"success": "post submitted"})
	}
}
