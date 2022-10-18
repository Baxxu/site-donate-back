package accesstoken

import (
	"github.com/Baxxu/site-donate-back/consts"
	"net/http"
)

func WriteToCookie(token string, writer http.ResponseWriter) {
	//срок жизни - день
	cookie := http.Cookie{
		Name:     cookieName,
		Value:    token,
		Path:     consts.ApiPath,
		Domain:   "testest.ru",
		MaxAge:   cookieMaxAge,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(writer, &cookie)
}
