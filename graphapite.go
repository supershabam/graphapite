package graphapite

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/supershabam/graphapite/structs"
)

type Graphapite struct {
	Handler http.Handler
	Store   Store
}

func NewGraphapite(store Store) *Graphapite {
	g := &Graphapite{Store: store}
	r := mux.NewRouter()
	r.HandleFunc("/metrics/find/", g.FindHandler).Methods("GET")
	r.HandleFunc("/render", g.RenderHandler).Methods("POST").Queries("format", "json")
	r.HandleFunc("/", g.NotFoundHandler)
	g.Handler = r
	return g
}

func (g Graphapite) FindHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	nodes, err := g.Store.Nodes(structs.Pattern(r.Form.Get("query")))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(nodes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func (g Graphapite) RenderHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))
}

func (g Graphapite) NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func (g Graphapite) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	g.Handler.ServeHTTP(w, r)
}
