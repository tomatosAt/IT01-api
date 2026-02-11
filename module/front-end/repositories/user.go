package repositories

import (
	"context"
	"errors"

	"github.com/tomatosAt/IT01-api/model"
	"gorm.io/gorm"
)

func (r *Repository) GetUserByFullNameAndDobRepo(ctx context.Context, tx *gorm.DB, fNameTh, lNameTh, dob string) error {
	// * tx
	if tx == nil {
		tx = r.DB().Ctx()
	}
	var count int64
	db := r.DB().Ctx().WithContext(ctx).Model(&model.User{})
	err := db.Where("first_name_th = ? AND last_name_th = ? AND birth_date = ?", fNameTh, lNameTh, dob).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("user already exists")
	}
	return nil
}

func (r *Repository) InsertUserRepo(ctx context.Context, tx *gorm.DB, data model.User) (model.User, error) {
	// * tx
	if tx == nil {
		tx = r.DB().Ctx()
	}
	if err := tx.WithContext(ctx).Create(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}
