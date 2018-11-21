package main

import (
	"encoding/json"
)

type AlexaResponse struct {
	Version  string   `json:"version"`
	Response Response `json:"response"`
}

type Response struct {
	ShouldEndSession bool         `json:"shouldEndSession"`
	OutputSpeech     OutputSpeech `json:"outputSpeech"`
}

type OutputSpeech struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func GenerateResponse(text string) []byte {
	newResponse := AlexaResponse{
		Version: "1.0",
		Response: Response{
			ShouldEndSession: true,
			OutputSpeech: OutputSpeech{
				Type: "PlainText",
				Text: text,
			},
		},
	}

	res, _ := json.Marshal(newResponse)
	return res
}
