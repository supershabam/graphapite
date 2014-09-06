package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/supershabam/graphapite"
	"github.com/supershabam/graphapite/structs"
)

type MockStore struct{}

func (s MockStore) Get(key structs.Key, start, end time.Time) ([]structs.Datapoint, error) {
	return []structs.Datapoint{
		structs.Datapoint{
			Time:  time.Now().Add(-time.Minute * 2),
			Value: 1,
		},
		structs.Datapoint{
			Time:  time.Now().Add(-time.Minute * 1),
			Value: 11,
		},
		structs.Datapoint{
			Time:  time.Now(),
			Value: 5,
		},
	}, nil
}

func (s MockStore) Nodes(pattern structs.Pattern) ([]structs.Node, error) {
	fmt.Printf("matching: %s\n", pattern)
	return []structs.Node{
		structs.Node{
			Path: []string{"some", "node", "path"},
			Name: "nodename",
			Leaf: true,
		},
	}, nil
}

func (s MockStore) Write(key structs.Key, datapoint structs.Datapoint) error {
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
