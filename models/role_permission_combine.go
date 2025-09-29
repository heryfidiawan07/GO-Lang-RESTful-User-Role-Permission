package models

type RolePermissionCombine struct {
	Role
	Permissions []Permission
}
