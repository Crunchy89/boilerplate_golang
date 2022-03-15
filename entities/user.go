package entities

type User struct {
	Model
	Username string `gorm:"type:varchar(100);unique" json:"username"`
	Email    string `gorm:"type:varchar(100);unique" json:"email"`
	Password string `gorm:"type:varchar(255)" json:"password"`
	RoleID   int64  `gorm:"not null" json:"-"`
	Role     *Role  `gorm:"foreignkey:RoleID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}

func (User) TableName() string {
	return "user"
}
