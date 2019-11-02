package model

import "time"

type Dataset struct {
	Id          string    `gorm:"primary_key" json:"id"`
	Uid         string    `json:"uid"`
	Path        string    `json:"path"`
	Description string    `json:"description"`
	Time        time.Time `gorm:"default:now()" json:"time"`
}
