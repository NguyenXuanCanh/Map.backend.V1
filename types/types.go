package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Weight   int     `json:"weight"`
	Size     int     `json:"size"`
	Price    float32 `json:"price"`
	Quantity int     `json:"quantity"`
}

type Package struct {
	Id         int                `json:"id"`
	Account_id int                `json:"account_id"`
	Total      int                `json:"total"`
	Date       primitive.DateTime `bson:"date"`
	Location   string             `json:"location"`
}

type Trip struct {
	Id         int                `json:"id"`
	Account_id int                `json:"account_id"`
	Package_id int                `json:"package_id"`
	Date       primitive.DateTime `bson:"date"`
}
type Job struct {
	Id          int      `json:"id"`
	Location    Location `json:"location"`
	Amount      []int    `json:"amount"`
	Description string   `json:"description"`
	// Service  []int
}

type Vehicle struct {
	Id       int      `json:"id"`
	Start    Location `json:"start"`
	End      Location `json:"end"`
	Capacity []int    `json:"capacity"`
}

type Location []float64
