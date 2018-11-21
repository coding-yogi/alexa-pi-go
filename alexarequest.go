package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AlexaRequest struct {
	Version string  `json:"version"`
	Session Session `json:"session"`
	Request Request `json:"request"`
}

type Session struct {
	New         bool        `json:"new"`
	SessionID   string      `json:"sessionId"`
	Application Application `json:"application"`
	User        User        `json:"user"`
}

type Application struct {
	ApplicationId string `json:"applicationId"`
}

type User struct {
	UserId string `json:"userId"`
}

type Request struct {
	Type   string `json:"type"`
	Intent Intent `json:"intent"`
}

type Intent struct {
	Name  string `json:"name"`
	Slots Slots  `json:"slots"`
}

type Slots struct {
	Device Slot `json:"device"`
	Status Slot `json:"status"`
}

type Slot struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

func OnOffIntentHandler(intent Intent) (string, error) {
	//Get device and status
	device := intent.Slots.Device.Value
	status := intent.Slots.Status.Value
	successMessage := "turned " + status + " " + device

	if device == "" || status == "" {
		fmt.Println("Device or status is not defined")
		return "", errors.New("Device or status is not defined")
	}

	return successMessage, nil
}

func AlexaRequestHandler(r *http.Request) (string, error) {
	//Unmarshal
	body, _ := ioutil.ReadAll(r.Body)
	alexaRequest := AlexaRequest{}
	unhandledIntent := "Sorry, I don't know how to handle this request"

	if err := json.Unmarshal(body, &alexaRequest); err != nil {
		return "", err
	}

	//Check if request is valid - Tokens , Users and all

	//Check Intent and call the required Handler
	intent := alexaRequest.Request.Intent
	if intent.Name == "GPIOControlIntent" {
		fmt.Println("Calling " + intent.Name)
		return OnOffIntentHandler(intent)
	}

	return unhandledIntent, nil
}
