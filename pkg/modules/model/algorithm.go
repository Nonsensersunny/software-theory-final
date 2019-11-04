package model

type Algorithm struct {
	Id          string `gorm:"primary_key" json:"id"`
	Path        string `json:"path"`
	Description string `json:"description"`
}
