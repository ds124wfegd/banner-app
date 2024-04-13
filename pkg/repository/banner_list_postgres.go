package repository

import (
	"fmt"
	"strings"

	bannerapp "github.com/ds124wfegd/banner-app"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type BannerListPostgres struct {
	db *sqlx.DB
}

func NewBannerListPostgres(db *sqlx.DB) *BannerListPostgres {
	return &BannerListPostgres{db: db}
}

func (r *BannerListPostgres) Create(list bannerapp.Banner) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (is_active, feature_id, tag_id_1, tag_id_2, tag_id_3 ) VALUES ($1, $2, $3, $4, $5) RETURNING id", bannerTable)
	row := tx.QueryRow(createListQuery, list.IsActive, list.FeatureId, list.TagIds_1, list.TagIds_2, list.TagIds_3)

	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	/*createContentBannerQuery := fmt.Sprintf("INSERT INTO %s (feature_id, banner_id) VALUES ($1, $2)", featureBannerTable)
	_, err = tx.Exec(createContentBannerQuery, list.FeatureId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	}*/

	return id, tx.Commit()
}

func (r *BannerListPostgres) GetAll() ([]bannerapp.Banner, error) {
	var lists []bannerapp.Banner

	query := fmt.Sprintf(`SELECT banners.id, banners.is_active, banners.feature_id, banners.tag_id_1, banners.tag_id_2, banners.tag_id_3 FROM %s`, bannerTable)
	err := r.db.Select(&lists, query)
	return lists, err
}

func (r *BannerListPostgres) GetByFeature(featureid int) ([]bannerapp.Banner, error) {
	var lists []bannerapp.Banner

	query := fmt.Sprintf(`SELECT banners.id, banners.is_active, banners.feature_id, banners.tag_id_1, banners.tag_id_2, banners.tag_id_3 FROM %s WHERE banners.feature_id = $1 `, bannerTable)
	err := r.db.Select(&lists, query, featureid)

	return lists, err
}

func (r *BannerListPostgres) GetByTag(tagid int) ([]bannerapp.Banner, error) {
	var lists []bannerapp.Banner

	query := fmt.Sprintf("SELECT banners.id, banners.is_active, banners.feature_id, banners.tag_id_1, banners.tag_id_2, banners.tag_id_3 FROM %s WHERE (banners.tag_id_1 = $1 or  banners.tag_id_2 = $1 or  banners.tag_id_3 = $1)", bannerTable)
	err := r.db.Select(&lists, query, tagid)

	return lists, err
}

/*func (r *BannerListPostgres) GetById(userId, listId int) (mp.MpList, error) {
	var list mp.MpList

	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description, tl.filepath, tl.price FROM %s tl
								INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2`,
		mpListsTable, usersListsTable)
	err := r.db.Get(&list, query, userId, listId)

	return list, err
}*/

func (r *BannerListPostgres) DeleteByFeature(listId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE banners.feature_id =$1", bannerTable)
	_, err := r.db.Exec(query, listId)

	return err
}

func (r *BannerListPostgres) UpdateBanner(bannerId int, input bannerapp.UpdateBanner) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.IsActive != nil {
		setValues = append(setValues, fmt.Sprintf("is_active=$%d", argId))
		args = append(args, *input.IsActive)
		argId++
	}

	if input.FeatureId != nil {
		setValues = append(setValues, fmt.Sprintf("feature_id=$%d", argId))
		args = append(args, *input.FeatureId)
		argId++
	}

	if input.TagIds_1 != nil {
		setValues = append(setValues, fmt.Sprintf("tag_id_1=$%d", argId))
		args = append(args, *input.TagIds_1)
		argId++
	}

	if input.TagIds_2 != nil {
		setValues = append(setValues, fmt.Sprintf("tag_id_2=$%d", argId))
		args = append(args, *input.TagIds_2)
		argId++
	}

	if input.TagIds_3 != nil {
		setValues = append(setValues, fmt.Sprintf("tag_id_3=$%d", argId))
		args = append(args, *input.TagIds_3)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s banners SET %s WHERE banners.id=%d`,
		bannerTable, setQuery, bannerId)

	_, err := r.db.Exec(query, args...)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)
	return err
}
