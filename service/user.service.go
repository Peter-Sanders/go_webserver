package service

import (
	"github.com/Peter-Sanders/go_webserver/db"
  "github.com/Peter-Sanders/go_webserver/util"
	"golang.org/x/crypto/bcrypt"
)

func NewUserServices(u User, uStore db.Store) *UserServices {

	return &UserServices{
		User:      u,
		UserStore: uStore,
	}
}

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
  FName string `json:"fname"`
  LName string `json:"lname"`
  Phone string `json:"phone"`
}

type UserServices struct {
	User      User
	UserStore db.Store
}

func (us *UserServices) CreateUser(u User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 8)
	if err != nil {
		return err
	}

  stmt := util.Get_sql("insert_user")

	_, err = us.UserStore.Db.Exec(
		stmt,
    u.FName,
    u.LName,
    u.Phone,
		u.Email,
		u.Username,
		string(hashedPassword),
	)

	return err
}

func (us *UserServices) CheckEmail(email string) (User, error) {

	query := util.Get_sql("get_user_by_email")

	stmt, err := us.UserStore.Db.Prepare(query)
	if err != nil {
		return User{}, err
	}

	defer stmt.Close()

	us.User.Email = email
	err = stmt.QueryRow(
		us.User.Email,
	).Scan(
		&us.User.ID,
    &us.User.FName,
    &us.User.LName,
    &us.User.Phone,
		&us.User.Email,
		&us.User.Password,
		&us.User.Username,
	)
	if err != nil {
		return User{}, err
	}

	return us.User, nil
}
