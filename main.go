package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

type URL string

// Marshal is take an Interace from the Goword turn that into JSON.
// Unmarshal is opposite, take some json to Goword

type URLDescription struct {
	URL         URL    `json:"url"` // field struct tag
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json: "-,"`
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL:         URL("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         URL("/blocks"),
			Method:      "POST",
			Description: "Add a block",
			Payload:     "data: string",
		},
	}

	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)
}

func (u URL) MarshalText() (text []byte, err error) {
	url := fmt.Sprintf("localhost%s%s", port, u)

	return []byte(url), nil
}

func main() {
	http.HandleFunc("/", documentation)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
