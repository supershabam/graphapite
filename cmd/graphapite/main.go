package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/supershabam/graphapite"
)

type MockStore struct{}

func (s MockStore) Get(key string, start, end time.Time) ([]graphapite.Datapoint, error) {
	return []graphapite.Datapoint{
		graphapite.Datapoint{
			Time:  time.Now().Add(-time.Minute * 2),
			Value: 1,
		},
		graphapite.Datapoint{
			Time:  time.Now().Add(-time.Minute * 1),
			Value: 11,
		},
		graphapite.Datapoint{
			Time:  time.Now(),
			Value: 5,
		},
	}, nil
}

func (s MockStore) Nodes(pattern string) ([]graphapite.Node, error) {
	fmt.Printf("matching: %s\n", pattern)
	return []graphapite.Node{
		graphapite.Node{
			Path: []string{"some", "node", "path"},
			Name: "nodename",
			Leaf: true,
		},
	}, nil
}

func (s MockStore) Write(key string, datapoint graphapite.Datapoint) error {
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
