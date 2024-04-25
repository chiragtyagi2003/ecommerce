package services

import (
	"github.com/Sebas3270/ecommerce-app-backend/db"
	"github.com/Sebas3270/ecommerce-app-backend/helpers"
	"github.com/Sebas3270/ecommerce-app-backend/models"
	"github.com/gofiber/fiber/v2"
)

type LogInStruct struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

// func LogIn(c *fiber.Ctx) error {

// 	l := new(LogInStruct)

// 	if err := c.BodyParser(l); err != nil {
// 		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
// 	}

// 	var user models.User
// 	db.MysqlConnection.Where("email = ?", l.Email).First(&user)

// 	if user.Password != l.Password {
// 		return c.Status(fiber.StatusNotFound).SendString("Check credentials")
// 	}

// 	token, err := helpers.GenerateJWT(user.ID)

// 	if err != nil {
// 		return c.Status(fiber.StatusNotFound).SendString("Server error, check logs")
// 	}

// 	return c.JSON(fiber.Map{
// 		"user":  user,
// 		"token": token,
// 	})
// }

func LogIn(c *fiber.Ctx) error {

	l := new(LogInStruct)

	if err := c.BodyParser(l); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	var user models.User
	err := db.MysqlConnection.Where("email = ?", l.Email).First(&user).Error
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if !helpers.CheckPasswordHash(l.Password, *user.Password) {
		return c.Status(fiber.StatusNotFound).SendString("Check credentials")
	}

	token, err := helpers.GenerateToken(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("There was an error")
	}

	user.Password = nil
	// user.Items = nil

	return c.JSON(fiber.Map{
		"user":  user,
		"token": token,
	})
}

func Resgister(c *fiber.Ctx) error {

	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	var checkUser models.User
	db.MysqlConnection.Where("email = ?", user.Email).First(&checkUser)
	if checkUser != (models.User{}) {
		return c.Status(fiber.StatusBadRequest).SendString("Email already logged in, use another one")
	}

	pass, err := helpers.HashPassword(*user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Server error")
	}
	*user.Password = pass

	if err := db.MysqlConnection.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Server error")
	}

	token, err := helpers.GenerateToken(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Server error")
	}

	user.Password = nil
	// user.Items = nil

	return c.JSON(fiber.Map{
		"user":  user,
		"token": token,
	})
}

func RenewToken(c *fiber.Ctx) error {

	userId := helpers.GetTokenData(c)

	var user models.User
	err := db.MysqlConnection.Where("id = ?", userId).First(&user).Error

	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	token, err := helpers.GenerateToken(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Server error")
	}

	user.Password = nil

	return c.JSON(fiber.Map{
		"user":  user,
		"token": token,
	})

}
