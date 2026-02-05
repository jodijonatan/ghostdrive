package models

import "time"

type Role string

const (
	RoleAdmin Role = "ADMIN"
	RoleUser  Role = "USER"
)

type User struct {
	ID       string `gorm:"primaryKey"`
	Email    string `gorm:"unique"`
	Password string
	Role     Role
}

type FileMeta struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Filename  string    `json:"filename"`
	OwnerID   string    `json:"owner_id"`
	CreatedAt time.Time `json:"created_at"`
}
