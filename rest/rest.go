package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/DongSeonYoo/go-coin/blockchain"
	"github.com/DongSeonYoo/go-coin/utils"
)

type url string

var port string

// Marshal is take an Interace from the Goword turn that into JSON.
// Unmarshal is opposite, take some json to Goword

type urlDescription struct {
	URL         url    `json:"url"` // field struct tag
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

type addBlockBody struct {
	Message string `json:"message"`
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         url("/blocks"),
			Method:      "POST",
			Description: "Add a block",
			Payload:     "data: string",
		},
	}

	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)
}

func (u url) MarshalText() (text []byte, err error) {
	url := fmt.Sprintf("localhost%s%s", port, u)

	return []byte(url), nil
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rw.Header().Add("Content-Type", "application/json")
		utils.HandleErr(json.NewEncoder(rw).Encode(blockchain.GetBlockChain().AllBlocks()))

	case "POST":
		var addBlockBody addBlockBody
		rw.Header().Add("Content-Type", "application/json")
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
		blockchain.GetBlockChain().AddBlock(addBlockBody.Message)

		rw.WriteHeader(http.StatusCreated)
	}
}

func Start(aPort int) {
	handler := http.NewServeMux()
	port := fmt.Sprintf(":%d", aPort)
	handler.HandleFunc("/", documentation)
	handler.HandleFunc("/blocks", blocks)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, handler))
}
