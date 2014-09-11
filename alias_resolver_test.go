package graphapite

import (
	"testing"
	"time"
)

func TestAlias(t *testing.T) {
	a := Alias{
		Resolver: MockResolver{
			ResolveFn: func(rawtarget string) ([]Series, error) {
				if rawtarget != "rawtarget" {
					t.Fatalf("expected rawtarget to be \"rawtarget\" not %s", rawtarget)
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
		},
	}

	series, err := a.Execute([]string{"rawtarget", "newname"})
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
