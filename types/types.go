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
	Id int `json:"id"`
	// Account_id  int                `json:"account_id"`
	Total       int                `json:"total"`
	Date        primitive.DateTime `bson:"date"`
	Status      string             `json:"status"`
	Location    Location           `json:"location"`
	Description string             `json:"description"`
	Volume      int                `json:"volume"`
	Weight      int                `json:"weight"`
	Type        int                `json:"type"`
}

type Trip struct {
	Id         int                `json:"id"`
	Account_id int                `json:"account_id"`
	Package_id int                `json:"package_id"`
	Date       primitive.DateTime `bson:"date"`
}

type TripAdd struct {
	Id   int    `json:"id"`
	Data string `json:"data"`
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
	Id primitive.ObjectID `bson:"_id,omitempty"`

	Brand       string `json:"brand"`
	License     string `json:"license"`
	Owner_name  string `json:"owner_name"`
	Tank_volume int    `json:"tank_volume"`
	Tank_weight string `json:"tank_weight"`
}

type History struct {
	Id         int                `json:"_id"`
	Account_id string             `json:"account_id"`
	Package_id int                `json:"package_id"`
	Date       primitive.DateTime `json:"date"`
}

type HistoryRes struct {
	Id          int                `json:"_id"`
	Account_id  string             `json:"account_id"`
	Package_id  int                `json:"package_id"`
	Date        primitive.DateTime `json:"date"`
	Status      string             `json:"status"`
	Total       int                `json:"total"`
	Description string             `json:"description"`
	Volume      int                `json:"volume"`
	Weight      int                `json:"weight"`
}

// type Route struct {
// 	Code  string `json:"code"`
// 	Paths []Path `json:"path"`
// }

// type Path struct {
// 	Distance     int           `json:"distance"`
// 	Instructions []Instruction `json:"instruction"`
// 	Bbox         []float64     `json:"bbox"`
// 	Points       string        `json:"points"`
// }

// type Instruction struct {
// 	Text       string `json:"text"`
// 	StreetName string `json:"street_name"`
// 	Distance   int    `json:"distance"`
// 	Time       int    `json:"time"`
// 	Interval   string `json:"interval"`
// 	Sign       string `json:"sign"`
// }

type ProfileImage struct {
	Account_id string           `json:"account_id"`
	Image      primitive.Binary `json:"image"`
}
type TripDB struct {
	Code   int `json:"code"`
	Sumary struct {
		Cost            int   `json:"cost"`
		Unassigned      int   `json:"unassigned"`
		Delivery        []int `json:"delivery"`
		Amount          []int `json:"amount"`
		Pickup          []int `json:"pickup"`
		Service         int   `json:"service"`
		Duration        int   `json:"duration"`
		Waiting_time    int   `json:"waiting_time"`
		Priority        int   `json:"priority"`
		Distance        int   `json:"distance"`
		Computing_times struct {
			Loading int `json:"loading"`
			Solving int `json:"solving"`
			Routing int `json:"routing"`
		}
	} `json:"computing_times"`
	Unassigned []any `json:"unassigned"`
	Routes     []struct {
		Vehicle      int   `json:"vehicle"`
		Cost         int   `json:"cost"`
		Delivery     []int `json:"delivery"`
		Amount       []int `json:"amount"`
		Pickup       []int `json:"pickup"`
		Service      int   `json:"service"`
		Duration     int   `json:"duration"`
		Waiting_time int   `json:"waiting_time"`
		Priority     int   `json:"priority"`
		Distance     int   `json:"distance"`
		Steps        []struct {
			Type         string   `json:"type"`
			Description  string   `json:"description"`
			Location     Location `json:"location"`
			Id           int      `json:"id"`
			Service      int      `json:"service"`
			Waiting_time int      `json:"waiting_time"`
			Job          int      `json:"job"`
			Load         []int    `json:"load"`
			Arrival      int      `json:"arrival"`
			Duration     int      `json:"duration"`
			Distance     int      `json:"distance"`
		} `json:"steps"`
		Geometry string `json:"geometry"`
	} `json:"routes"`
}

type Notification struct {
	Account_id string             `json:"account_id"`
	Package_id string             `json:"package_id"`
	Type       string             `json:"type"`
	Time       primitive.DateTime `json:"time"`
}

type ClockinStatus struct {
	Status     string `json:"status"`
	Account_id string `json:"account_id"`
}

type Location []float64
