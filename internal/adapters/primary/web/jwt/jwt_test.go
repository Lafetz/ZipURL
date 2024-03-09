package jwt_auth

import (
	"log"
	"testing"

	"github.com/lafetz/url-shortner/internal/core/domain"
	"github.com/stretchr/testify/assert"
)

func TestJwt(t *testing.T) {
	t.Run("Successfuly create jwt", func(t *testing.T) {
		user := domain.NewUser("username", "email@Email.com", []byte("stuff"))
		token, err := CreateJwt(user)
		if err != nil {
			log.Fatal(err)
		}
		assert.True(t, len(token) > 0)
		assert.Nil(t, err)

	})

	t.Run("Successfuly parse token", func(t *testing.T) {
		user := domain.NewUser("username", "email@Email.com", []byte("stuff"))
		token, err := CreateJwt(user)
		if err != nil {
			log.Fatal(err)
		}
		_, err = PareseJwt(token)
		assert.Nil(t, err)
	})

}
