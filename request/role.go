package request

type RoleStore struct {
	Name        string   `form:"name" json:"name" xml:"name" binding:"required"`
	Desctiption string   `form:"description" json:"description" xml:"description"`
	Permissions []string `form:"permissions" json:"permissions" xml:"permissions"`
}

type RoleUpdate struct {
	Name        string   `form:"name" json:"name" xml:"name" binding:"required"`
	Desctiption string   `form:"description" json:"description" xml:"description"`
	Permissions []string `form:"permissions" json:"permissions" xml:"permissions"`
}
