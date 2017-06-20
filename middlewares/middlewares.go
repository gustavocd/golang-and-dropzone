package middlewares

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// IsMethodPost checks if the request was sent by http method post.
func IsMethodPost(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		h(w, r, p)
	}
}
