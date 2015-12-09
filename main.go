package main

//go:generate go-bindata -pkg presentation -o presentation/bindata.go templates

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/agonzalezro/md2slides/presentation"
	"github.com/gorilla/mux"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	availableThemes = []string{"reveal"}
	theme           *string

	outputFile  = kingpin.Flag("output", "output file where to write the HTML.").Default("/dev/stdout").Short('o').OpenFile(os.O_CREATE|os.O_WRONLY, 0644)
	startDaemon = kingpin.Flag("daemon", "start a simple HTTP serving your slides.").Short('d').Bool()
	port        = kingpin.Flag("port", "port where to run the server.").Default("8080").Int()
	config      = kingpin.Flag("theme-config", "configuration for the theme (JS file)").Short('c').File()

	source = kingpin.Arg("source", "Markdown source file.").Required().File()
)

func init() {
	themesHelp := fmt.Sprintf("Choose one: [%s].", strings.Join(availableThemes, ", "))
	theme = kingpin.Flag("theme", themesHelp).Default("reveal").String()

	kingpin.CommandLine.HelpFlag.Short('h')
}

func main() {
	kingpin.Parse()
	if !contains(availableThemes, *theme) {
		kingpin.Fatalf("theme: %s not found", *theme)
	}

	p, err := presentation.NewFromFile(*source)
	ifErrFatal(err)

	p.Theme = *theme

	if *startDaemon {
		r := mux.NewRouter()

		r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			p.Reload()
			ifErrFatal(p.WriteWithConfig(w, *config))
		})

		port := ":" + strconv.Itoa(*port)
		log.Println("Serving slides at", port)
		log.Fatal(http.ListenAndServe(port, r))

		return
	}

	// Write it just if we don't serve it
	ifErrFatal(p.WriteWithConfig(*outputFile, *config))
}
