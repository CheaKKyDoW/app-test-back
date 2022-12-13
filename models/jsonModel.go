package models

import "time"

type JsonReturn struct {
	Time        time.Time   `json:"time"`
	Title       string      `json:"title"`
	Status      int         `json:"status"`
	Description string      `json:"description"`
	Result      interface{} `json:"result"`
}
