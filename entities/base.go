package entities

import "time"

type Model struct {
	ID        int64     `gorm:"primaryKey autoIncrement" json:"id"`
	UUID      string    `gorm:"type:char(36);unique" json:"uuid"`
	CreatedAt time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP;onUpdate:CURRENT_TIMESTAMP" json:"updated_at"`
	IsDeleted bool      `gorm:"type:tinyint(1);default:0" json:"is_deleted"`
	IsActive  bool      `gorm:"type:tinyint(1);default:1" json:"is_active"`
}
