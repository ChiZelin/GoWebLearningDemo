package models

type Usertable struct {
	Id       int
	Username string
	Age      int
	Email    string
	AddTime  int
}

func (u *Usertable) TableName() string {
	return "usertable"
}
