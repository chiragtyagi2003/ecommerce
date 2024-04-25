package helpers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	"time"
)

type TokenClaims struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}

var sampleSecretKey = []byte("MySecretKey")

func GenerateToken(userId uint) (string, error) {
	// Create the Claims
	claims := jwt.MapClaims{
		"id":  userId,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return t, nil
}

func GetTokenData(c *fiber.Ctx) uint {

	auth := c.Locals("auth").(*jwt.Token)
	claims := auth.Claims.(jwt.MapClaims)
	userId := uint(claims["id"].(float64))

	return userId
}

// func GenerateJWT(userId uint) (string, error) {

// 	expireToken := time.Now().Add(time.Hour * 2).Unix()

// 	// Set-up claims
// 	claims := TokenClaims{
// 		ID: userId,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: expireToken,
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	// claims := token.Claims.(jwt.MapClaims)
// 	// claims["exp"] = time.Now().Add(2 * time.Hour)
// 	// claims["authorized"] = true
// 	// claims["id"] = userId
// 	// claims["name"] = userName
// 	tokenString, err := token.SignedString(sampleSecretKey)
// 	if err != nil {
// 		return "", err
// 	}

// 	return tokenString, nil
// }

// func ValidateJwt(receivedToken string) (uint, error) {

// 	if receivedToken == "" {
// 		return 0, errors.New("No token in request")
// 	}

// 	var keyfunc jwt.Keyfunc = func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Unexpected signin method")
// 		}
// 		return []byte(sampleSecretKey), nil
// 	}

// 	claims := jwt.MapClaims{}
// 	_, err := jwt.ParseWithClaims(receivedToken, claims, keyfunc)

// 	if err != nil {
// 		return 0, err
// 	}

// 	id := uint(claims["id"].(float64))

// 	return id, nil

// }
