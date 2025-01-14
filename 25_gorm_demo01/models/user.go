package models

type User struct { //default tablename is 'users'
	Id       int
	Username string
	Age      int
	Email    string
	AddTime  int
}

func (user User) TableName() string {
	return "user" //config table name as 'user', because the default name is users
}
