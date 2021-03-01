package user

import (
	"fmt"
	"time"

	"github.com/olliepotter/ea-services/eadb"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type GormCommon struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type User struct {
	GormCommon
	Email  string `json:"email"`
	RoleID int    `json:"role_id"`
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	tx.Model(u).Update("email", "rileypotter17@gmail.com")
	fmt.Println(u.ID)
	return
}

func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Println(u)
	fmt.Println(u.Email)
	return
}

// GetUser retrieves a user from the database
func GetUser(c *fiber.Ctx) error {
	db := eadb.DBConn
	var user []User
	db.Find(&user)
	return c.JSON(user)
}

// CreateUser ....
func CreateUser(c *fiber.Ctx) error {
	db := eadb.DBConn
	var user User
	user.Email = "olliepotter16@gmail.com"
	user.RoleID = 0
	db.Create(&user)
	return c.JSON(user)
}

// DeleteUser ....
func DeleteUser(c *fiber.Ctx) error {
	db := eadb.DBConn
	var user User
	db.First(&user, 22)
	db.Delete(&user, 22)
	return c.SendStatus(fiber.StatusOK)
}
