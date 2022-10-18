package accesstoken

import (
	"net/http"
)

func ReadFromCookie(request *http.Request) (token string, err error) {
	cookie, err := request.Cookie(cookieName)
	if err != nil {
		return
	}

	return cookie.Value, nil
}
