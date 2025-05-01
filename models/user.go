package models

import "time"

type User struct {
	Id               uint32
	Name             string
	Password         string
	Avatar           string
	Email            string
	EmailVerified_at time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        time.Time
}
