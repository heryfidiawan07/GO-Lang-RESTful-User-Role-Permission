package request

type RolePost struct {
	Name string `form:"name" json:"name" xml:"name"  binding:"required"`
	Permissions []string `form:"permissions" json:"permissions" xml:"permissions"`
}

type RolePut struct {
	Name string `form:"name" json:"name" xml:"name"  binding:"required"`
	Permissions []string `form:"permissions" json:"permissions" xml:"permissions"`
}