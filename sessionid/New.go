package sessionid

import (
	"context"
	"crypto/rand"
	"log"
	mathRand "math/rand"
	"time"
)

func New(userId int) (sessionId, privateKey, refreshToken []byte, err error) {
	sessionId = make([]byte, 32)
	_, err = mathRand.Read(sessionId)
	if err != nil {
		log.Printf("%s", err)
		return
	}

	privateKey = make([]byte, 32)
	_, err = rand.Read(privateKey)
	if err != nil {
		log.Printf("%s", err)
		return
	}

	refreshToken = make([]byte, 32)
	_, err = rand.Read(refreshToken)
	if err != nil {
		log.Printf("%s", err)
		return
	}

	timeNow := time.Now().Unix()

	_, err = DataBase.Pool.Exec(context.Background(),
		`insert into sessions (id,
                      user_id,
                      private_key,
                      refresh_token,
                      creation_time,
                      last_access_time)
values ($1, $2, $3, $4, $5, $6);`,
		sessionId,
		userId,
		privateKey,
		refreshToken,
		timeNow,
		timeNow,
	)
	if err != nil {
		log.Printf("%s", err)
		return
	}

	return sessionId, privateKey, refreshToken, nil
}
