package refreshtoken

import (
	"context"
	"log"
)

func Delete(token []byte) (err error) {
	_, err = DataBase.Pool.Exec(context.Background(),
		`delete
from sessions
where refresh_token = $1;`,
		token)
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	return nil
}
