package model

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

func (a *Algorithm) BeforeCreate(scope *gorm.Scope) error {
	id := uuid.NewV4()
	return scope.SetColumn("Id", id.String())
}

func (d *Dataset) BeforeCreate(scope *gorm.Scope) error {
	id := uuid.NewV4()
	return scope.SetColumn("Id", id.String())
}

func (p *Prediction) BeforeCreate(scope *gorm.Scope) error {
	id := uuid.NewV4()
	return scope.SetColumn("Id", id.String())
}

func (u *User) BeforeCreate(scope *gorm.Scope) error {
	id := uuid.NewV4()
	return scope.SetColumn("Id", id.String())
}
