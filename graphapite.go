package graphapite

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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

func (g Graphapite) GetSeries(target string) ([]Series, error) {
	if strings.HasPrefix(target, "alias(") && strings.HasSuffix(target, ")") {
		target = strings.TrimPrefix(target, "alias")
		target = strings.TrimSuffix(target, ")")
		parts := strings.SplitN(target, ",", 2)
		series, err := g.GetSeries(parts[0])
		if err != nil {
			return series, nil
		}
		return []Series{}, fmt.Errorf("NOT IMPLEMENTED BITCH")
	}

	datapoints, err := g.Store.Get(Key(target), time.Now(), time.Now())
	if err != nil {
		return []Series{}, err
	}
	return []Series{
		Series{
			Name:       target,
			Datapoints: datapoints,
		},
	}, nil
}

func (g Graphapite) RenderHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seri := []Series{}
	for _, target := range r.Form["target"] {
		series, err := g.GetSeries(target)
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
