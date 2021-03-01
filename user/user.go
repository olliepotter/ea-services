package user

import (
	"fmt"

	"github.com/olliepotter/ea-services/eadb"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// User describes a user object in the database
type User struct {
	eadb.BaseUUID
	Email           string `json:"email" gorm:"unique"`
	EmailVerified   bool   `json:"email_verified"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	PrimaryRider    int    `json:"primary_rider"`
	CustomerID      string `json:"customer_id"`
	EmContactName   string `json:"em_contact_name"`
	EmContactNumber string `json:"em_contact_number"`
	RoleID          int    `json:"role_id"`
}

// HOOK
func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	tx.Model(u).Update("email", "rileypotter17@gmail.com")
	fmt.Println(u.ID)
	return
}

// HOOK
func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
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
	err := db.Create(&user)
	if err.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error,
		})
	}
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
