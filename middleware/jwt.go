package middleware

import (
	"app/config"
	"app/models"
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Auth(action string) gin.HandlerFunc {
	return check(action)
}

func permission(action string, c *gin.Context) bool {
	// Special cases that don't need database check
	if action == "me" || action == "except" {
		return true
	}

	userID := c.MustGet("jwt_user_id")

	// Check if user is superadmin
	var user models.User
	if err := config.DB.First(&user, "id = ?", userID).Error; err == nil {
		if user.Username == "superadmin" {
			return true
		}
	}

	// Single query to check if user has the specific permission
	var count int64
	err := config.DB.Table("permissions").
		Joins("JOIN role_permissions ON permissions.id = role_permissions.permission_id").
		Joins("JOIN users ON role_permissions.role_id = users.role_id").
		Where("users.id = ? AND permissions.name = ?", userID, action).
		Count(&count).Error

	if err != nil {
		fmt.Printf("Error checking permission: %v\n", err)
		return false
	}

	return count > 0
}

func check(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		bearerToken := strings.Split(authorization, " ")

		if len(bearerToken) != 2 {
			c.JSON(401, gin.H{"message": "Unauthorized !"})
			c.Abort()
			return
		}

		// token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
		token, _ := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// fmt.Println(claims["user_id"], claims["username"])
			fmt.Println("user_id", claims["user_id"])
			var user models.User
			if err := config.DB.First(&user, "id = ?", claims["user_id"]).Error; err != nil {
				c.JSON(404, gin.H{"message": "Data not found !"})
				c.Abort()
				return
			}

			c.Set("jwt_user_id", claims["user_id"])
			fmt.Println("**** action **** ", permission(action, c))

			if !permission(action, c) {
				c.JSON(401, gin.H{"message": "Permission denied !"})
				c.Abort()
				return
			}
		} else {
			// fmt.Println(err)
			c.JSON(422, gin.H{"message": "Invalid Token !"})
			c.Abort()
			return
		}
	}
}
