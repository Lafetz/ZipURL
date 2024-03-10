package jwtauth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lafetz/url-shortner/internal/core/domain"
)

type UserToken struct {
	Id       string
	Email    string
	Username string
}

type UserClaim struct {
	jwt.RegisteredClaims
	Id       string
	Email    string
	Username string
}

func (u *UserClaim) GetUserToken() *UserToken {
	return &UserToken{
		Id:       u.Id,
		Email:    u.Email,
		Username: u.Email,
	}

}

var (
	ErrInvalidToken = errors.New("token not valid")
)

func CreateJwt(user *domain.User) (string, error) {
	KEY := os.Getenv("JWT_TOKEN")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim{RegisteredClaims: jwt.RegisteredClaims{Issuer: "github.com/lafetz/snippitstash",
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Audience:  jwt.ClaimStrings{"github.com/lafetz/ZipUrl"},
		NotBefore: jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour))},
		Id:       user.Id.String(),
		Email:    user.Email,
		Username: user.Username})

	jwtToken, err := token.SignedString([]byte(KEY)) //

	return jwtToken, err

}
func PareseJwt(jwtToken string) (*UserClaim, error) {
	KEY := os.Getenv("JWT_TOKEN")
	token, err := jwt.ParseWithClaims(jwtToken, &UserClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(KEY), nil
	})
	if err != nil {
		return nil, ErrInvalidToken
	}

	if claims, ok := token.Claims.(*UserClaim); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
