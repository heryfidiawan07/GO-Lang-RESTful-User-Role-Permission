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
	if action == "except" {
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

		// Check if Authorization header exists
		if authorization == "" {
			c.JSON(401, gin.H{"status": false, "message": "Authorization header is required"})
			c.Abort()
			return
		}

		bearerToken := strings.Split(authorization, " ")

		// Check if Bearer token format is correct
		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
			c.JSON(401, gin.H{"status": false, "message": "Invalid authorization header format. Expected: Bearer <token>"})
			c.Abort()
			return
		}

		// Check if token is not empty
		if bearerToken[1] == "" {
			c.JSON(401, gin.H{"status": false, "message": "Token cannot be empty"})
			c.Abort()
			return
		}

		// Parse and validate JWT token
		token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
			// Validate signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		// Handle JWT parsing errors
		if err != nil {
			var errorMessage string
			if ve, ok := err.(*jwt.ValidationError); ok {
				switch {
				case ve.Errors&jwt.ValidationErrorExpired != 0:
					errorMessage = "Token has expired"
				case ve.Errors&jwt.ValidationErrorNotValidYet != 0:
					errorMessage = "Token is not valid yet"
				case ve.Errors&jwt.ValidationErrorMalformed != 0:
					errorMessage = "Token is malformed"
				default:
					errorMessage = "Token validation failed"
				}
			} else {
				errorMessage = "Invalid token format"
			}
			c.JSON(401, gin.H{"status": false, "message": errorMessage})
			c.Abort()
			return
		}

		// Check if token is valid and extract claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Check if user_id exists in claims
			userID, exists := claims["user_id"]
			if !exists || userID == "" {
				c.JSON(401, gin.H{"status": false, "message": "Token missing user_id claim"})
				c.Abort()
				return
			}

			fmt.Println("user_id", userID)

			// Verify user exists in database
			var user models.User
			if err := config.DB.First(&user, "id = ?", userID).Error; err != nil {
				c.JSON(404, gin.H{"status": false, "message": "User not found"})
				c.Abort()
				return
			}

			// Set user_id in context for further use
			c.Set("jwt_user_id", userID)
			fmt.Println("**** action **** ", permission(action, c))

			// Check permissions
			if !permission(action, c) {
				c.JSON(403, gin.H{"status": false, "message": "Access denied: insufficient permissions"})
				c.Abort()
				return
			}
		} else {
			c.JSON(401, gin.H{"status": false, "message": "Invalid token"})
			c.Abort()
			return
		}
	}
}
