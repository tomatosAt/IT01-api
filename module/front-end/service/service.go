package services

import (
	"github.com/tomatosAt/IT01-api/module/front-end/ports"
	repositories "github.com/tomatosAt/IT01-api/module/front-end/repositories"
)

type Service struct {
	repo ports.Repository
}

func New(repo *repositories.Repository) *Service {
	return &Service{repo: repo}
}
