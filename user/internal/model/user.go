package model

type User struct {
	Id     int64  `db:"id" json:"id"`
	Name   string `db:"name" json:"name"`
	Gender string `db:"gender" json:"gender"`
}

func (u User) TableName() any {
	return "user"
}
