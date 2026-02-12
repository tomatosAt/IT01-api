package services

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/tomatosAt/IT01-api/model"
	"github.com/tomatosAt/IT01-api/module/front-end/dto"
	"github.com/tomatosAt/IT01-api/module/front-end/mapper"
	"github.com/tomatosAt/IT01-api/pkg/util"
	"gorm.io/gorm"
)

func (s *Service) UserSVC(ctx context.Context, data dto.UserPayload) (*dto.DataInsertUser, int, error) {
	var result dto.DataInsertUser
	dobDate, err := time.Parse("2006-01-02", data.Dob)
	if err != nil {
		return &result, http.StatusBadRequest, errors.New("invalid date format")
	}
	dobYear := dobDate.Year()
	currentYear := time.Now().Year()
	if dobYear > currentYear {
		return &result, http.StatusBadRequest, errors.New("birth year cannot be in the future")
	}
	// ENCRYPT ข้อมูล
	encryptedData, _ := util.EncryptList(s.repo.AppCfg().Secret.EncryptKey, data.FirstNameTH, data.LastNameTH)
	if err := s.repo.GetUserByFullNameAndDobRepo(ctx, nil, encryptedData[0], encryptedData[1], data.Dob); err != nil {
		return &result, http.StatusBadRequest, err
	} else {
		mapperPreRegister, err := mapper.InsertUserMapper(data, encryptedData)
		if err != nil {
			return &result, http.StatusBadRequest, err
		}
		userRepo, err := s.repo.InsertUserRepo(ctx, nil, mapperPreRegister)
		if err != nil {
			return &result, http.StatusInternalServerError, errors.New("server error")
		}
		result = mapper.ResponseUserInsertMapper(userRepo)
	}
	return &result, http.StatusOK, nil
}

func (s *Service) DashboardUser(ctx context.Context, limit, page int) (dto.ResponseDataDashBoardUser, int, error) {
	var result dto.ResponseDataDashBoardUser
	var user []model.User
	offset := (page - 1) * limit
	if err := s.repo.GetAllUserRepo(ctx, nil, &user, limit, offset); err != nil {
		return result, http.StatusBadRequest, err
	}
	dataDecrypt, err := s.DecryptDashboradUser(user)
	if err != nil {
		return result, http.StatusBadRequest, err
	}
	result = mapper.ResponseDashBoardUserMapper(dataDecrypt)
	return result, http.StatusOK, nil
}

func (s *Service) DecryptDashboradUser(userList []model.User) ([]model.User, error) {
	if len(userList) == 0 {
		return userList, nil
	}
	for i := range userList {
		decryptData, _ := util.DecryptList(
			s.repo.AppCfg().Secret.EncryptKey,
			userList[i].FirstNameTH,
			userList[i].LastNameTH,
		)
		userList[i].FirstNameTH = decryptData[0]
		userList[i].LastNameTH = decryptData[1]
	}
	return userList, nil
}

func (s *Service) DecryptUserByID(user *model.User) error {
	decryptData, err := util.DecryptList(
		s.repo.AppCfg().Secret.EncryptKey,
		user.FirstNameTH,
		user.LastNameTH,
	)
	if err != nil {
		return err
	}
	user.FirstNameTH = decryptData[0]
	user.LastNameTH = decryptData[1]
	return err
}

func (s *Service) UserDetailByIDSVC(ctx context.Context, userId string) (dto.DataDashBoardUser, int, error) {
	var res dto.DataDashBoardUser
	getUserByID, err := s.repo.GetUserByUserIDRepo(ctx, nil, userId)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return res, http.StatusInternalServerError, err
		}
	}
	if getUserByID != nil {
		res = mapper.ResponseUserByIDMapper(getUserByID)
	}
	if err := s.DecryptUserByID(getUserByID); err != nil {
		return res, http.StatusInternalServerError, err
	}
	res = mapper.ResponseUserByIDMapper(getUserByID)
	return res, http.StatusOK, nil
}
