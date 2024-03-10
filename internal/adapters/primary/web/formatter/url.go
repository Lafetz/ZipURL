package jsonformatter

import (
	"time"

	"github.com/google/uuid"
	"github.com/lafetz/url-shortner/internal/core/domain"
)

type UrlResp struct {
	Id          uuid.UUID `json:"id"`
	UserId      uuid.UUID `json:"userId"`
	ShortUrl    string    `json:"shortUrl"`
	OriginalUrl string    `json:"originalUrl"`
	CreatedAt   time.Time `json:"createdAt"`
}

func NewUrlResp(url *domain.Url) *UrlResp {

	return &UrlResp{
		Id:          url.Id,
		UserId:      url.UserId,
		ShortUrl:    url.ShortUrl,
		OriginalUrl: url.OriginalUrl,
		CreatedAt:   url.CreatedAt,
	}
}
func NewUrlsResp(urls []*domain.Url) []*UrlResp {
	resp := []*UrlResp{}
	for _, url := range urls {

		resp = append(resp, NewUrlResp(url))
	}
	return resp
}
