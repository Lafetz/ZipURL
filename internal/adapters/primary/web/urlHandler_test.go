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

type mockUrlService struct {
}

func (srv *mockUrlService) GetUrls(userId uuid.UUID) ([]*domain.Url, error) {
	urls := []*domain.Url{domain.NewUrl(userId, "test.com"), domain.NewUrl(userId, "test2.com")}
	return urls, nil
}
func (srv *mockUrlService) GetUrl(shortUrl string) (*domain.Url, error) {
	return domain.NewUrl(uuid.New(), "test.com"), nil
}
func (srv *mockUrlService) AddUrl(url *domain.Url) (*domain.Url, error) {
	id := uuid.New().String()
	truncatedID := id[:7]
	url.ShortUrl = truncatedID
	return url, nil
}
func (srv *mockUrlService) DeleteUrl(shorturl string, userId uuid.UUID) error {
	return nil
}

func TestCreateUrl(t *testing.T) {
	mockService := mockUrlService{}
	createUrlHandler := createUrl(&mockService)
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/urls", requireAuth(), createUrlHandler)
	user := domain.NewUser("username", "email@Email.com", []byte("stuff"))
	token, err := createJwt(user)
	if err != nil {
		log.Fatal(err)
	}
	cookie := &http.Cookie{
		Name:  "Authorization",
		Value: token,
	}
	tests := []struct {
		name string
		json *strings.Reader
		code int
	}{
		{
			name: "OriginalUrl Required",
			json: strings.NewReader(string(`{}`)),
			code: 422,
		}, {
			name: "Successfully create url",
			json: strings.NewReader(string(`{	 "originalUrl":"wxosoitsorlxd"}`)),
			code: 201,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodPost, "/urls", tt.json)
			req.Header.Add("Content-Type", "application/json")
			req.AddCookie(cookie)
			if err != nil {
				log.Fatal(err)
			}
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.code, w.Code)

		})
	}
	t.Run("Not authorized", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodPost, "/urls", strings.NewReader(string(`{}`)))
		req.Header.Add("Content-Type", "application/json")

		if err != nil {
			log.Fatal(err)
		}
		router.ServeHTTP(w, req)
		assert.Equal(t, 401, w.Code)
	})
}

func TestDeleteUrl(t *testing.T) {
	mockService := mockUrlService{}
	deleteUrlHandler := deleteUrl(&mockService)
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.DELETE("/urls/:shorturl", requireAuth(), deleteUrlHandler)
	user := domain.NewUser("username", "email@Email.com", []byte("stuff"))
	token, err := createJwt(user)
	if err != nil {
		log.Fatal(err)
	}
	cookie := &http.Cookie{
		Name:  "Authorization",
		Value: token,
	}
	t.Run("Not authorized", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodDelete, "/urls/sampleurl", nil)
		if err != nil {
			log.Fatal(err)
		}
		router.ServeHTTP(w, req)
		assert.Equal(t, 401, w.Code)
	})
	t.Run("successfully delete url", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodDelete, "/urls/sampleurl", nil)
		req.AddCookie(cookie)
		if err != nil {
			log.Fatal(err)
		}
		router.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
	})
}

func TestGetUrls(t *testing.T) {
	mockService := mockUrlService{}
	getUrlHandler := getUrls(&mockService)
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/urls", requireAuth(), getUrlHandler)
	user := domain.NewUser("username", "email@Email.com", []byte("stuff"))
	token, err := createJwt(user)
	if err != nil {
		log.Fatal(err)
	}
	cookie := &http.Cookie{
		Name:  "Authorization",
		Value: token,
	}
	t.Run("Not authorized", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "/urls", nil)
		if err != nil {
			log.Fatal(err)
		}
		router.ServeHTTP(w, req)
		assert.Equal(t, 401, w.Code)
	})
	t.Run("get urls", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "/urls", nil)
		req.AddCookie(cookie)
		if err != nil {
			log.Fatal(err)
		}
		router.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
	})
}
