package model

type User struct {
	Id       int
	Username string
	Age      int
	Email    string
	AddTime  int
	Password string
}

func (user User) TableName() string {
	return "user" //config table name as 'user', because the default name is users
}
