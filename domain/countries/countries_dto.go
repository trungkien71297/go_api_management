package countries

type Country struct {
	Id     int64  `json:"id"`
	Phone  int32  `json:"phone"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}
