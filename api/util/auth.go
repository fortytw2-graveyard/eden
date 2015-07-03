package util

import (
	"errors"
	"net/http"

	"github.com/fortytw2/abdi"
	"github.com/fortytw2/eden/datastore"
	"github.com/fortytw2/eden/model"
)

// Authenticate checks for correct credentials via HTTP basic auth
func Authenticate(r *http.Request, users datastore.UserService) (u *model.User, err error) {
	username, password, ok := r.BasicAuth()
	if !ok {
		err = errors.New("basic auth not OK")
		return
	}
	u, err = users.GetUserByUsername(username)
	if err != nil {
		err = errors.New("no user found")
		return
	}
	if err = abdi.Check(password, u.PasswordHash); err != nil {
		return
	}

	return
}
