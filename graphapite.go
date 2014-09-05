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

// FindNode provides this god awful datatype
// https://github.com/graphite-project/graphite-web/blob/master/webapp/graphite/metrics/views.py
type FindNode struct {
	AllowChildren bool   `json:"allowChildren"`
	Expandable    bool   `json:"expandable"`
	Id            string `json:"id"`
	Leaf          bool   `json:"leaf"`
	Text          string `json:"text"`
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
	fr := FindResponse{
		FindNode{
			Id:            "stats.gauges.echo_server",
			Expandable:    true,
			Text:          "echo_server",
			Leaf:          true,
			AllowChildren: true,
		},
	}
	b, err := json.Marshal(fr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func (g Graphapite) NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func (g Graphapite) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	g.Handler.ServeHTTP(w, r)
}
