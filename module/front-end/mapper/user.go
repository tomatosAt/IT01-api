package mapper

import (
	"time"

	"github.com/tomatosAt/IT01-api/model"
	"github.com/tomatosAt/IT01-api/module/front-end/dto"
)

func InsertUserMapper(payload dto.UserPayload, data []string) (model.User, error) {
	birthDate, err := time.Parse("02-01-2006", payload.Dob)
	if err != nil {
		return model.User{}, err
	}
	return model.User{
		FirstNameTH: data[0],
		LastNameTH:  data[1],
		BirthDate:   birthDate,
		Address:     payload.Addresses,
		Model: model.Model{
			CreatedBy: "Admin",
			UpdatedBy: "Admin",
		},
	}, nil
}
