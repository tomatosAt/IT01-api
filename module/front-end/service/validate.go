package services

import (
	"context"
	"errors"
	"time"

	"github.com/tomatosAt/IT01-api/module/front-end/dto"
)

func (s *Service) CheckFormatPreRegisterSVC(ctx context.Context, data *dto.UserPayload) error {
	if _, err := s.DOBCheckSVC(ctx, data.Dob); err != nil {
		return err
	}
	return nil
}

func (s *Service) DOBCheckSVC(ctx context.Context, dob string) (string, error) {
	_, errDob := time.Parse("02-01-2006", dob)
	if errDob != nil {
		return "", errors.New("birth_date must be YYYY-MM-DD format")
	}
	return dob, nil
}
