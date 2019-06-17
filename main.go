package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gaborszakacs/yo-app/chatbot"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/responses", responsesHandler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

// NewResponseParams is the format expected when creating a new response
type NewResponseParams struct {
	Input string
}

// Conversation is responded when a new response is requsted
type Conversation struct {
	Input    string `json:"input"`
	Response string `json:"response"`
}

func responsesHandler(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var rp NewResponseParams
	err := decoder.Decode(&rp)

	if err != nil {
		panic(err)
	}

	resp := Conversation{Response: chatbot.Respond(rp.Input), Input: rp.Input}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
