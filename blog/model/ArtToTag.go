package model

import "gorm.io/gorm"

type ArtToTag struct {
	Aid uint `json:"aid"`
	Tid uint `json:"tid"`
	gorm.Model
}
