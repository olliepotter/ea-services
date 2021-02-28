package eadb

import (
	"gorm.io/gorm"
)

// Global database variable that can be passed around
var (
	DBConn *gorm.DB
)
