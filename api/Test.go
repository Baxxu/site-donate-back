package api

import (
	"encoding/json"
	"github.com/Baxxu/site-donate-back/accesstoken"
	"log"
	"net/http"
)

func Test(writer http.ResponseWriter, request *http.Request) {
	accessToken, err := accesstoken.ReadFromCookie(request)
	if err != nil {
		log.Printf("%s", err)

		data, _ := json.Marshal(TestApiResponse{
			Ok:      false,
			Message: "error: cookie AccessToken not found",
		})

		writer.Header().Set(`Content-Type`, `application/json`)
		writer.Write(data)

		return
	}

	err = accesstoken.Validate(accessToken)
	if err != nil {
		data, _ := json.Marshal(TestApiResponse{
			Ok:      false,
			Message: "error: AccessToken invalid",
		})

		writer.Header().Set(`Content-Type`, `application/json`)
		writer.Write(data)

		return
	}

	data, _ := json.Marshal(TestApiResponse{
		Ok:      true,
		Message: "AccessToken valid",
	})

	writer.Header().Set(`Content-Type`, `application/json`)
	writer.Write(data)
}

type TestApiResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message,omitempty"`
}
