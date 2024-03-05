package web

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lafetz/url-shortner/internal/core/domain"
)

type UserClaim struct {
	jwt.RegisteredClaims
	ID       string
	Email    string
	Username string
}

var (
	ERRINVALIDTOKEN = errors.New("token not valid")
)

func createJwt(user *domain.User) (string, error) {
	KEY := "temporaryKEy"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim{RegisteredClaims: jwt.RegisteredClaims{
		Issuer:    "github.com/lafetz/snippitstash",
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Audience:  jwt.ClaimStrings{"github.com/lafetz/snippitstash"},
		NotBefore: jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	},
		ID:    "1",
		Email: user.Email,

		Username: "First Last"})

	jwtToken, err := token.SignedString([]byte(KEY)) //
	return jwtToken, err

}
func pareseJwt(jwtToken string) error {
	KEY := "temporaryKEy"
	var userClaim UserClaim
	token, err := jwt.ParseWithClaims(jwtToken, userClaim, func(token *jwt.Token) (interface{}, error) {
		return []byte(KEY), nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return ERRINVALIDTOKEN
	}
	return nil
}
