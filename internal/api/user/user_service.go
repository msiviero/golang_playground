package user

import (
	"github.com/jmoiron/sqlx"

	pb "dev.msiviero/example/internal/grpc_gen"

	"go.uber.org/zap"
)

type UserService struct {
	db *sqlx.DB
}

func NewUserService(db *sqlx.DB) UserService {
	return UserService{db: db}
}

func (userService UserService) GetUser(id int32) pb.UserMessage {

	user := User{}
	err := userService.db.Get(&user, "SELECT * FROM users WHERE id=?", id)

	if err != nil {
		zap.L().Error(err.Error())
	}

	return pb.UserMessage{
		Id:    user.Id,
		Age:   user.Age,
		Name:  user.Name,
		Email: user.Email,
	}
}

func (userService UserService) PutUser(user *pb.UserMessage) error {
	_, err := userService.db.NamedExec("INSERT INTO users (age, name, email) VALUES (:age, :name, :email)", User{
		Age:   user.Age,
		Name:  user.Name,
		Email: user.Email,
	})

	return err
}

type User struct {
	Id    int32  `db:"id"`
	Age   int32  `db:"age"`
	Name  string `db:"name"`
	Email string `db:"email"`
}
