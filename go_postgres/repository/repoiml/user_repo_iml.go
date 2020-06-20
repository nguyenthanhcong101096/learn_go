package repoiml

import (
	"database/sql"
	"fmt"
	models "go_postgres/model"
	repo "go_postgres/repository"
)

type UserRepoIml struct {
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) repo.UserRepo {
	return &UserRepoIml{
		Db: db,
	}
}

func (u *UserRepoIml) Select() ([]models.User, error) {
	users := make([]models.User, 0)

	rows, err := u.Db.Query("SELECT * FROM users")

	if err != nil {
		return users, err
	}

	for rows.Next() {
		user := models.User{}
		fmt.Println(rows)
		err := rows.Scan(&user.Id, &user.Email, &user.Role)

		if err != nil {
			break
		}

		users = append(users, user)
	}

	err = rows.Err()

	if err != nil {
		return users, err
	}

	return users, nil
}
