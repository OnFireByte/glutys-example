package reqcontext

import (
	"errors"
	"net/http"
)

// do not use type alias here, it will cause error in generated code. Use type definition only.
type Username string

func ParseUsername(r *http.Request) (Username, error) {
	user := r.Header.Get("username")
	if user == "" {
		return "", errors.New("no username in request header")
	}
	return Username(user), nil
}
