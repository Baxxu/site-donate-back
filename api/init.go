package api

import (
	"log"
	"net/http"
)

var (
	Client *http.Client
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}
