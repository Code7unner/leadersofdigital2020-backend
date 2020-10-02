package nexmo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/code7unner/leadersofdigital2020-backend/utils"
	"net/http"
	"strconv"
)

type Payload struct {
	From      string `json:"from"`
	To        string `json:"to"`
	Text      string `json:"text"`
	ApiKey    string `json:"api_key"`
	ApiSecret string `json:"api_secret"`
}

func SendMessage(name, phone, key, secret string) (*http.Response, error) {
	code := strconv.Itoa(utils.GenerateRandom4DigitNumber())

	data := &Payload{
		From:      name,
		To:        phone,
		Text:      fmt.Sprintf("Dear %s, your code is %s", name, code),
		ApiKey:    key,
		ApiSecret: secret,
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://rest.nexmo.com/sms/json", body)
	if err != nil {
		return nil, err
	}

	// Ensure headers
	req.SetBasicAuth(key, secret)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return resp, nil
}
