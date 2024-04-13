package service

import (
	bannerapp "github.com/ds124wfegd/banner-app"
	"github.com/ds124wfegd/banner-app/pkg/repository"
)

type BannerListService struct {
	repo repository.BannerList
}

func NewBannerListService(repo repository.BannerList) *BannerListService {
	return &BannerListService{repo: repo}
}

func (s *BannerListService) Create(list bannerapp.Banner) (int, error) {
	return s.repo.Create(list)
}

func (s *BannerListService) GetAll() ([]bannerapp.Banner, error) {
	return s.repo.GetAll()
}

func (s *BannerListService) GetByFeature(featureid int) ([]bannerapp.Banner, error) {
	return s.repo.GetByFeature(featureid)
}

func (s *BannerListService) GetByTag(tagid int) ([]bannerapp.Banner, error) {
	return s.repo.GetByTag(tagid)
}

func (s *BannerListService) DeleteByFeature(featureid int) error {
	return s.repo.DeleteByFeature(featureid)
}

func (s *BannerListService) UpdateBanner(bannerId int, input bannerapp.UpdateBanner) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateBanner(bannerId, input)
}
