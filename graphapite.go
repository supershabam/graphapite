package graphapite

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Graphapite struct {
	Handler  http.Handler
	Store    Store
	Resolver Resolver
}

func NewGraphapite(store Store) *Graphapite {
	g := &Graphapite{Store: store}

	resolver := StoreResolver{
		Store:     store,
		Functions: map[string]SeriesFn{},
	}

	resolver.Functions["alias"] = Alias
	resolver.Functions["aliasByNode"] = AliasByNode

	g.Resolver = resolver

	r := mux.NewRouter()
	r.HandleFunc("/metrics/find/", g.FindHandler).Methods("GET")
	r.HandleFunc("/render", g.RenderHandler).Methods("POST")
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
	nodes, err := g.Store.Nodes(Pattern(r.Form.Get("query")))
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
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seri := []Series{}
	for _, rawtarget := range r.Form["target"] {
		var target Target
		err = target.Parse(rawtarget)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// TODO parse from and until variables
		series, err := g.Resolver.Resolve(target, time.Now(), time.Now())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		seri = append(seri, series...)
	}

	b, err := json.Marshal(seri)
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
