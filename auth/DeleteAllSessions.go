package auth

import (
	"github.com/Baxxu/site-donate-back/accesstoken"
	"github.com/Baxxu/site-donate-back/refreshtoken"
	"log"
	"net/http"
)

func DeleteAllSessions(writer http.ResponseWriter, request *http.Request) {
	refreshToken, err := refreshtoken.ReadFromCookie(request)
	if err != nil {
		log.Printf("%s", err)
		return
	}

	refreshtoken.DeleteAll(refreshToken)
	refreshtoken.DeleteFromCookie(writer)

	accesstoken.DeleteFromCookie(writer)

	writer.Header().Set(`Content-Type`, `application/json`)
	writer.Write([]byte(`{"ok":true}`))
}
