package sessionid

import (
	"github.com/Baxxu/site-donate-back/sql"
	"log"
	mathRand "math/rand"
	"time"
)

var (
	DataBase sql.DataBase
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	mathRand.Seed(time.Now().Unix())
}
