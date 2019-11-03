package model

type User struct {
	Id       string `gorm:"primary_key" json:"id"`
	Mail     string `json:"mail"`
	Password string `json:"password"`
}
