package main

import (
	"api_test/router"
)

func main() {

	// go func() {
	// 	err := exec.Command("explorer", "http://127.0.0.1:3000").Run()
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// }()
	e := router.New()

	e.Start("127.0.0.1:3000")
}
