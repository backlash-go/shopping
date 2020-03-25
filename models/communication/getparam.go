package communication

type GetLoginParam struct {
	Id    int `json:"id"`
	Cellphone string  `json:"cellphone"`
	Password string `json:"password"`
	Method string  `json:"method"`
}
