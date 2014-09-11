package graphapite

import (
	"encoding/json"
	"time"
)

type Datapoint struct {
	Time  time.Time
	Value float64
}

func (d Datapoint) MarshalJSON() ([]byte, error) {
	return json.Marshal([]interface{}{
		d.Value,
		d.Time.UnixNano() / 1e9,
	})
}
