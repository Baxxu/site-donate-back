package accesstoken

import (
	"errors"
	"log"
)

const (
	cookieName   string = "AccessToken"
	cookieMaxAge int    = 60 * 60 * 24
)

var (
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
