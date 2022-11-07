package controllers

import (
	"api_test/webserver"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {
	// User ID from path `users/:id`
	// id := c.Param("id")
	id := c.FormValue("id")
	// return c.HTML(http.StatusOK, "<b>Thank you! "+id+"</b>")
	return c.String(http.StatusOK, id)

}

func Login(c echo.Context) error {

	user, err := webserver.DBCon.Query("SELECT * FROM mock_data limit 1")
	if err != nil {
		fmt.Println("err", err)
		return nil
	}

	cols, _ := user.Columns()

	data := make(map[string]string)

	if user.Next() {
		columns := make([]string, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}

		user.Scan(columnPointers...)

		for i, colName := range cols {
			data[colName] = columns[i]
		}
	}

	// Make a sample token
	// In a real world situation, this token will have been acquired from
	// some other API call (see Example_getTokenViaHTTP)
	// token, err := createToken("foo")
	// fatal(err)

	// // Make request.  See func restrictedHandler for example request processor
	// req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:%v/restricted", serverPort), nil)
	// fatal(err)
	// req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	// res, err := http.DefaultClient.Do(req)
	// fatal(err)

	// // Read the response body
	// buf := new(bytes.Buffer)
	// io.Copy(buf, res.Body)
	// res.Body.Close()
	// fmt.Println(buf.String())

	u := &Data_{
		Data:   data,
		Status: http.StatusOK,
	}

	// a, _ := json.Marshal(u)
	return c.JSON(http.StatusOK, u)

}

type Data_ struct {
	Data   map[string]string `json:"data"`
	Status int               `json:"status"`
}
