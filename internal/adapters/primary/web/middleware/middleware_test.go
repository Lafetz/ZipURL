package middleware

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	jwt_auth "github.com/lafetz/url-shortner/internal/adapters/primary/web/jwt"
	"github.com/lafetz/url-shortner/internal/core/domain"
	"github.com/stretchr/testify/assert"
)

func TestRequireAuth(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/", RequireAuth())
	t.Run("Returns error if cookie is missing", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodPost, "/", nil)
		if err != nil {
			log.Fatal(err)
		}
		router.ServeHTTP(w, req)
		assert.Equal(t, 401, w.Code)

	})
	t.Run("Returns error if token is invalid", func(t *testing.T) {
		w := httptest.NewRecorder()
		cookie := &http.Cookie{
			Name:  "Authorization",
			Value: "token",
		}
		req, err := http.NewRequest(http.MethodPost, "/", nil)
		req.AddCookie(cookie)

		if err != nil {
			log.Fatal(err)
		}
		router.ServeHTTP(w, req)
		assert.Equal(t, 401, w.Code)

	})
	t.Run("Successful if token is valid", func(t *testing.T) {
		w := httptest.NewRecorder()
		user := domain.NewUser("username", "email@Email.com", []byte("stuff"))
		token, err := jwt_auth.CreateJwt(user)
		if err != nil {
			log.Fatal(err)
		}
		cookie := &http.Cookie{
			Name:  "Authorization",
			Value: token,
		}
		req, err := http.NewRequest(http.MethodPost, "/", nil)
		req.AddCookie(cookie)

		if err != nil {
			log.Fatal(err)
		}
		router.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)

	})

}
