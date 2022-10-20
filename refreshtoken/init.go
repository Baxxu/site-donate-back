package refreshtoken

import (
	"errors"
	"log"
)

const (
	cookieName   string = "RefreshToken"
	cookieMaxAge int    = 60 * 60 * 24 * 365
)

var (
	ErrInvalid = errors.New("error refresh token invalid")
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}
