package service

import (
	bannerapp "github.com/ds124wfegd/banner-app"
	"github.com/ds124wfegd/banner-app/pkg/repository"
)

type BannerContentService struct {
	repo        repository.BannerList
	contentRepo repository.BannerContent
}

func NewBannerContentService(repo repository.BannerList, contentRepo repository.BannerContent) *BannerContentService {
	return &BannerContentService{repo: repo, contentRepo: contentRepo}

}

func (s *BannerContentService) CreateContent(bannerId int, content bannerapp.Content) (int, error) {
	//	_, err := s.contentRepo.GetBannerById(bannerId)
	//	if err != nil {
	//		return 0, err
	//	}

	return s.contentRepo.CreateContent(bannerId, content)
}
func (s *BannerContentService) GetAllContent(bannerId int) ([]bannerapp.Content, error) {
	return s.contentRepo.GetAllContent(bannerId)
}

func (s *BannerContentService) DeleteContent(contentId int) error {
	return s.contentRepo.DeleteContent(contentId)
}

/*	Create(bannerId int, content bannerapp.Content) (int, error)




func (s *MpItemService) GetById(userId, itemId int) (mp.MpItem, error) {
	return s.repo.GetById(userId, itemId)
}


}
DeleteContent(bannerId int) error
func (s *MpItemService) Update(userId, itemId int, input mp.UpdateItemInput) error {
	return s.repo.Update(userId, itemId, input)
}*/
