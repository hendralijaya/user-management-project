package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(userId string, Username string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaims struct {
	UserId string `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer string
}

func NewJWTService() JWTService {
	return &jwtService{
		issuer: "golang-jwt",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		secretKey = "md5(rahasia)"
	}
	return secretKey
}	

func (j *jwtService) GenerateToken(UserId string, Username string) (string, error) {
	claims := &jwtCustomClaims{
		UserId,
		Username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			Issuer:    j.issuer,
			IssuedAt: time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}