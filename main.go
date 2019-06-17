package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gaborszakacs/yo/response"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

// NewResponseParams is the format expected when creating a new response
type NewResponseParams struct {
	To string
}

// ResponseResponse is the format expected when creating a new response
type ResponseResponse struct {
	response string
}

func responsesHandler(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var rp NewResponseParams
	err := decoder.Decode(&rp)

	if err != nil {
		panic(err)
	}

	fmt.Println(rp.To)
	resp := response.R{Response: "jani"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/responses", responsesHandler)
	http.ListenAndServe(":8080", nil)
}
