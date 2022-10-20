package sessionid

import (
	"log"
	mathRand "math/rand"
	"time"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	mathRand.Seed(time.Now().Unix())
}
