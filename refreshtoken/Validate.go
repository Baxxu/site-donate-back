package refreshtoken

import (
	"context"
	"github.com/Baxxu/site-donate-back/sql"
	"log"
	"time"
)

// Validate
//
// Чекает токен.
//
// Обновляет last_access_time в базе.
//
// Время жизни = last_access_time + год
func Validate(token []byte) (sessionId []byte, userId int, privateKey []byte, err error) {
	timeNow := time.Now()
	var sessionIdTmp *[]byte
	var userIdTmp *int
	var privateKeyTmp *[]byte
	var ok bool
	err = sql.DataBase.Pool.QueryRow(context.Background(),
		`select id, user_id, private_key, ok from refresh_token_validate($1,$2,$3) as (id bytea, user_id bigint, private_key bytea, ok boolean)`,
		token, timeNow.Unix(), timeNow.AddDate(-1, 0, 0).Unix()).
		Scan(&sessionIdTmp, &userIdTmp, &privateKeyTmp, &ok)
	if err != nil {
		log.Printf("%s", err)
		return nil, 0, nil, err
	}

	if !ok {
		return nil, 0, nil, ErrInvalid
	}

	return *sessionIdTmp, *userIdTmp, *privateKeyTmp, nil
}
