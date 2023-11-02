package user

import (
	"github.com/jmoiron/sqlx"

	"go.uber.org/zap"
)

type UserService struct {
	db *sqlx.DB
}

func NewUserService(db *sqlx.DB) UserService {
	return UserService{db: db}
}

func (userService UserService) GetUsers(cb func(User)) {
	rows, err := userService.db.Queryx("SELECT * FROM users")

	if err != nil {
		zap.L().Error(err.Error())
	}

	for rows.Next() {
		user := User{}

		err = rows.StructScan(&user)

		if err != nil {
			zap.L().Error(err.Error())
			continue
		}
		cb(user)
	}
}

func (userService UserService) GetUser(id int32) (User, error) {

	user := User{}
	err := userService.db.Get(&user, "SELECT * FROM users WHERE id=?", id)

	if err != nil {
		zap.L().Error(err.Error())
	}

	return user, err
}

func (userService UserService) PutUser(user User) error {
	_, err := userService.db.NamedExec("INSERT INTO users (age, name, email) VALUES (:age, :name, :email)", user)
	return err
}

type User struct {
	Id    int32  `db:"id"`
	Age   int32  `db:"age"`
	Name  string `db:"name"`
	Email string `db:"email"`
}
