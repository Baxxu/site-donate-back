package auth

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/Baxxu/site-donate-back/accesstoken"
	"github.com/Baxxu/site-donate-back/keys"
	"github.com/Baxxu/site-donate-back/refreshtoken"
	"github.com/Baxxu/site-donate-back/sessionid"
	"github.com/Baxxu/site-donate-back/sql"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

//id, first_name, last_name, username, photo_url, auth_date and hash fields

/*
auth_date
first_name
hash
id
last_name
photo_url
username
*/

func Telegram(writer http.ResponseWriter, request *http.Request) {
	//параметры запроса из URL
	params := request.URL.Query()

	//сортирую параметры по алфавиту
	keyValues := make([]string, 0, len(params))
	for key, value := range params {
		if key == "hash" {
			continue
		}
		if len(value) != 1 {
			return
		}
		keyValue := fmt.Sprintf("%s=%s", key, value[0])
		keyValues = append(keyValues, keyValue)
	}

	sort.Strings(keyValues)

	dataCheckString := strings.Join(keyValues, "\n")

	//чекаю хеши https://core.telegram.org/widgets/login
	sha256Hasher := sha256.New()
	sha256Hasher.Write(keys.TelegramBotKey)
	botKeyHash := sha256Hasher.Sum(nil)

	hmacHasher := hmac.New(sha256.New, botKeyHash)
	hmacHasher.Write([]byte(dataCheckString))
	dataCheckStringHmac := hmacHasher.Sum(nil)

	//если хеши не совпали то это скам
	hashHex, err := hex.DecodeString(params.Get("hash"))
	if err != nil {
		return
	}
	if !bytes.Equal(dataCheckStringHmac, hashHex) {
		return
	}

	//нет даты авторизации (по сути даты создания токена) - тоже скам
	authDate, err := strconv.Atoi(params.Get("auth_date"))
	if err != nil {
		log.Printf("%s", err)
		return
	}

	//чекаю свежесть токена (день)
	if authDate < (int(time.Now().AddDate(0, 0, -1).Unix())) {
		return
	}

	//!!!телеграм токен норм, значит нужно делать свой!!!

	//telegram_id
	telegramId, err := strconv.Atoi(params.Get("id"))
	if err != nil {
		log.Printf("%s", err)
		return
	}

	//user_id
	var userId int
	err = sql.DataBase.Pool.QueryRow(context.Background(),
		`select get_user_id_with_telegram_id($1);`,
		telegramId).
		Scan(&userId)
	if err != nil {
		log.Printf("%s", err)
		return
	}

	//!!!новая сессия!!!

	sessionId, privateKey, refreshToken, err := sessionid.New(userId)
	if err != nil {
		return
	}

	accessToken := accesstoken.New(userId, sessionId, privateKey)

	//кладу токены в куки

	refreshtoken.WriteToCookie(refreshToken, writer)

	accesstoken.WriteToCookie(accessToken, writer)

	writer.Header().Set(`Content-Type`, `application/json`)
	writer.Write([]byte(`{"ok":true}`))
}
