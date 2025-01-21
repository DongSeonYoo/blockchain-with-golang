package explorer

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/DongSeonYoo/go-coin/blockchain"
)

const (
	templateDir string = "templates/"
)

var templates *template.Template

// render our block
type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(w http.ResponseWriter, r *http.Request) {
	data := homeData{PageTitle: "Home", Blocks: blockchain.BlockChain().AllBlocks()}

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
		blockchain.BlockChain().AddBlock(data)
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	}
}

func Start(port int) {
	handler := http.NewServeMux()

	templates = template.Must(templates.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))

	handler.HandleFunc("/", home)
	handler.HandleFunc("/add", add)
	fmt.Printf("Listening on http://localhost:%d\n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}
