package utils

import (
	"ginapi/pkg/setting"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(setting.JwtSecret)

// Claims this is for Claims
type Claims struct {
	LanguageType string  `json:"language_type"`
	ProductID    int     `json:"product_id"`
	UserID       string  `json:"user_id"`
	TransferID   string  `json:"transfer_id"`
	TotalAmount  float64 `json:"total_amount"`
	jwt.StandardClaims
}

// GenerateToken this is GenerateToken
func GenerateToken(LanguageType string, ProductID int, UserID string, TransferID string, TotalAmount float64) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	claims := Claims{
		LanguageType,
		ProductID,
		UserID,
		TransferID,
		TotalAmount,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin.com",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// ParseToken this is ParseToken
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
