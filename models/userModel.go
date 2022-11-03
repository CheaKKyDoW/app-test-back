package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// type User struct {
// 	Name string `gorm:"<-:create"`          // allow read and create
// 	Name string `gorm:"<-:update"`          // allow read and update
// 	Name string `gorm:"<-"`                 // allow read and write (create and update)
// 	Name string `gorm:"<-:false"`           // allow read, disable write permission
// 	Name string `gorm:"->"`                 // readonly (disable write permission unless it configured)
// 	Name string `gorm:"->;<-:create"`       // allow read and create
// 	Name string `gorm:"->:false;<-:create"` // createonly (disabled read from db)
// 	Name string `gorm:"-"`                  // ignore this field when write and read with struct
// 	Name string `gorm:"-:all"`              // ignore this field when write, read and migrate with struct
// 	Name string `gorm:"-:migration"`        // ignore this field when migrate with struct
// }
