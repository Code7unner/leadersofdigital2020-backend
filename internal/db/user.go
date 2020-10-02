package db

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Sex      string `json:"sex"`
	Role     string `json:"role"`
}

func (u User) GetId() int64 {
	return u.Id
}
