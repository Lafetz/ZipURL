package web

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
	"github.com/lafetz/url-shortner/internal/core/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockUsersService struct {
	mock.Mock
}

func (srv *mockUsersService) GetUser(username string) (*domain.User, error) {

	return domain.NewUser(username, "username@email.com", []byte("password")), nil
}
func (srv *mockUsersService) AddUser(user *domain.User) (*domain.User, error) {

	return user, nil
}

func (srv *mockUsersService) DeleteUser(id uuid.UUID) error {
	return nil
}

func TestCreateUser(t *testing.T) {
	mockService := &mockUsersService{}
	createUserHandler := createUser(mockService)
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/signup", createUserHandler)
	t.Run("Request Missing Json", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodPost, "/signup", nil)
		if err != nil {
			log.Fatal(err)
		}
		router.ServeHTTP(w, req)
		assert.Equal(t, 400, w.Code)
		// assert.Equal(t, , w.Body.String())
	})
	t.Run("Email,Username and Password Required", func(t *testing.T) {
		jsonParam := `{}`
		json := strings.NewReader(string(jsonParam))
		w := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodPost, "/signup", json)
		if err != nil {
			log.Fatal(err)
		}
		router.ServeHTTP(w, req)
		assert.Equal(t, 422, w.Code)

	})

	t.Run("Invalid email", func(t *testing.T) {
		jsonParam := `{	 "username":"wxosoitsorlxd",
		"email":"dsxssoasfdgmail.com",
		"password":"87s654321"	}`
		json := strings.NewReader(string(jsonParam))
		w := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodPost, "/signup", json)
		if err != nil {
			log.Fatal(err)
		}
		router.ServeHTTP(w, req)
		assert.Equal(t, 422, w.Code)
	})

	t.Run("Password too short", func(t *testing.T) {
		jsonParam := `{	 "username":"wxosoitsorlxd",
		"email":"dsxssoasfd@gmail.com",
		"password":"1234567"	}`
		json := strings.NewReader(string(jsonParam))
		w := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodPost, "/signup", json)
		if err != nil {
			log.Fatal(err)
		}
		router.ServeHTTP(w, req)
		assert.Equal(t, 422, w.Code)

	})
	t.Run("Successfuly create User", func(t *testing.T) {
		jsonParam := `{	 "username":"wxosoitsorlxd",
		"email":"dsxssoasfd@gmail.com",
		"password":"12345678"	}`
		json := strings.NewReader(string(jsonParam))
		w := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodPost, "/signup", json)
		if err != nil {
			log.Fatal(err)
		}
		router.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)

	})
}
