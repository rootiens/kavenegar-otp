package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func hasValidToken(tokens []Token) bool {
	for _, token := range tokens {
		if token.Key == "token" && token.Value != "" {
			return true
		}
	}
	return false
}

func createUrl(payload Otp) (string, error) {
	if payload.ApiKey == "" {
		return "", errors.New("Empty API-KEY")
	}

	if payload.Receptor == "" {
		return "", errors.New("Empty Receptor")
	}

	if payload.Template == "" {
		return "", errors.New("Empty Template")
	}

	if payload.Type == "" {
		return "", errors.New("Empty Type")
	}

	if !hasValidToken(payload.Tokens) {
		return "", errors.New("Token undefined or has empty value")
	}

	url := "https://api.kavenegar.com/v1/" + payload.ApiKey + "/verify/lookup.json?receptor=" + payload.Receptor

	for _, token := range payload.Tokens {
		url += "&" + token.Key + "=" + token.Value
	}

	url += "&template=" + payload.Template

	return url, nil
}

func Send(payload Otp) (response, error) {
	url, err := createUrl(payload)
	if err != nil {
		return response{}, err
	}

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return response{}, err
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return response{}, err
	}

	var resp response
	if err := json.Unmarshal(body, &resp); err != nil {
		return response{}, err
	}

	if res.StatusCode != 200 {
		return response{}, fmt.Errorf("StatusCode: %d. Message: %s", resp.Return.Status, resp.Return.Message)
	}

	return resp, nil
}
