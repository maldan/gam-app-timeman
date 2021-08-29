package api

import "time"

type ArgsEmpty struct {
}

type ArgsDate struct {
	Date time.Time `json:"date"`
}
