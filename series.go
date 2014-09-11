package graphapite

type Series struct {
	Key        Key         `json:"-"`
	Name       string      `json:"target"`
	Datapoints []Datapoint `json:"datapoints"`
}
