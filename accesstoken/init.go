package accesstoken

import (
	"errors"
	"github.com/Baxxu/site-donate-back/sql"
	"log"
)

const (
	cookieName   string = "AccessToken"
	cookieMaxAge int    = 60 * 60 * 24
)

var (
	DataBase sql.DataBase

	ErrParsing = errors.New("error parsing assess token")

	ErrInvalid = errors.New("error access token invalid")
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}

type token_ struct {
	creationTime int
	sessionId    []byte
	userId       int
}
