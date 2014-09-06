package structs

import "time"

type Datapoint struct {
	Time  time.Time
	Value float64
}
