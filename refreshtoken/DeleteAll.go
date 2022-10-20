package refreshtoken

import (
	"context"
	"github.com/Baxxu/site-donate-back/sql"
	"log"
)

func DeleteAll(token []byte) (err error) {
	_, err = sql.DataBase.Pool.Exec(context.Background(),
		`delete
from sessions
where user_id in (select user_id
                  from sessions
                  where refresh_token = $1);`,
		token)
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	return nil
}
