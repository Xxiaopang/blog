package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"project/config"
	"time"
)

type MyClaims struct {
	UID   uint `json:"uid"`
	Level int  `json:"level"`
	jwt.RegisteredClaims
}

const (
	issuer  = "video"
	subject = "login"
)

var key = []byte(config.JWTKey)

func CreateToken(uid uint, level int) (string, error) {
	expires := time.Now().Add(time.Hour * 48)
	claims := MyClaims{
		UID:   uid,
		Level: level,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,
			Subject:   subject,
			Audience:  nil,
			ExpiresAt: jwt.NewNumericDate(expires),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}

func CheckToken(token string) (*MyClaims, bool) {
	claims, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if myClaims, ok := claims.Claims.(*MyClaims); ok && claims.Valid {
		return myClaims, true
	}
	return nil, false
}
func RefreshClaims(claims *MyClaims) (string, error) {
	expires := time.Now().Add(7 * 24 * time.Hour)
	newclaims := &MyClaims{
		UID:   claims.UID,
		Level: claims.Level,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,
			Subject:   subject,
			Audience:  nil,
			ExpiresAt: jwt.NewNumericDate(expires),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, newclaims)
	return token.SignedString(key)

}
