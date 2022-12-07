package api

import (
	"log"
	"net/http"

	"git.furqansoftware.net/hjr265/xform/cfg"
	"git.furqansoftware.net/hjr265/xform/pipe"
	"github.com/gorilla/mux"
)

func handleRun(wr http.ResponseWriter, r *http.Request) {
	pid := mux.Vars(r)["pid"]
	cp, ok := cfg.Registry[pid]
	if !ok {
		http.Error(wr, "Not Found", http.StatusNotFound)
		return
	}
	p, err := pipe.NewPipeline(cp)
	if err != nil {
		log.Println("pipeline:", err)
		http.Error(wr, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	p.Values.Set("@Request", r)
	p.Values.Set("@ResponseWriter", wr)
	err = p.Run()
	if err != nil {
		log.Println("pipeline run:", err)
		http.Error(wr, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func init() {
	Router.NewRoute().Methods("POST").Path("/api/{pid}/run").HandlerFunc(handleRun)
}
