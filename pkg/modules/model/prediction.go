package model

import "time"

type Prediction struct {
	Id       string    `gorm:"primary_key" json:"id"`
	Aid      string    `json:"aid"`
	TrainId  string    `json:"train_id"`
	TestId   string    `json:"test_id"`
	Accuracy float64   `json:"accuracy"`
	Path     string    `json:"path"`
	Result   string    `json:"result"`
	Time     time.Time `gorm:"default:now()" json:"time"`
}
