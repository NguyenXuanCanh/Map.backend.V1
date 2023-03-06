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
	// guidance rule:
	// 		TURN_SHARP_LEFT=-3
	// 		TURN_LEFT=-2
	// 		TURN_SLIGHT_LEFT=-1
	// 		CONTINUE_ON_STREET=0
	// 		TURN_SLIGHT_RIGHT=1
	// 		TURN_RIGHT=2
	// 		TURN_SHARP_RIGHT=3
	// 		FINISH=4
	// 		VIA_REACHED=5
	// 		USE_ROUNDABOUT=6
	// 		KEEP_RIGHT=7
	// 		UTURN=-98
}
type Location []float64
