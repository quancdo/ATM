package models

import (
	uuid "github.com/gofrs/uuid"
)

type Account struct {
	ID      uuid.UUID `sql:"type:varchar(36)" gorm:"primary_key"`
	Name    string    `sql:"type:varchar(50)"`
	Api_key string    `sql:"type:varchar(32)"`
	Balance int32
	Pin     string
}
