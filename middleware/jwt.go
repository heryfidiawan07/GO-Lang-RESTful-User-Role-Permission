package middleware

import (
	"fmt"
	"os"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"app/config"
	"app/models"
)

func Auth(action string) gin.HandlerFunc {
	return check(action)
}

func permission(action string, c *gin.Context) bool {
	var user models.User
	config.DB.First(&user, "id = ?", c.MustGet("jwt_user_id"))

	var role models.Role
	config.DB.First(&role, "id = ?", user.RoleId)

	var rolePermissions []models.RolePermission
	config.DB.Where("role_id = ?", user.RoleId).Find(&rolePermissions)

	rolePermissionId := make([]string, len(rolePermissions))
	for key,value := range rolePermissions {
		rolePermissionId[key] = value.PermissionId
    }
	
	var permissions []models.Permission
	config.DB.Where("id IN ?", rolePermissionId).Find(&permissions)

	if action == "me" || action == "except" {
		return true
	}

	for _,value := range permissions {
		if(action == value.Name) {
			return true
		}
    }

	return false
}

func check(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		bearerToken   := strings.Split(authorization, " ")

		if len(bearerToken) != 2 {
			c.JSON(401, gin.H{
				"message": "Unauthorized !",
			})
			c.Abort()
			return
		}

		// token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
		token, _ := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// fmt.Println(claims["user_id"], claims["username"])
			fmt.Println("user_id", claims["user_id"])
			var user models.User
			if err := config.DB.First(&user, "id = ?", claims["user_id"]).Error; err != nil {
				c.JSON(404, gin.H{
					"message": "Data not found !",
				})
				c.Abort()
				return
			}
			
			c.Set("jwt_user_id", claims["user_id"])
			// c.Set("jwt_user_role", claims["user_role"])

			fmt.Println("**** action **** ", permission(action, c))
			if permission(action, c) == false {
				c.JSON(401, gin.H{
					"message": "Permission denied !",
				})
				c.Abort()
				return
			}
		} else {
			// fmt.Println(err)
			c.JSON(422, gin.H{
				"message": "Invalid Token !",
			})
			c.Abort()
			return
		}
	}
}