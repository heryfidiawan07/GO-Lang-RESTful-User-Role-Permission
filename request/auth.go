package request

type Login struct {
	Username string `form:"username" json:"username" xml:"username" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

type Register struct {
	Name     string `form:"name" json:"name" xml:"name" binding:"required"`
	Username string `form:"username" json:"username" xml:"username" binding:"required"`
	Email    string `form:"email" json:"email" xml:"email" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

type RefreshToken struct {
	RefreshToken string `form:"refresh_token" json:"refresh_token" xml:"refresh_token" binding:"required"`
}
