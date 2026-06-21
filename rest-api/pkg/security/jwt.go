package security

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type (
	Ijwt interface {
		GenerateToken(email, username, id string, refresh bool) (string, error)
		ValidateToken(tokenString string, refresh bool) (jwt.MapClaims, error)
	}
	jwtImpl struct{}
)

func NewJwtImpl() Ijwt {
	return &jwtImpl{}
}

const tokenIssuer = "rest-api"

func (j *jwtImpl) GenerateToken(email, username, id string, refresh bool) (string, error) {

	jwtKey := os.Getenv("PASSWORD_TOKEN") // Asegúrate de configurar esta variable de entorno
	tokenType := "access"
	if refresh {
		jwtKey = os.Getenv("PASSWORD_REFRESH_TOKEN")
		tokenType = "refresh"
	}
	if jwtKey == "" {
		return "", errors.New("jwt key is not configured")
	}

	// Definir los claims
	claims := jwt.MapClaims{
		"email":    email,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // Expiración en 1 hora
		"iss":      tokenIssuer,
		"type":     tokenType,
		"username": username,
		"id":       id,
	}
	// Crear el token con el método de firma HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Firmar el token con una clave secreta
	secret := []byte(jwtKey)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		fmt.Println("Error al firmar el token:", err)
		return "", err
	}
	return tokenString, nil
}

func (j *jwtImpl) ValidateToken(tokenString string, refresh bool) (jwt.MapClaims, error) {

	jwtKey := os.Getenv("PASSWORD_TOKEN")
	if refresh {
		jwtKey = os.Getenv("PASSWORD_REFRESH_TOKEN")
	}
	if jwtKey == "" {
		return nil, errors.New("jwt key is not configured")
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	if !j.isTokenValidByExpiration(claims) {
		return nil, errors.New("token is expired")
	}

	if claims["iss"] != tokenIssuer {
		return nil, errors.New("invalid token issuer")
	}

	if refresh {
		if claims["type"] != "refresh" {
			return nil, errors.New("invalid token type")
		}
	} else {
		if claims["type"] != "access" {
			return nil, errors.New("invalid token type")
		}
	}

	return claims, nil
}

func (j *jwtImpl) isTokenValidByExpiration(claims jwt.MapClaims) bool {
	expirationTime, err := claims.GetExpirationTime()
	if err != nil || expirationTime == nil {
		return false
	}

	return expirationTime.After(time.Now())
}
