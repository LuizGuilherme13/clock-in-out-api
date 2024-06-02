package employee

type Model struct {
	Id       int64  `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
}
