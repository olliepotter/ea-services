package eadb

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// BaseID provides base fields for DB objects using auto-increment ID
type BaseID struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// BaseUUID provides base fields for DB objects using UUID
type BaseUUID struct {
	ID        uuid.UUID      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// BeforeCreate auto generates UUID
func (base *BaseUUID) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = uuid.NewV4()
	return
}

// Global database variable that can be passed around
var (
	DBConn *gorm.DB
)
