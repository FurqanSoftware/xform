package main

import (
	"flag"
	"net/http"
	"os"

	"git.furqansoftware.net/hjr265/xform/api"
	"git.furqansoftware.net/hjr265/xform/cfg"
	"github.com/gorilla/mux"
)

var pdir = flag.String("pdir", "./pipelines", "Directory where pipeline specifications are stored.")

func main() {
	flag.Parse()

	cfg.LoadPipelines(*pdir)

	r := mux.NewRouter()
	r.NewRoute().
		PathPrefix("/api").
		Handler(api.Router)
	http.Handle("/", r)

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
