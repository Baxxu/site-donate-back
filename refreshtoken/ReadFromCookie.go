package refreshtoken

import (
	"encoding/hex"
	"log"
	"net/http"
)

func ReadFromCookie(request *http.Request) (token []byte, err error) {
	cookie, err := request.Cookie(cookieName)
	if err != nil {
		return nil, err
	}

	if len(cookie.Value) != 64 {
		return nil, ErrInvalid
	}

	token, err = hex.DecodeString(cookie.Value)
	if err != nil {
		log.Printf("%s", err)
		return
	}

	return token, nil
}
