package graphapite

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Graphapite struct {
	Handler http.Handler
	Store   Store
}

type FindResponse []FindNode

type FindNode struct {
}

func NewGraphapite(store Store) *Graphapite {
	g := &Graphapite{Store: store}
	r := mux.NewRouter()
	r.HandleFunc("/metrics/find/", g.FindHandler).Methods("GET")
	r.HandleFunc("/", g.NotFoundHandler)
	g.Handler = r
	return g
}

func (g Graphapite) FindHandler(w http.ResponseWriter, r *http.Request) {
	fr := FindResponse{}
	b, err := json.Marshal(fr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(b)
}

func (g Graphapite) NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func (g Graphapite) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.Handler.ServeHTTP(w, r)
}
