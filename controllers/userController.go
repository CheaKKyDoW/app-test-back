package controllers

import (
	"bytes"
	"fmt"
	"io"
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

	// Make a sample token
	// In a real world situation, this token will have been acquired from
	// some other API call (see Example_getTokenViaHTTP)
	token, err := createToken("foo")
	fatal(err)

	// Make request.  See func restrictedHandler for example request processor
	req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:%v/restricted", serverPort), nil)
	fatal(err)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	res, err := http.DefaultClient.Do(req)
	fatal(err)

	// Read the response body
	buf := new(bytes.Buffer)
	io.Copy(buf, res.Body)
	res.Body.Close()
	fmt.Println(buf.String())

	return c.String(http.StatusOK, "Login")

}
