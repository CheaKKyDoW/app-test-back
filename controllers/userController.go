package controllers

import (
	"api_test/middlewares"
	"api_test/models"
	"api_test/webserver"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	// User ID from path `users/:id`
	// id := c.Param("id")
	res := &models.JsonReturn{
		Data:   nil,
		Status: http.StatusOK,
	}

	username := c.FormValue("username")
	password := c.FormValue("password")
	if username != "" && password != "" {
		// if { check id/pass from DB

		jwt, err := middlewares.JwtTest(username)
		if err != nil {
			return err
		}

		data := []byte(`{"token":"` + jwt + `" }`)
		var jsonMap map[string]string
		json.Unmarshal([]byte(data), &jsonMap)
		if err != nil {
			panic(err)
		}

		res.Data = jsonMap

		return c.JSON(http.StatusOK, res)

		// }
	}

	return c.JSON(http.StatusOK, res)

}

func GetUser(c echo.Context) error {
	lim := c.FormValue("limit")
	user, err := webserver.DBCon.Query("SELECT  first_name, last_name, email, gender, ip_address FROM users limit " + lim)
	if err != nil {
		fmt.Println("GetUser err DB", err)
		return nil
	}
	defer user.Close()
	cols, _ := user.Columns()

	// data := make(map[string]string)

	for user.Next() {

		columns := make([]string, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		if err := user.Scan(columnPointers...); err != nil {
			log.Fatal(err)
		}

		// for i, colName := range cols {
		// 	data[colName] = columns[i]
		// }
		//

		m := make(map[string]interface{})

		for i := range columns {

			if columns[i] == "nil" {
				m[columns[i]] = nil
			} else {
				m[columns[i]] = columns[i]
			}
		}
		v.Data = append(v.Data, m)
		//
	}

	u := models.JsonReturn{
		Data:   v.Data,
		Status: http.StatusOK,
	}

	return c.JSON(http.StatusOK, u)

}

var v struct {
	Data []interface{} // `json:"data"`
}
