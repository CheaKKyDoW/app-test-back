package models

import (
	"database/sql"
	"time"
)

type UserProfile struct {
	ID           uint
	Name         string
	Username     string `query:"username"`
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type User struct {
	Username string `query:"username"`
	Password string `query:"password"`
	Email    string `query:"email"`
	Token    string `query:"token"`
}
