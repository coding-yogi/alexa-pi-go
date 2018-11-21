package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {

		//logRequest(r)
		message, err := AlexaRequestHandler(r)
		if err != nil {
			fmt.Println(err)
			message = "Sorry, I could not handle this request"
		}
		res := GenerateResponse(message)
		w.Header().Set("content-type", "application/json")
		w.Write(res)
	})

	http.ListenAndServe(":8080", nil)
}

// formatRequest generates ascii representation of a request
func logRequest(r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))
}
