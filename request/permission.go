package request

type PermissionStore struct {
	ParentMenu string `form:"parent_menu" json:"parent_menu" xml:"parent_menu"`
	ParentId   string `form:"parent_id" json:"parent_id" xml:"parent_id"`
	Name       string `form:"name" json:"name" xml:"name" binding:"required"`
	Alias      string `form:"alias" json:"alias" xml:"alias" binding:"required"`
	Url        string `form:"url" json:"url" xml:"url"`
	Icon       string `form:"icon" json:"icon" xml:"icon"`
}
