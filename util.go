package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var providers = map[string]object{
	"openrouter": {
		"url":    "https://openrouter.ai/api/v1",
		"apiKey": "",
	},
	"aihubmix": {
		"url":    "https://aihubmix.com/v1/",
		"apiKey": "",
	},
	"gptsapi": {
		"url":    "https://api.gptsapi.net/v1",
		"apiKey": "",
	},
}

var provider = ""

func Completion() {
	url := providers[provider]["url"].(string)
	apiKey := providers[provider]["apiKey"].(string)

	body := make(object)
	body["model"] = "gpt-5"
	body["messages"] = []object{
		{
			"role":    "user",
			"content": "What is the meaning of life?",
		},
	}

	data, _ := json.Marshal(body)
	payload := strings.NewReader(string(data))

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	res, _ := http.DefaultClient.Do(req)

	if res.StatusCode == 200 {
		fmt.Println("")
	}
}
