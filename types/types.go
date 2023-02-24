package types

type Product struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Weight   int     `json:"weight"`
	Size     int     `json:"size"`
	Price    float32 `json:"price"`
	Quantity int     `json:"quantity"`
}

type Package struct {
	Id       string    `json:"id"`
	Products []Product `json:"products"`
	Location string    `json:"location"`
}

type Job struct {
	Id       string
	Location []float64
	// Service  []int
}

type Vehicle struct {
	Id    string
	Start []float64
	End   []float64
}
