package accesstoken

import (
	"bytes"
	"context"
	"encoding/hex"
	"github.com/jackc/pgx/v5"
	"log"
	"strconv"
	"strings"
	"time"
)

func Validate(tokenRaw string) (err error) {
	//парсим
	paramsRaw := strings.Split(tokenRaw, "&")

	if len(paramsRaw) != 4 {
		return ErrParsing
	}

	var (
		tokenFromClient token_
		hashFromClient  []byte
	)
	for _, s := range paramsRaw {
		param := strings.Split(s, "=")
		if len(param) != 2 {
			return ErrParsing
		}
		switch param[0] {
		case "creation_time":
			time_, err := strconv.Atoi(param[1])
			if err != nil {
				log.Printf("%s", err)
				return err
			}
			tokenFromClient.creationTime = time_
		case "session_id":
			id, err := hex.DecodeString(param[1])
			if err != nil {
				log.Printf("%s", err)
				return err
			}
			tokenFromClient.sessionId = id
		case "user_id":
			id, err := strconv.Atoi(param[1])
			if err != nil {
				log.Printf("%s", err)
				return err
			}
			tokenFromClient.userId = id
		case "hash":
			hash_, err := hex.DecodeString(param[1])
			if err != nil {
				log.Printf("%s", err)
				return err
			}
			hashFromClient = hash_
		}
	}
	// todo можно менять last_access_time и здесь тоже
	//проверяем
	var privateKey []byte
	err = DataBase.Pool.QueryRow(context.Background(),
		`select private_key from sessions where id = $1;`,
		tokenFromClient.sessionId).
		Scan(&privateKey)
	if err != nil {
		if err != pgx.ErrNoRows {
			log.Printf("%s", err)
			return err
		}
		return err
	}
	_, hash := hash_(tokenFromClient, privateKey)

	if !bytes.Equal(hash, hashFromClient) {
		return ErrInvalid
	}

	//старше суток - тухлый
	if tokenFromClient.creationTime < int(time.Now().AddDate(0, 0, -1).Unix()) {
		return ErrInvalid
	}

	return nil
}
