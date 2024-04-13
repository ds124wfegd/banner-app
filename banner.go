package bannerapp

import "errors"

type Banner struct {
	Id        int `json:"id" db:"id"`
	IsActive  int `json:"isActive" db:"is_active" binding:"required"`
	FeatureId int `json:"featureId" db:"feature_id" binding:"required"`
	TagIds_1  int `json:"tagId_1" db:"tag_id_1" binding:"required"`
	TagIds_2  int `json:"tagId_2" db:"tag_id_2" binding:"required"`
	TagIds_3  int `json:"tagId_3" db:"tag_id_3" binding:"required"`
}

type Content struct {
	Id        int    `json:"id" db:"id"`
	BannerId  int    `json:"bannerId" db:"banner_id" binding:"required"`
	Title     string `json:"title" db:"title" binding:"required"`
	SomeTitle string `json:"someTitle" db:"some_title" binding:"required"`
	Text      string `json:"text" db:"text" binding:"required"`
	SomeText  string `json:"someText" db:"some_text" binding:"required"`
	SomeUrl   string `json:"someUrl" db:"some_url" binding:"required"`
}

type FeatureBanner struct {
	Id        int
	FeatureId int
	BannerId  int
}

type TagBanner struct {
	Id       int
	TagId    int
	BannerId int
}

type UpdateBanner struct {
	IsActive  *int `json:"isActive"`
	FeatureId *int `json:"featureId"`
	TagIds_1  *int `json:"tagId_1"`
	TagIds_2  *int `json:"tagId_2"`
	TagIds_3  *int `json:"tagId_3"`
}

func (i UpdateBanner) Validate() error {
	if i.IsActive == nil && i.FeatureId == nil && i.TagIds_1 == nil && i.TagIds_2 == nil && i.TagIds_3 == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
