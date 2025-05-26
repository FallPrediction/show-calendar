package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Event struct {
	Id            uint32
	Name          string
	ShowId        uint32
	OgImage       string
	OgUrl         string
	OgTitle       string
	OgDescription string
	StartDate     datatypes.Date
	EndDate       datatypes.Date
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}
