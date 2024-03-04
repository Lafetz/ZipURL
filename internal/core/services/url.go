package services

import (
	"strconv"

	"github.com/google/uuid"
	"github.com/lafetz/url-shortner/internal/core/domain"
	"github.com/lafetz/url-shortner/internal/core/ports"
)

type UrlService struct {
	repo ports.UrlRepository
}

func (srv *UrlService) GetUrls(userId uuid.UUID) ([]*domain.Url, error) {
	return srv.repo.GetUrls(userId)
}
func (srv *UrlService) GetUrl(shortUrl string) (*domain.Url, error) {
	return srv.repo.GetUrl(shortUrl)
}
func (srv *UrlService) AddUrl(url *domain.Url) (*domain.Url, error) {
	totalUrls, err := srv.repo.TotalUrls()
	if err != nil {
		return nil, err
	}
	//
	id := uuid.New().String()
	truncatedID := id[:7] + strconv.Itoa(totalUrls)
	//

	url.ShortUrl = truncatedID
	return srv.repo.AddUrl(url)
}
func (srv *UrlService) UpdateUrl(url *domain.Url) error {

	return srv.repo.UpdateUrl(url)
}
func (srv *UrlService) DeleteUrl(id uuid.UUID) error {
	return srv.repo.DeleteUrl(id)
}