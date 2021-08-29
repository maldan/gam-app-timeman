package core

import "time"

var DataDir = ""

type Task struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Start       time.Time `json:"start"`
	Stop        time.Time `json:"stop"`
}
