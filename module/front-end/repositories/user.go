package repositories

import (
	"context"

	"github.com/tomatosAt/IT01-api/model"
	"gorm.io/gorm"
)

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
