package service

import (
	bannerapp "github.com/ds124wfegd/banner-app"
	"github.com/ds124wfegd/banner-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user bannerapp.User) (int, error)
	GenerateToken(username, password string) (string, error)
	CreateAdmin(admin bannerapp.Admin) (int, error)
	GenerateTokenAdmin(adminUsername, adminPassword string) (string, error)
	ParseToken(token string) (int, error)
}

type BannerList interface {
	Create(list bannerapp.Banner) (int, error)
	GetAll() ([]bannerapp.Banner, error)
	GetByFeature(featureid int) ([]bannerapp.Banner, error)
	GetByTag(tagid int) ([]bannerapp.Banner, error)

	DeleteByFeature(featureid int) error
	UpdateBanner(bannerId int, input bannerapp.UpdateBanner) error
}

type BannerContent interface {
	CreateContent(bannerId int, content bannerapp.Content) (int, error)
	GetAllContent(bannerId int) ([]bannerapp.Content, error)
	DeleteContent(bannerId int) error
}

type Service struct {
	Authorization
	BannerContent
	BannerList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		BannerList:    NewBannerListService(repos.BannerList),
		BannerContent: NewBannerContentService(repos.BannerList, repos.BannerContent),
	}
}
