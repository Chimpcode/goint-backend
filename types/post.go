package types

import "time"


type Location struct {
	Longitude float64 `json:"longitude" form:"longitude"`
	Latitude float64 `json:"latitude" form:"latitude"`
}

type Post struct {
	Id string `json:"id" storm:"id"`

	By string `json:"by" form:"by"`

	CreatedAt time.Time `json:"created_at" storm:"index"`
	Type string `json:"type" form:"type"`
	Title string `json:"title" fako:"title" form:"title"`
	Image string `json:"image" form:"image"`
	Description string `json:"description" form:"description"`
	Address string `json:"address" fako:"street_address" form:"address"`

	Location Location `json:"location" form:"location"`
	//Promotion string `json:"promotion"`

	Stock int `json:"stock" form:"stock"`

}

type PostMinified struct {
	Id string `json:"id" storm:"id"`

	By string `json:"by" form:"by"`

	Title string `json:"title" fako:"title" form:"title"`
	Location Location `json:"location" form:"location"`

	Stock int `json:"stock" form:"stock"`
}
