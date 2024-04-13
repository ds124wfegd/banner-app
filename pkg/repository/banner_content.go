package repository

import (
	"fmt"

	bannerapp "github.com/ds124wfegd/banner-app"
	"github.com/jmoiron/sqlx"
)

type bannerContentPostgres struct {
	db *sqlx.DB
}

func NewBannerContentPostgres(db *sqlx.DB) *bannerContentPostgres {
	return &bannerContentPostgres{db: db}
}

func (r *bannerContentPostgres) CreateContent(bannerId int, content bannerapp.Content) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createContentQuery := fmt.Sprintf("INSERT INTO %s ( banner_id, title, some_title, text, some_text, some_url ) VALUES ($1, $2, $3, $4, $5,$6) RETURNING content.id", contentTable)
	row := tx.QueryRow(createContentQuery, content.BannerId, content.Title, content.SomeTitle, content.Text, content.SomeText, content.SomeUrl)
	err = row.Scan(&id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *bannerContentPostgres) GetAllContent(bannerId int) ([]bannerapp.Content, error) {
	var content []bannerapp.Content
	query := fmt.Sprintf(`SELECT content.id, content.banner_id, content.title, content.some_title, content.text, content.some_text, content.some_url FROM %s content WHERE content.banner_id = %d `, contentTable, bannerId)
	if err := r.db.Select(&content, query); err != nil {
		return nil, err
	}

	return content, nil
}

func (r *bannerContentPostgres) DeleteContent(bannerId int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE content.id=$1`, contentTable)
	_, err := r.db.Exec(query, bannerId)
	return err
}
