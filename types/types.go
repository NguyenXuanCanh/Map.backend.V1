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
