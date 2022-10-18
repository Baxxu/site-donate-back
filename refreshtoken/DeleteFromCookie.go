package refreshtoken

import (
	"github.com/Baxxu/site-donate-back/consts"
	"net/http"
)

func DeleteFromCookie(writer http.ResponseWriter) (err error) {
	cookie := http.Cookie{
		Name:     cookieName,
		Value:    "DELETED",
		Path:     consts.AuthPath,
		Domain:   "testest.ru",
		MaxAge:   -1,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(writer, &cookie)

	return nil
}
