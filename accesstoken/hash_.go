package accesstoken

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func hash_(token token_, privateKey []byte) (data string, hash []byte) {
	data = fmt.Sprintf("creation_time=%d&session_id=%s&user_id=%d", token.creationTime, hex.EncodeToString(token.sessionId), token.userId)

	mac := hmac.New(sha256.New, privateKey)

	mac.Write([]byte(data))

	hash = mac.Sum(nil)

	return
}
