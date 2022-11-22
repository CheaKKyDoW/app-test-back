package models

// type JsonReturn struct {
// 	Status int         `json:"status"`
// 	Time   time.Time   `json:"time"`
// 	Data   interface{} `json:"data"`
// }

type JsonReturn struct {
	Title       string      `json:"title"`
	Status      int         `json:"status"`
	Description string      `json:"description"`
	Result      interface{} `json:"result"`
}
