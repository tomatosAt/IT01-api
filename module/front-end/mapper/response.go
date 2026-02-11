package mapper

import (
	"github.com/tomatosAt/IT01-api/model"
	"github.com/tomatosAt/IT01-api/module/front-end/dto"
)

func ResponseUserInsertMapper(user model.User) dto.DataInsertUser {
	return dto.DataInsertUser{
		ID:           user.Id.String(),
		CreateTimeAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
