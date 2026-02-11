package ports

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/tomatosAt/IT01-api/config"
	"github.com/tomatosAt/IT01-api/model"
	"github.com/tomatosAt/IT01-api/module/front-end/dto"
	"github.com/tomatosAt/IT01-api/pkg/database"
	"gorm.io/gorm"
)

type Repository interface {
	Module() string
	AppCfg() *config.Config
	Log() *logrus.Entry
	DB() *database.Client
	InsertUserRepo(ctx context.Context, tx *gorm.DB, data model.User) (model.User, error)
	GetUserByFullNameAndDobRepo(ctx context.Context, tx *gorm.DB, fNameTh, lNameTh, dob string) error
}

type Service interface {
	CheckFormatPreRegisterSVC(ctx context.Context, data *dto.UserPayload) error
	UserSVC(ctx context.Context, data dto.UserPayload) (*dto.DataInsertUser, int, error)
}
