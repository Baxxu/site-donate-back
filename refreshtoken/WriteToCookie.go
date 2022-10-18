package refreshtoken

import (
	"encoding/hex"
	"github.com/Baxxu/site-donate-back/consts"
	"net/http"
)

func WriteToCookie(token []byte, writer http.ResponseWriter) {
	tokenStr := hex.EncodeToString(token)

	//срок жизни - год
	cookie := http.Cookie{
		Name:     cookieName,
		Value:    tokenStr,
		Path:     consts.AuthPath,
		Domain:   "testest.ru",
		MaxAge:   cookieMaxAge,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(writer, &cookie)
}
