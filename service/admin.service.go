package service

import (
	"github.com/Peter-Sanders/go_webserver/db"
	// "golang.org/x/crypto/bcrypt"
)

func NewAdminServices(u User, uStore db.Store) *AdminServices {

	return &AdminServices{
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

type AdminServices struct {
	User      User
	UserStore db.Store
}

func (us *AdminServices) Login(username string) (User, error) {

	query := db.Get_sql("admin/get_user_by_username")

	stmt, err := us.UserStore.Db.Prepare(query)
	if err != nil {
		return User{}, err
	}

	defer stmt.Close()

	us.User.Username = username
	err = stmt.QueryRow(
		us.User.Username,
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
