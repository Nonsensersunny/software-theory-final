package model

import "time"

type Prediction struct {
	Id       string    `gorm:"primary_key" json:"id"`
	Aid      string    `json:"aid"`
	Did      string    `json:"did"`
	Accuracy float32   `json:"accuracy"`
	Path     string    `json:"path"`
	Result   string    `json:"result"`
	Time     time.Time `gorm:"default:now()" json:"time"`
}
