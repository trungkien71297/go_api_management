package user_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

const (
	mysql_username = "musn"
	mysql_password = "mpwd"
	mysql_host     = "msh"
	mysql_schema   = "msc"
)

var (
	Client *sql.DB
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	iq2 := "INSERT INTO user_account(username, first_name, last_name, email, pwd) values (?,?,?,?,?);"
	var (
		username = os.Getenv(mysql_username)
		password = os.Getenv(mysql_password)
		host     = os.Getenv(mysql_host)
		schema   = os.Getenv(mysql_schema)
	)
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, schema)

	fmt.Println(dataSourceName)
	Client, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	var res sql.Result
	res, err = Client.Exec(iq2, "U", "a", "y", "s", "sc", "as")
	if err != nil {
		fmt.Println("Dc day")
	}

	_ = res
	log.Println("Connected to db")
}
