package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

var SecretKey = "Secret_key_f3Nfienfeo33ieuooreuna;ae9383n3283084njnkj2"

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webToken, err := token.SignedString([]byte(SecretKey))

	return webToken, err
}

func VerifyToken(authorizationHeader string) (*jwt.Token, error) {
	// Pastikan header Authorization tidak kosong
	if authorizationHeader == "" {
		return nil, fmt.Errorf("Authorization header is empty")
	}

	// Pisahkan header menjadi dua bagian: "Bearer" dan token
	parts := strings.Split(authorizationHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return nil, fmt.Errorf("Invalid Authorization header format")
	}

	// Ambil token dari bagian kedua
	tokenString := parts[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SecretKey), nil
	})

	if err != nil {
		return nil, err
	}
	return token, nil
}

func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	token, err := VerifyToken(tokenString)
	if err != nil {
		return nil, err
	}
	claims, isOk := token.Claims.(jwt.MapClaims)

	if isOk && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("Invalid token")
}
