package graphapite

import (
	"testing"
	"time"
)

func TestAlias(t *testing.T) {
	r := MockResolver{
		ResolveFn: func(target string, from, until time.Time) ([]Series, error) {
			if target != "rawtarget" {
				t.Fatalf("expected rawtarget to be \"rawtarget\" not %s", target)
			}
			return []Series{
				Series{
					Name: "jerks",
					Datapoints: []Datapoint{
						Datapoint{
							Time:  time.Now(),
							Value: 42.0,
						},
					},
				},
			}, nil
		},
	}

	series, err := Alias(r, []string{"rawtarget", "newname"}, time.Now().Add(-time.Minute), time.Now())
	if err != nil {
		t.Fatal(err)
	}
	if len(series) != 1 {
		t.Fatalf("expected just one series, but got %d", len(series))
	}
	if series[0].Name != "newname" {
		t.Fatalf("expected name to be \"newname\" not %s", series[0].Name)
	}
}
