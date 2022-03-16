package entities

import (
	"gorm.io/gorm"
)

type User struct {
	Model
	Username string `gorm:"type:varchar(100);unique" json:"username"`
	Email    string `gorm:"type:varchar(100);unique" json:"email"`
	Password string `gorm:"type:varchar(255)" json:"password"`
	RoleID   int64  `gorm:"not null" json:"-"`
	Role     *Role  `gorm:"foreignkey:RoleID;constraint;onUpdate:RESTRICT,onDelete:RESTRICT" json:"user"`
}

func (User) TableName() string {
	return "user"
}

type UserRepository interface {
	SaveUser(db *gorm.DB) (*User, error)
	FindAllUsers(db *gorm.DB) (*[]User, error)
}
