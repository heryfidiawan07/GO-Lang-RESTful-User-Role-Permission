package request

type UserStore struct {
	Name string `form:"name" json:"name" xml:"name"  binding:"required"`
	Username string `form:"username" json:"username" xml:"username"  binding:"required"`
	Email string `form:"email" json:"email" xml:"email"  binding:"required"`
}

type UserUpdate struct {
	Name string `form:"name" json:"name" xml:"name"  binding:"required"`
	Username string `form:"username" json:"username" xml:"username"  binding:"required"`
	Email string `form:"email" json:"email" xml:"email"  binding:"required"`
}