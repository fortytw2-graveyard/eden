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

type registerData struct {
	Username, Email, Password string
}

// NewUser creates a user
func NewUser(ds *datastore.Datastore) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			util.JSONError(w, err, http.StatusBadRequest)
			return
		}

		var rd registerData
		err = json.Unmarshal(body, &rd)
		if err != nil {
			util.JSONError(w, err, http.StatusBadRequest)
			return
		}

		user, err := model.CreateUser(rd.Username, rd.Email, rd.Password)
		if err != nil {
			util.JSONError(w, err, http.StatusInternalServerError)
			return
		}

		err = ds.CreateUser(user)
		if err != nil {
			util.JSONError(w, err, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "user created successfully"})
	}
}
