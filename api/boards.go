package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/fortytw2/eden/api/util"
	"github.com/fortytw2/eden/datastore"
	"github.com/fortytw2/eden/model"
	"github.com/julienschmidt/httprouter"
)

// GetBoards returns all boards
func GetBoards(ds *datastore.Datastore) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		boards, err := ds.GetBoards(0)
		if err != nil {
			util.JSONError(w, err, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(boards)
	}
}

// GetBoardPosts returns all the posts on a certain board
func GetBoardPosts(ds *datastore.Datastore) httprouter.Handle {
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

type newBoardData struct {
	Name    string
	Summary string
}

// CreateBoard creates a new board
func CreateBoard(ds *datastore.Datastore) httprouter.Handle {
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

		var nbd newBoardData
		err = json.Unmarshal(body, &nbd)
		if err != nil {
			util.JSONError(w, err, http.StatusBadRequest)
			return
		}

		b := &model.Board{
			Name:    nbd.Name,
			Summary: nbd.Summary,
			Creator: user.Username,
		}

		err = ds.CreateBoard(b)
		if err != nil {
			util.JSONError(w, err, http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"success": "board submitted for approval"})
	}
}
