package model

type Algorithm struct {
	Id          string `gorm:"primary_key" json:"id"`
	Name        string `gorm:"name" json:"name"`
	Path        string `json:"path"`
	Description string `json:"description"`
}
