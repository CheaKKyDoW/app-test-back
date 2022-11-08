package models

import "time"

type JsonReturn struct {
	Status int           `json:"status"`
	Time   time.Time     `json:"time"`
	Data   []interface{} `json:"data"`
}
