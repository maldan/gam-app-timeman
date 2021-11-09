package core

import "time"

var DataDir = ""

type Activity struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

type Task struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Start       time.Time `json:"start"`
	Stop        time.Time `json:"stop"`
}

type Sleep struct {
	Id          string    `json:"id"`
	Description string    `json:"description"`
	Start       time.Time `json:"start"`
	Stop        time.Time `json:"stop"`
}

type Config struct {
	ActivityColor map[string]string `json:"activityColor"`
}

var AppConfig = Config{}
