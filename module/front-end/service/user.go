package services

import (
	"context"
	"errors"
	"net/http"

	"github.com/tomatosAt/IT01-api/module/front-end/dto"
	"github.com/tomatosAt/IT01-api/module/front-end/mapper"
	"github.com/tomatosAt/IT01-api/pkg/util"
)

func (s *Service) UserSVC(ctx context.Context, data dto.UserPayload) (*dto.DataInsertUser, int, error) {
	var result dto.DataInsertUser
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
