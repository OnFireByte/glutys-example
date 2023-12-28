package reqcontext

import (
	"errors"
	"net/http"
)

type Username string

func ParseUsername(r *http.Request) (Username, error) {
	user := r.Header.Get("username")
	if user == "" {
		return "", errors.New("no username in request header")
	}
	return Username(user), nil
}
