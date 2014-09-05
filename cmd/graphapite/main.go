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

func (s MockStore) Write(key graphapite.Key, datapoint graphapite.Datapoint) error {
	return nil
}

// Children("") -> []string{"actual_metric", "a_prefix."}
// Find([]string{}, "") -> []string{"actual_metric", "directory."}
// Find([]string{}, "act") -> []string{"actual_metric"}
// Children("a_prefix.") -> []string{"a_prefix.actual_metric", "a_prefix.next_folder."}
func (s MockStore) Find(path []string, prefix string) ([]string, error) {
}

func (s MockStore) Keys(pattern Pattern) []Key {
}

func (s MockStore) Find(pattern graphapite.Pattern) ([]graphapite.FindNode, error) {
	return []graphapite.Key{
		"some.path.*",
		"some.path.that.is.deeper",
	}
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
