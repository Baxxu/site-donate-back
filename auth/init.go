package auth

import (
	"github.com/Baxxu/site-donate-back/sql"
	"log"
)

var (
	DataBase sql.DataBase
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}
