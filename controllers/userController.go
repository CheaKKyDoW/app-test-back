package controllers

import (
	"api_test/models"
	"api_test/webserver"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	// User ID from path `users/:id`
	// id := c.Param("id")
	id := c.FormValue("id")
	// return c.HTML(http.StatusOK, "<b>Thank you! "+id+"</b>")
	return c.String(http.StatusOK, id)

}

func GetUser(c echo.Context) error {
	lim := c.FormValue("limit")
	user, err := webserver.DBCon.Query("SELECT * FROM mock_data limit " + lim)
	if err != nil {
		fmt.Println("err", err)
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
			m[columns[i]] = columns[i]
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
