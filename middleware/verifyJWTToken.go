package middleware

import (
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	// Mem-parse dan memverifikasi token menggunakan secret key
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	// Jika token tidak valid, atau ada kesalahan parsing
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, fmt.Errorf("signature tidak valid")
		}
		return nil, fmt.Errorf("token parsing error: %v", err)
	}

	// Jika token kadaluarsa
	if !token.Valid {
		return nil, fmt.Errorf("token tidak valid")
	}

	return claims, nil
}

func main() {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."

	// Verifikasi token
	claims, err := VerifyToken(tokenString)
	if err != nil {
		log.Printf("Token tidak valid: %v\n", err)
	} else {
		fmt.Printf("Token valid untuk pengguna: %s\n", claims.Username)
	}
}
