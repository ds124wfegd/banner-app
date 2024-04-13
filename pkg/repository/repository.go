package repository

import (
	bannerapp "github.com/ds124wfegd/banner-app"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user bannerapp.User) (int, error)
	GetUser(username, password string) (bannerapp.User, error)
	CreateAdmin(admin bannerapp.Admin) (int, error)
	GetAdmin(adminUsername, adminPassword string) (bannerapp.Admin, error)
}

type BannerList interface {
	Create(list bannerapp.Banner) (int, error)
	GetAll(adminId int) ([]bannerapp.Banner, error)
	GetByFeature(bannerid int) ([]bannerapp.Banner, error)
	GetByTag(bannerid int) ([]bannerapp.Banner, error)

	DeleteByFeature(featureid int) error
	UpdateBanner(bannerId int, input bannerapp.UpdateBanner) error
}

type BannerContent interface {
	CreateContent(bannerId int, content bannerapp.Content) (int, error)
	GetAllContent(bannerId int) ([]bannerapp.Content, error)
	DeleteContent(contentId int) error
}

type Repository struct {
	Authorization
	BannerList
	BannerContent
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		BannerList:    NewBannerListPostgres(db),
		BannerContent: NewBannerContentPostgres(db),
	}
}
