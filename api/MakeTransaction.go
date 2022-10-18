package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Baxxu/keyGen"
	"github.com/Baxxu/site-donate-back/keys"
	"io"
	"log"
	"net/http"
	"strconv"
)

//TODO MakeTransaction, когда подключу переводы

func MakeTransaction(writer http.ResponseWriter, request *http.Request) {
	if request.ContentLength > 2048 {
		return
	}

	dataTemp, err := io.ReadAll(http.MaxBytesReader(writer, request.Body, 2048))
	if err != nil {
		log.Printf("%s", err)
	}
	request.Body.Close()

	var makeTransactionRequest MakeTransactionRequest
	err = json.Unmarshal(dataTemp, &makeTransactionRequest)
	if err != nil {
		log.Printf("%s", err)
	}

	value, err := strconv.Atoi(string(dataTemp))
	if err != nil {
		return
	}

	if value < 1 {
		return
	}

	payment := Payment{
		Amount: Amount{
			Value:    string(dataTemp),
			Currency: "RUB",
		},
		Capture: true,
		Confirmation: Confirmation{
			Type:      "redirect",
			ReturnUrl: "https://testest.ru/return_url",
		},
	}

	paymentJson, err := json.Marshal(payment)
	if err != nil {
		log.Printf("Error doing JSON Marshal\n%s\n", err)
		return
	}

	req, err := http.NewRequest("POST", "https://api.yookassa.ru/v3/payments", bytes.NewReader(paymentJson))

	req.SetBasicAuth(keys.ShopId, keys.ShopKey)

	randKey, _ := keyGen.GetRandKey(8)

	req.Header.Set("Content-Type", `application/json`)
	req.Header.Set("Idempotence-Key", string(randKey))

	resp, err := Client.Do(req)
	if err != nil {
		log.Printf("Error doing POST request\n%s\n", err)
		return
	}
	dataTempTemp, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp.Body.Close()
	fmt.Printf("%s\n", dataTempTemp)
}

type MakeTransactionRequest struct {
	Amount Amount `json:"amount"`
	Token  string `json:"token"`
}

type MakeTransactionResponse struct {
	Ok         bool `json:"ok"`
	ErrorToken bool `json:"error_token,omitempty"`
}

type Payment struct {
	Amount       Amount       `json:"amount"`
	Capture      bool         `json:"capture"`
	Confirmation Confirmation `json:"confirmation"`
	Description  string       `json:"description,omitempty"`
}

type Amount struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

type Confirmation struct {
	Type      string `json:"type"`
	ReturnUrl string `json:"return_url"`
}
