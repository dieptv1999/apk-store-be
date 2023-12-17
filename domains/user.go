package domains

import (
	"github.com/dipeshdulal/clean-gin/dto"
	"github.com/dipeshdulal/clean-gin/models"
	"gorm.io/gorm"
)

type UserService interface {
	WithTrx(trxHandle *gorm.DB) UserService
	GetOneUser(id uint) (models.User, error)
	GetAllUser() ([]models.User, error)
	CreateUser(dto.RegisterDto) error
	UpdateUser(models.User) error
	DeleteUser(id uint) error
	GetOneByUserName(username string) (models.User, error)
	LoginByUserName(username string, password string) (*models.User, error)
}
