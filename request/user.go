package request

type Login struct {
	Username string `form:"username" json:"username" xml:"username"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required"`
}

type Register struct {
	Name string `form:"name" json:"name" xml:"name"  binding:"required"`
	Username string `form:"username" json:"username" xml:"username"  binding:"required"`
	Email string `form:"email" json:"email" xml:"email"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required"`
}

type ChangePassword struct {
	OldPassword string `form:"old_password" json:"old_password" xml:"old_password"  binding:"required"`
	NewPassword string `form:"new_password" json:"new_password" xml:"new_password"  binding:"required"`
}

type UserStore struct {
	Name string `form:"name" json:"name" xml:"name"  binding:"required"`
	Username string `form:"username" json:"username" xml:"username"  binding:"required"`
	Email string `form:"email" json:"email" xml:"email"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required"`
}

type UserUpdate struct {
	Name string `form:"name" json:"name" xml:"name"  binding:"required"`
	Username string `form:"username" json:"username" xml:"username"  binding:"required"`
	Email string `form:"email" json:"email" xml:"email"  binding:"required"`
}
