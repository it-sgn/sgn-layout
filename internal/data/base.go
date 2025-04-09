package data

import (
	"time"

	"gorm.io/gorm"
)

type BaseFields struct {
	Id        int64          `gorm:"primarykey;type:int;comment:primary key id"`
	CreatedAt time.Time      `gorm:"column:created_at;not null;comment:creation time"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;comment:Update time"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;comment:Delete time"`
}

// Paginate
func Paginate(page, pageSize int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(int(offset)).Limit(int(pageSize))
	}
}
