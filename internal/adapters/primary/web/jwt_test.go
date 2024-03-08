package web

import (
	"fmt"
	"log"
	"testing"

	"github.com/lafetz/url-shortner/internal/core/domain"
	"github.com/stretchr/testify/assert"
)

func TestJwt(t *testing.T) {
	t.Run("Successfuly create jwt", func(t *testing.T) {
		user := domain.NewUser("username", "email@Email.com", []byte("stuff"))
		token, err := createJwt(user)
		if err != nil {
			log.Fatal(err)
		}
		assert.True(t, len(token) > 0)
		assert.Nil(t, err)
	})
	t.Run("Successfuly parse token", func(t *testing.T) {
		user := domain.NewUser("username", "email@Email.com", []byte("stuff"))
		token, err := createJwt(user)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", err)
		_, err = pareseJwt(token)
		assert.Nil(t, err)
	})
	// t.Run("Fails when Key is changed", func(t *testing.T) {
	// 	user := domain.NewUser("username", "email@Email.com", []byte("stuff"))
	// 	token, err := createJwt(user)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("%s", err)
	// 	_, err = pareseJwt(token)
	// 	assert.ErrorIs(t, err, ErrInvalidToken)
	// })
}
