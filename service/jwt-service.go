package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Contract for JWTService
type JWTService interface {
	GenerateToken(UserID string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

// instance of JWTService
func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "bookstore",
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	fmt.Println("secretKey => ", secretKey)

	if secretKey != "" {
		secretKey = "sangatrahasiaheyy"
	}

	return secretKey
}

func (j *jwtService) GenerateToken(UserID string) string {
	claims := &jwtCustomClaims{
		UserID,
		jwt.StandardClaims{
			// ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			// expire 1 tahun
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		fmt.Println("error Genereate Token => ", err)
		panic(err)
	}
	return t
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}
