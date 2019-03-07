package auth

import (
    "fmt"
	"io/ioutil"
	"github.com/dgrijalva/jwt-go"

)
func Validate() bool{
	b, err := ioutil.ReadFile("creds.txt")
    if err != nil {
        fmt.Print(err)
    }
    fmt.Println(string(b))
    token, _ := jwt.Parse(string(b), func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("There was an error")
        }
		return []byte("secret"), nil
        })
        
    if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true
	}
	return false
}
