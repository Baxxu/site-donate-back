package api

import (
	"log"
	"net/http"
	"time"
)

var (
	Client = http.Client{
		Timeout: time.Second * 60,
	}
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}
