package countries

import (
	"go_api_management/datasources/mysql/user_db"
	"go_api_management/utils/errors"
)

const (
	getAllQuery = "SELECT * FROM countries;"
)

func GetAllCountries() ([]*Country, *errors.RestError) {
	var countries []*Country

	if err := user_db.Client.Ping(); err != nil {
		panic(err)
	}

	res, err := user_db.Client.Query(getAllQuery)
	if err != nil {
		panic(err)
	}
	for res.Next() {
		country := Country{}
		res.Scan(&country.Id, &country.Phone, &country.Code, &country.Name, &country.Symbol)
		countries = append(countries, &country)
	}
	return countries, nil
}
