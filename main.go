package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/DongSeonYoo/go-coin/blockchain"
)

const (
	templateDir string = "templates/"
	port        string = ":4000"
)

var templates *template.Template

// render our block
type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(w http.ResponseWriter, r *http.Request) {
	data := homeData{PageTitle: "Home", Blocks: blockchain.GetBlockChain().AllBlocks()}

	if err := templates.ExecuteTemplate(w, "home", data); err != nil {
		panic(err)
	}
}

func add(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if err := templates.ExecuteTemplate(rw, "add", nil); err != nil {
			panic(err)
		}
	case "POST":
		if err := r.ParseForm(); err != nil {
			panic(err)
		}
		data := r.Form.Get("blockData")
		blockchain.GetBlockChain().AddBlock(data)
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	}
}

func main() {
	templates = template.Must(templates.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))

	http.HandleFunc("/", home)
	http.HandleFunc("/add", add)
	fmt.Printf("Listening on http://localhost%s", port)

	log.Fatal(http.ListenAndServe(port, nil))
}
