package domain

type RolePermissions struct {
	RoleID  string   `json:"role_id"`
	PermIDs []string `json:"perm_id"`
}
