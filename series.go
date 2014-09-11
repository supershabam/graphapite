package graphapite

type Series struct {
	Name       string      `json:"target"`
	Datapoints []Datapoint `json:"datapoints"`
}
