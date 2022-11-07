package main

import (
	"api_test/router"
	"api_test/webserver"
)

func main() {

	// go func() {
	// 	err := exec.Command("explorer", "http://127.0.0.1:3000").Run()
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// }()
	e := router.New()
	webserver.PostgresConn()
	e.Start("127.0.0.1:3000")
}
