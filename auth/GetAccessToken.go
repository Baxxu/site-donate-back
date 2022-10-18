package auth

import (
	"encoding/json"
	"github.com/Baxxu/site-donate-back/accesstoken"
	"github.com/Baxxu/site-donate-back/refreshtoken"
	"log"
	"net/http"
)

func GetAccessToken(writer http.ResponseWriter, request *http.Request) {
	refreshToken, err := refreshtoken.ReadFromCookie(request)
	if err != nil {
		resp, err := json.Marshal(GetAccessTokenResponse{
			Ok:      false,
			Message: "error: no RefreshToken cookie",
		})
		if err != nil {
			log.Printf("%s", err)
			return
		}

		writer.Header().Set(`Content-Type`, `application/json`)
		writer.Write(resp)

		return
	}

	sessionId, userId, privateKey, err := refreshtoken.Validate(refreshToken)
	if err != nil {
		log.Printf("%s", err)

		resp, err := json.Marshal(GetAccessTokenResponse{
			Ok:      false,
			Message: "error: RefreshToken invalid",
		})
		if err != nil {
			log.Printf("%s", err)
			return
		}

		writer.Header().Set(`Content-Type`, `application/json`)
		writer.Write(resp)

		return
	}

	accessToken := accesstoken.New(userId, sessionId, privateKey)

	accesstoken.WriteToCookie(accessToken, writer)

	resp, err := json.Marshal(GetAccessTokenResponse{
		Ok:          true,
		AccessToken: accessToken,
	})
	if err != nil {
		log.Printf("%s", err)
		return
	}

	writer.Header().Set(`Content-Type`, `application/json`)
	writer.Write(resp)
}

type GetAccessTokenResponse struct {
	Ok          bool   `json:"ok"`
	Message     string `json:"message,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
}
