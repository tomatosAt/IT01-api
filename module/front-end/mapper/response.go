package mapper

import (
	"time"

	"github.com/tomatosAt/IT01-api/model"
	"github.com/tomatosAt/IT01-api/module/front-end/dto"
)

func ResponseUserInsertMapper(user model.User) dto.DataInsertUser {
	return dto.DataInsertUser{
		ID:           user.Id.String(),
		CreateTimeAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ResponseDashBoardUserMapper(user []model.User) dto.ResponseDataDashBoardUser {
	var dashboardUser []dto.DataDashBoardUser
	for _, user := range user {
		dashboardUser = append(dashboardUser, dto.DataDashBoardUser{
			ID:          user.Id.String(),
			FirstNameTH: user.FirstNameTH,
			LastNameTH:  user.LastNameTH,
			FullNameTH:  user.FirstNameTH + " " + user.LastNameTH,
			Dob:         user.BirthDate.Format("2006-01-02"),
			Age:         CalculateAge(user.BirthDate),
			Addresses:   user.Address,
		})
	}
	return dto.ResponseDataDashBoardUser{
		AllUsers: dashboardUser,
		Total:    len(user),
	}
}

func ResponseUserByIDMapper(user *model.User) dto.DataDashBoardUser {
	return dto.DataDashBoardUser{
		ID:          user.Id.String(),
		FirstNameTH: user.FirstNameTH,
		LastNameTH:  user.LastNameTH,
		FullNameTH:  user.FirstNameTH + " " + user.LastNameTH,
		Dob:         user.BirthDate.Format("2006-01-02"),
		Age:         CalculateAge(user.BirthDate),
		Addresses:   user.Address,
	}
}

func CalculateAge(birthDate time.Time) int {
	now := time.Now()
	age := now.Year() - birthDate.Year()

	if now.YearDay() < birthDate.YearDay() {
		age--
	}
	return age
}
