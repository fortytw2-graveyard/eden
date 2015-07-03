package web

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Homepage renders the non-JSON homepage
func Homepage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Write([]byte("HI"))

	return
}
