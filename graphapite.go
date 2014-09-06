package graphapite

import (
	"encoding/json"
	"net/http"
	"time"

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

func (g Graphapite) GetSeries(target string) ([]structs.Series, error) {
	transformer := AliasTransformer{
		NewName: "lulz",
	}
	return transformer.Transform([]structs.Series{
		structs.Series{
			Name: "omg.test",
			TimesortedDatapoints: []structs.Datapoint{
				structs.Datapoint{
					Time:  time.Now().Add(-time.Minute),
					Value: 14.4,
				},
				structs.Datapoint{
					Time:  time.Now(),
					Value: 24.4,
				},
			},
		},
	})
}

func (g Graphapite) RenderHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seri := []structs.Series{}
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
