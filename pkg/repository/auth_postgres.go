package repository

import (
	"errors"
	"fmt"
	"os"

	bannerapp "github.com/ds124wfegd/banner-app"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user bannerapp.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil

}

func (r *AuthPostgres) CreateAdmin(admin bannerapp.Admin) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (id, username, password_hash, admin_status, system_password) values ($1, $2, $3, $4, $5) RETURNING (id)", adminsTable)
	row := r.db.QueryRow(query, admin.Id, admin.AdminUsername, admin.AdminPassword, admin.AdminStatus, admin.SystemPasword)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	if admin.SystemPasword != os.Getenv("SYSTEM_PASSWORD") {
		err := errors.New("Access is denied - password_system is not correct")
		return 0, err
	} else {
		return id, nil
	}

}

func (r *AuthPostgres) GetUser(username, password string) (bannerapp.User, error) {
	var user bannerapp.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}

func (r *AuthPostgres) GetAdmin(adminUsername, adminPassword string) (bannerapp.Admin, error) {
	var admin bannerapp.Admin
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", adminsTable)
	err := r.db.Get(&admin, query, adminUsername, adminPassword)

	return admin, err
}
