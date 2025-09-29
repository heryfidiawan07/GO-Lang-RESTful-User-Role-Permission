package controller

import (
	"app/config"
	"app/helper"
	"app/models"
	"app/request"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/danilopolani/gocialite/structs"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Redirect to correct oAuth URL
func RedirectHandler(c *gin.Context) {
	// Retrieve provider from route
	provider := c.Param("provider")

	// In this case we use a map to store our secrets, but you can use dotenv or your framework configuration
	// for example, in revel you could use revel.Config.StringDefault(provider + "_clientID", "") etc.
	providerSecrets := map[string]map[string]string{
		"github": {
			"clientID":     os.Getenv("CLIENT_ID_GITHUB"),
			"clientSecret": os.Getenv("CLIENT_SECRETS_GITHUB"),
			"redirectURL":  os.Getenv("AUTH_REDIRECT_URL") + "/github/callback/",
		},
		"google": {
			"clientID":     os.Getenv("CLIENT_ID_GOOGLE"),
			"clientSecret": os.Getenv("CLIENT_SECRETS_GOOGLE"),
			"redirectURL":  os.Getenv("AUTH_REDIRECT_URL") + "/google/callback/",
		},
	}

	providerScopes := map[string][]string{
		// "github": []string{"public_repo"},
		"github": []string{},
		"google": []string{},
	}

	providerData := providerSecrets[provider]
	actualScopes := providerScopes[provider]
	authURL, err := config.Gocial.New().
		Driver(provider).
		Scopes(actualScopes).
		Redirect(
			providerData["clientID"],
			providerData["clientSecret"],
			providerData["redirectURL"],
		)

	// Check for errors (usually driver not valid)
	if err != nil {
		c.Writer.Write([]byte("Error: " + err.Error()))
		return
	}

	// Redirect with authURL
	c.Redirect(http.StatusFound, authURL)
}

// Handle callback of provider
func CallbackHandler(c *gin.Context) {
	// Retrieve query params for state and code
	state := c.Query("state")
	code := c.Query("code")
	provider := c.Param("provider")

	// Handle callback and check for errors
	user, token, err := config.Gocial.Handle(state, code)
	if err != nil {
		c.Writer.Write([]byte("Error: " + err.Error()))
		return
	}

	// Print in terminal user information
	fmt.Printf("%#v", token)
	fmt.Printf("%#v", user)
	fmt.Printf("%#v", provider)

	var newUser = getOrRegisterUser(provider, user)
	var newToken = createToken(&newUser)

	c.JSON(200, gin.H{
		"token": newToken,
		"data":  newUser,
	})

	// If no errors, show provider name
	// c.Writer.Write([]byte("Hi, " + user.FullName))
}

func Login(c *gin.Context) {
	var valid request.Login
	if err := c.ShouldBind(&valid); err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.First(&user, "username = ?", valid.Username).Error; err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": "Username / password invalid!"})
		return
	}

	validPassword := helper.ComparePassword(user.Password, []byte(valid.Password))
	if !validPassword {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": "Username / password invalid!"})
		return
	}

	refreshTokenData := models.RefreshToken{
		Revoked:   false,
		ExpiredAt: time.Now().AddDate(0, 0, 7),
		UserId:    user.Id,
	}

	if err := config.DB.Create(&refreshTokenData).Error; err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": err})
		return
	}

	var token = createToken(&user)
	var refreshToken = createRefreshToken(&refreshTokenData)

	result := struct {
		token         string
		refresh_token string
		user          models.User
	}{token, refreshToken, user}

	c.JSON(404, gin.H{"status": true, "data": result, "message": nil})
}

func Register(c *gin.Context) {
	var valid request.Register
	if err := c.ShouldBind(&valid); err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": err.Error()})
		return
	}

	data := models.User{
		Username: valid.Username,
		Name:     valid.Name,
		Email:    valid.Email,
		Password: helper.HashAndSalt([]byte(valid.Password)),
	}

	if err := config.DB.Create(&data).Error; err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": err})
		return
	}

	c.JSON(201, gin.H{"status": true, "data": data, "message": "success"})
}

func getOrRegisterUser(provider string, user *structs.User) models.User {
	var userData models.User

	// config.DB.Where("provider = ? AND social_id = ? ", provider, user.ID).First(&userData)
	// config.DB.Where(&models.User{Provider: provider, SocialId: user.ID}).First(&userData)

	if err := config.DB.Where("provider = ? AND social_id = ? ", provider, user.ID).First(&userData).Error; err != nil {
		fmt.Println("# # Register # #")
		fmt.Println("# # # # # err", err)
		fmt.Println("# # Register # #")

		newUser := models.User{
			Username: user.Username,
			Name:     user.FullName,
			Email:    user.Email,
			SocialId: user.ID,
			Provider: provider,
			Avatar:   user.Avatar,
		}

		config.DB.Create(&newUser)
		return newUser
	}

	fmt.Println("* * Login * *")
	fmt.Println("* * * * * userData", userData)
	fmt.Println("* * Login * *")

	return userData
}

func RefreshToken(c *gin.Context) {
	var valid request.RefreshToken
	if err := c.ShouldBind(&valid); err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": err.Error()})
		return
	}

	token, _ := jwt.Parse(valid.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("refresh_id", claims["refresh_id"])
		var refreshToken models.RefreshToken
		if err := config.DB.First(&refreshToken, "id = ?", claims["refresh_id"]).Error; err != nil {
			c.JSON(404, gin.H{"status": false, "data": nil, "message": "Refresh token not found!"})
			return
		}

		if refreshToken.Revoked {
			c.JSON(404, gin.H{"status": false, "data": nil, "message": "Refresh token has been revoked!"})
			return
		}

		var user models.User
		if err := config.DB.First(&user, "id = ?", refreshToken.UserId).Error; err != nil {
			c.JSON(404, gin.H{"status": false, "data": nil, "message": "User not found!"})
			return
		}

		c.JSON(200, gin.H{
			"status":  true,
			"message": "success",
			"data": map[string]interface{}{
				"token": createToken(&user),
			},
		})
	} else {
		// fmt.Println(err)
		c.JSON(422, gin.H{"status": false, "data": nil, "message": "Invalid Refresh Token!"})
		c.Abort()
		return
	}
}

func RevokeRefreshToken(c *gin.Context) {
	var valid request.RefreshToken
	if err := c.ShouldBind(&valid); err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.First(&user, "id = ?", c.MustGet("jwt_user_id")).Error; err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": "Data not found!"})
		return
	}

	token, _ := jwt.Parse(valid.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("refresh_id", claims["refresh_id"])
		var refreshToken models.RefreshToken
		if err := config.DB.First(&refreshToken, "id = ?", claims["refresh_id"]).Error; err != nil {
			c.JSON(404, gin.H{"status": false, "data": nil, "message": "Refresh token not found!"})
			return
		}

		if refreshToken.Revoked {
			c.JSON(404, gin.H{"status": false, "data": nil, "message": "Refresh token has been revoked!"})
			return
		}

		if refreshToken.UserId != user.Id {
			c.JSON(404, gin.H{"status": false, "data": nil, "message": "Refresh token not math with auth id!"})
			return
		}

		data := models.RefreshToken{
			Revoked: true,
		}

		if err := config.DB.Model(&refreshToken).Updates(&data).Error; err != nil {
			c.JSON(404, gin.H{"status": false, "data": nil, "message": err})
			return
		}

		c.JSON(200, gin.H{"status": true, "data": nil, "message": "success"})
	} else {
		// fmt.Println(err)
		c.JSON(422, gin.H{"status": false, "data": nil, "message": "Invalid Refresh Token!"})
		c.Abort()
		return
	}
}

func createToken(user *models.User) string {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.Id,
		// "exp": time.Now().Add(time.Minute * 15).Unix(), // expired_at 15 minutes
		"exp": time.Now().AddDate(0, 0, 1).Unix(), // expired_at 1 days
		"iat": time.Now().Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		fmt.Println("tokenString err", err)
	}
	return tokenString
}

func createRefreshToken(refreshToken *models.RefreshToken) string {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"status":  true,
		"message": "success",
		"data": map[string]interface{}{
			"refresh_id": refreshToken.Id,
			"exp":        time.Now().AddDate(0, 0, 2).Unix(), // expired_at 2 days
			"iat":        time.Now().Unix(),
		},
	})

	refreshTokenString, err := jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		fmt.Println("refreshTokenString err", err)
	}
	return refreshTokenString
}
