package utils

import "github.com/golang-jwt/jwt/v5"

var SecretKey = "Secret_key_f3Nfienfeo33ieuooreuna;ae9383n3283084njnkj2"

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webToken, err := token.SignedString([]byte(SecretKey))

	return webToken, err
}
