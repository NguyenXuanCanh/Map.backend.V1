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
	Id          int                `json:"id"`
	Account_id  int                `json:"account_id"`
	Total       int                `json:"total"`
	Date        primitive.DateTime `bson:"date"`
	Status      string             `json:"status"`
	Location    Location           `json:"location"`
	Description string             `json:"description"`
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

type VehicleDB struct {
	Id          string `json:"_id"`
	Brand       string `json:"brand"`
	License     string `json:"license"`
	Owner_name  string `json:"owner_name"`
	Tank_volume int    `json:"tank_volume"`
	Tank_weight int    `json:"tank_weight"`
}

type History struct {
	Id         int                `json:"_id"`
	Account_id int                `json:"account_id"`
	Package_id int                `json:"package_id"`
	Date       primitive.DateTime `json:"date"`
}

type Route struct {
	Code  string `json:"code"`
	Paths []Path `json:"path"`
}

type Path struct {
	Distance     int           `json:"distance"`
	Instructions []Instruction `json:"instruction"`
	Bbox         []float64     `json:"bbox"`
	Points       string        `json:"points"`
}

type Instruction struct {
	Text       string `json:"text"`
	StreetName string `json:"street_name"`
	Distance   int    `json:"distance"`
	Time       int    `json:"time"`
	Interval   string `json:"interval"`
	Sign       string `json:"sign"`
}
type Location []float64
