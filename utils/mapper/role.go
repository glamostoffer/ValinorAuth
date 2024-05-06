package mapper

import "github.com/glamostoffer/ValinorAuth/pkg/consts"

var (
	Roles = map[int]string{
		consts.AdminRoleID: consts.Admin,
		consts.UserRoleID:  consts.User,
	}
)
