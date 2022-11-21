package controllers

import (
	"api_test/middlewares"
	"api_test/models"
	"api_test/webserver"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func CheckAccount(c models.User) bool {
	var data string
	err := webserver.DBCon.QueryRow("SELECT first_name FROM public.users where first_name = $1 and password =$2;", c.Username, c.Password).Scan(&data)
	return err == nil

}

func RegisterAccount(c echo.Context) error {
	var id_check string
	var user models.User
	var dt = time.Now()
	datetime := dt.Format(time.RFC3339Nano)
	// var id_return int
	err_bind := c.Bind(&user)
	if err_bind != nil {
		return c.String(http.StatusBadRequest, "bad request params")
	}
	err := webserver.DBCon.QueryRow("SELECT id FROM public.users where first_name = $1 or email = $2;", user.Username, user.Email).Scan(&id_check)

	if err == nil {
		return c.JSON(http.StatusBadRequest, "Register Account Error")
	}
	//
	if id_check == "" {

		jwt, err := middlewares.JwtTest(user.Password)
		if err != nil {
			log.Fatalf("JWT Error%s", err)
		}
		query := "INSERT INTO users (first_name, last_name, email, gender, password, token,created) VALUES ($1,$2,$3,$4,$5,$6,$7) ;"
		insertResult, err := webserver.DBCon.Exec(query, user.Username, user.Password, user.Email, "", "", jwt, datetime)
		if err != nil {
			log.Fatalf("impossible insert Users: %s", err)
		}
		// id, err := insertResult.LastInsertId()
		if err != nil {
			log.Fatalf("impossible to retrieve last inserted id: %s", err)
		}
		log.Printf("inserted id: %d", insertResult)
		return c.JSON(http.StatusOK, "Registed")
	}
	//
	return c.JSON(http.StatusOK, "Register Failed")
}

func Login(c echo.Context) error {
	// User ID from path `users/:id`
	// id := c.Param("id")
	res := &models.JsonReturn{
		Data:   nil,
		Status: http.StatusOK,
	}

	var user models.User
	err := c.Bind(&user)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	if user.Username != "" && user.Password != "" {
		auth_account := CheckAccount(user)
		if auth_account {

			jwt, err := middlewares.JwtTest(user.Username)
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

		} else {
			return c.JSON(http.StatusOK, &models.JsonReturn{
				Data:   "ไม่พบข้อมูล",
				Status: http.StatusOK,
			})
		}
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
