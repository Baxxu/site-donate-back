package refreshtoken

import (
	"errors"
	"github.com/Baxxu/site-donate-back/sql"
	"log"
)

const (
	cookieName   string = "RefreshToken"
	cookieMaxAge int    = 60 * 60 * 24 * 365
)

var (
	DataBase sql.DataBase

	ErrInvalid = errors.New("error refresh token invalid")
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}
