package main

import (
	"flag"
	"net/http"
	"time"

	"github.com/supershabam/graphapite"
)

type MockStore struct{}

func (s MockStore) Get(key graphapite.Key, start, end time.Time) ([]graphapite.Datapoint, error) {
	return []graphapite.Datapoint{}, nil
}

func (s MockStore) Nodes(pattern graphapite.Pattern) ([]graphapite.Node, error) {
	return []graphapite.Node{
		graphapite.Node{
			Path: []string{"some", "node", "path"},
			Name: "nodename",
			Leaf: true,
		},
	}, nil
}

func (s MockStore) Write(key graphapite.Key, datapoint graphapite.Datapoint) error {
	return nil
}

func main() {
	addr := flag.String("addr", ":8080", "http listen addr")
	flag.Parse()

	store := MockStore{}

	g := graphapite.NewGraphapite(store)
	server := http.Server{
		Addr:    *addr,
		Handler: g,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
