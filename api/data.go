package api

import "time"

type Days struct {
	NumberOfDays *int `json:numberofdays",omitempty"`
}

type DateTime struct {
	IncrementedDateTime time.Time `json:`
}

type Sum struct {
	Total int `json:`
}

type Numbers struct {
	First  *int `json:"first,omitempty"`
	Second *int `json:"second,omitempty"`
}
