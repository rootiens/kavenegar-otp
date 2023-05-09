package main

type Token struct {
	Key   string
	Value string
}

type Otp struct {
	ApiKey   string
	Receptor string
	Template string
	Type     string
	Tokens   []Token
}

type response struct {
	Return struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	} `json:"return"`
	Entries []struct {
		MessageID  int    `json:"messageid"`
		Message    string `json:"message"`
		Status     int    `json:"status"`
		StatusText string `json:"statustext"`
		Sender     string `json:"sender"`
		Receptor   string `json:"receptor"`
		Date       int    `json:"date"`
		Cost       int    `json:"cost"`
	} `json:"entries"`
}
