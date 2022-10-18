package accesstoken

import (
	"encoding/hex"
	"fmt"
	"time"
)

func New(userId int, sessionId, privateKey []byte) (token string) {
	data, hash := hash_(
		token_{
			creationTime: int(time.Now().Unix()),
			sessionId:    sessionId,
			userId:       userId,
		},
		privateKey,
	)

	hashStr := hex.EncodeToString(hash)

	return fmt.Sprintf("%s&hash=%s", data, hashStr)
}
