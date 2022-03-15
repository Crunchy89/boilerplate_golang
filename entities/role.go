package entities

type Role struct {
	Model
	Role      string `gorm:"type:varchar(50)" json:"role"`
	CanRead   bool   `gorm:"type:tinyint(1);default:1" json:"can_read"`
	CanWrite  bool   `gorm:"type:tinyint(1);default:1" json:"can_write"`
	CanDelete bool   `gorm:"type:tinyint(1);default:1" json:"can_delete"`
	CanUpdate bool   `gorm:"type:tinyint(1);default:1" json:"can_update"`
}

func (Role) TableName() string {
	return "role"
}
