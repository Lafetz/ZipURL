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
)

type mockUsersService struct {
}

func (srv *mockUsersService) GetUser(username string) (*domain.User, error) {
	hashP, err := hashPassword("password")
	if err != nil {
		log.Fatal(err)
	}
	return domain.NewUser(username, "username@email.com", hashP), nil
}
func (srv *mockUsersService) AddUser(user *domain.User) (*domain.User, error) {

	return user, nil
}

func (srv *mockUsersService) DeleteUser(id uuid.UUID) error {
	return nil
}

func TestCreateUser(t *testing.T) {
	mockService := mockUsersService{}
	createUserHandler := createUser(&mockService)
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/signup", createUserHandler)
	tests := []struct {
		name string
		json *strings.Reader
		code int
	}{
		{
			name: "Email,Username and Password Required",
			json: strings.NewReader(string(`{}`)),
			code: 422,
		}, {
			name: "Invalid email",
			json: strings.NewReader(string(`{	 "username":"wxosoitsorlxd",
			 	"email":"dsxssoasfdgmail.com",
				"password":"letsgooguel"	}`)),
			code: 422,
		}, {
			name: "Password too short",
			json: strings.NewReader(string(`{	 "username":"wxosoitsorlxd",
			 	"email":"dsxssoasfd@gmail.com",
				"password":"lets"	}`)),
			code: 422,
		}, {
			name: "Successfuly create User",
			json: strings.NewReader(string(`{	 "username":"wxosoitsorlxd",
			"email":"dsxssoasfd@gmail.com",
				"password":"password"	}`)),
			code: 201,
		}}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodPost, "/signup", tt.json)
			req.Header.Add("Content-Type", "application/json")
			if err != nil {
				log.Fatal(err)
			}
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.code, w.Code)

		})
	}

}
func TestSignIn(t *testing.T) {
	mockService := &mockUsersService{}
	createUserHandler := signin(mockService)
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/signin", createUserHandler)
	tests := []struct {
		name string
		json *strings.Reader
		code int
	}{
		{
			name: "Username and Password Required",
			json: strings.NewReader(string(`{}`)),
			code: 422,
		}, {
			name: "Invalid password",
			json: strings.NewReader(string(`{	 "username":"wxosoitsorlxd",
			 	"email":"dsxssoasfdgmail.com",
				"password":"letsgooguel"	}`)),
			code: 401,
		}, {
			name: "Successfuly sign in",
			json: strings.NewReader(string(`{"username":"wxosoitsorlxd",
				"password":"password"	}`)),
			code: 200,
		}}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			w := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodPost, "/signin", tt.json)
			req.Header.Add("Content-Type", "application/json")
			if err != nil {
				log.Fatal(err)
			}
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.code, w.Code)

		})
	}

}
