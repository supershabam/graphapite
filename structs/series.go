package structs

type Series struct {
	// Name by default is the key that generated this series, but functions may
	// change the name
	Name string `json:"target"`
	// TimesortedDatapoints are datapoints indexed from 0 -> N in increasing unix time
	// order. Order is not enforced, but functions creating Series should preserve
	// this ordering.
	TimesortedDatapoints []Datapoint `json:"datapoints"`
}
