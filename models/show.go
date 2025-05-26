package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Show struct {
	Id        uint32
	Name      string
	TicketUrl string
	StartDate datatypes.Date
	EndDate   datatypes.Date
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
