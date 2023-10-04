package repository

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type user struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *user {
	return &user{db}
}

func (repo user) Create(user models.User) (uint64, error) {
	statement, err := repo.db.Prepare(
		"insert into users (name, nickName, email, password) values (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.NickName, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastIsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(lastIsertID), nil
}

func (repo user) GetUsers(nameOrNickname string) ([]models.User, error) {
	nameOrNickname = fmt.Sprintf("%%%s%%", nameOrNickname)

	conn, err := repo.db.Query(
		"select id, name, nickName, email, createdAt from users where name like ? or nickName like ?",
		nameOrNickname, nameOrNickname,
	)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	var users []models.User

	for conn.Next() {
		var user models.User
		if err = conn.Scan(
			&user.ID,
			&user.Name,
			&user.NickName,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil

}

func (repo user) GetUser(id uint64) (models.User, error) {

	conn, err := repo.db.Query(
		"select id, name, nickName, email, createdAt from users where id = ?",
		id,
	)
	if err != nil {
		return models.User{}, err
	}

	defer conn.Close()
	var user models.User
	if conn.Next() {
		if err = conn.Scan(
			&user.ID,
			&user.Name,
			&user.NickName,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}
