package users

import (
	"fmt"
	"log"
	_ "time"

	"github.com/trungkien71297/go_api_management/datasources/mysql/user_db"
	"github.com/trungkien71297/go_api_management/utils/errors"
)

const (
	insertQuery = "INSERT INTO user_account(username, first_name, last_name, email, pwd) values(?,?,?,?,?);"
	iq2         = "INSERT INTO user_account(username, first_name, last_name, email, pwd) values (?,?,?,?,?);"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Save() *errors.RestError {
	fmt.Println(iq2)
	fmt.Println(user_db.Client.Stats())
	if err := user_db.Client.Ping(); err != nil {

		fmt.Println("Co bi chi mo")
	}

	stmt, err := user_db.Client.Prepare(iq2)
	if err != nil {
		return &errors.RestError{
			Message: "Cant insert",
			Error:   "Ble ble",
			Code:    301,
		}
	}

	defer stmt.Close()

	//	now := time.Now()
	//	user.DateCreated = now.Format("2006-01-02T15:04:05.000")

	log.Println(user.Lastname)
	insertRes, err := stmt.Exec(user.Username, user.Firstname, user.Lastname, user.Email, user.Password)
	if err != nil {
		return &errors.RestError{
			Message: "Cant insert",
			Error:   "Ble ble",
			Code:    301,
		}

	}
	user.Id, err = insertRes.LastInsertId()

	if err != nil {
		return &errors.RestError{
			Message: "Cant insert",
			Error:   "Ble ble",
			Code:    301,
		}
	}
	return nil
}

func Get(userId int64) (*User, *errors.RestError) {
	res := userDB[userId]

	if err := user_db.Client.Ping(); err != nil {
		panic(err)
	}

	if res == nil {
		return nil, &errors.RestError{}
	}
	return res, nil
}
