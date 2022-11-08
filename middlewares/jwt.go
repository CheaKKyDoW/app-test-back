package middlewares

import (
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func JwtTest(c echo.Context) (string error) {

	// Set custom claims
	claims := &jwtCustomClaims{
		"Jon Snow",
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute + 1).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})

	// if err != nil {
	// 	v, _ := err.(*jwt.ValidationError)
	// 	switch v.Errors {
	// 	case jwt.ValidationErrorSignatureInvalid:
	// 		// token invalid
	// 		response := map[string]string{"message": "Unauthorized"}
	// 		helper.ResponseJSON(w, http.StatusUnauthorized, response)
	// 		return
	// 	case jwt.ValidationErrorExpired:
	// 		// token expired
	// 		response := map[string]string{"message": "Unauthorized, Token expired!"}
	// 		helper.ResponseJSON(w, http.StatusUnauthorized, response)
	// 		return
	// 	default:
	// 		response := map[string]string{"message": "Unauthorized"}
	// 		helper.ResponseJSON(w, http.StatusUnauthorized, response)
	// 		return
	// 	}
	// }

	// if !token.Valid {
	// 	response := map[string]string{"message": "Unauthorized"}
	// 	helper.ResponseJSON(w, http.StatusUnauthorized, response)
	// 	return
	// }

}
