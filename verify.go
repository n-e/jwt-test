package main

import (
	"fmt"

	"github.com/lestrrat-go/jwx/v2/jwt"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

func main() {
	tokenString := "eyJhbGciOiJFZERTQSJ9.eyJhIjoiYiIsImV4cCI6MTY1NTI0NzA2NC4xMjl9.Du6LnQMyr_gtAkrJzHZx5yvD3w0jXiZo09QJOVVlcgK8qVxyGz-fHoYMcicDJIwJB2WV-DFgJUAy1K8QZ_Y6AQ"

	src := "-----BEGIN PUBLIC KEY-----\nMCowBQYDK2VwAyEAW7rgytlpvkVL/Z/jCraGtJMNQM6d7EAXBMsNzHMpnwY=\n-----END PUBLIC KEY-----"

	key, err := jwk.ParseKey([]byte(src), jwk.WithPEM(true))
	if err != nil {
		fmt.Printf("failed to parse key in PEM format: %s\n", err)
		return
	}

	verifiedToken, err := jwt.Parse([]byte(tokenString), jwt.WithKey(jwa.EdDSA, key), jwt.WithValidate(true))
	if err != nil {
		fmt.Printf("failed to verify JWT: %s\n", err)
		return
	}
	fmt.Println(verifiedToken.PrivateClaims())
}
