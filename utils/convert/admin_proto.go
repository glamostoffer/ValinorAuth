package convert

import (
	"github.com/glamostoffer/ValinorAuth/internal/model"
	"github.com/glamostoffer/ValinorAuth/utils/mapper"
	adminProto "github.com/glamostoffer/ValinorProtos/auth/admin_auth"
)

func UserToProto(user model.User) *adminProto.User {
	return &adminProto.User{
		Id:        user.ID,
		Login:     user.Username,
		Role:      mapper.Roles[user.Role],
		CreatedAt: user.CreatedAt.Unix(),
		UpdatedAt: user.UpdatedAt.Unix(),
	}
}

func UsersToProto(users []model.User) []*adminProto.User {
	protoUsers := make([]*adminProto.User, 0, len(users))
	for _, user := range users {
		protoUsers = append(protoUsers, UserToProto(user))
	}
	return protoUsers
}
