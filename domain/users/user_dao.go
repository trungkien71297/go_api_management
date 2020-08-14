package users

import (
	"github.com/trungkien71297/go_api_management/datasources/mysql/user_db"
	"github.com/trungkien71297/go_api_management/utils/errors"
	"log"
	"time"
)

const (
	insertQuery    = "INSERT INTO user_account(username, first_name, last_name, email, pwd, date_created) values(?,?,?,?,?,?);"
	searchQuery    = "SELECT * FROM user_account where id=?;"
	searchAllQuery = "SELECT * FROM user_account;"
	editUser       = "UPDATE user_account SET  username=?, first_name=?, last_name=?, email=?, pwd=? WHERE id=?"
	deleteUser     = "DELETE FROM user_account WHERE id=?"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Save() *errors.RestError {
	if err := user_db.Client.Ping(); err != nil {
		panic(err)
	}

	stmt, err := user_db.Client.Prepare(insertQuery)
	if err != nil {
		return &errors.RestError{
			Message: "Cant insert",
			Error:   "Ble ble",
			Code:    301,
		}
	}

	defer stmt.Close()

	now := time.Now()
	//user.DateCreated = now.Format("2006-01-02T15:04:05.000")

	user.DateCreated = now
	log.Println(user.Lastname)
	insertRes, err := stmt.Exec(user.Username, user.Firstname, user.Lastname, user.Email, user.Password, user.DateCreated)
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
	var user User
	if err := user_db.Client.Ping(); err != nil {
		panic(err)
	}

	st := user_db.Client.QueryRow(searchQuery, userId)
	st.Scan(&user.Id, &user.Username, &user.Firstname, &user.Lastname, &user.Email, &user.Password, &user.DateCreated)
	if st.Err() != nil {
		return nil, &errors.RestError{
			Message: "Cant find",
			Error:   "Ble ble",
			Code:    301,
		}
	}

	return &user, nil
}

func GetAllUser() ([]*User, *errors.RestError) {
	var (
		users []*User
	)

	if err := user_db.Client.Ping(); err != nil {
		panic(err)
	}

	rows, err := user_db.Client.Query(searchAllQuery)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.Username, &user.Firstname, &user.Lastname, &user.Email, &user.Password, &user.DateCreated)
		users = append(users, &user)
	}
	return users, nil
}

func (user *User) EditUser() (bool, *errors.RestError) {
	if err := user_db.Client.Ping(); err != nil {
		panic(err)
	}

	res, err := user_db.Client.Exec(editUser, user.Username, user.Firstname, user.Lastname, user.Email, user.Password, user.Id)
	if err != nil {
		panic(err)
	}
	if r, _ := res.RowsAffected(); r > 0 {
		return true, nil
	}
	return false, nil
}

func (user *User) DeletetUser() (bool, *errors.RestError) {
	if err := user_db.Client.Ping(); err != nil {
		panic(err)
	}

	res, err := user_db.Client.Exec(deleteUser, user.Id)
	if err != nil {
		panic(err)
	}
	if r, _ := res.RowsAffected(); r > 0 {
		return true, nil
	}
	return false, nil
}
