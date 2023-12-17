package services

import (
	"errors"
	"github.com/dipeshdulal/clean-gin/domains"
	"github.com/dipeshdulal/clean-gin/dto"
	"github.com/dipeshdulal/clean-gin/lib"
	"github.com/dipeshdulal/clean-gin/models"
	"github.com/dipeshdulal/clean-gin/repository"
	"gorm.io/gorm"
)

// UserService service layer
type UserService struct {
	logger       lib.Logger
	repository   repository.UserRepository
	snowflake    lib.Snowflake
	passwordHash lib.PasswordHash
	env          lib.Env
}

// NewUserService creates a new userservice
func NewUserService(env lib.Env, logger lib.Logger, repository repository.UserRepository) domains.UserService {
	return UserService{
		logger:     logger,
		repository: repository,
		env:        env,
	}
}

// WithTrx delegates transaction to repository database
func (s UserService) WithTrx(trxHandle *gorm.DB) domains.UserService {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}

// GetOneUser gets one user
func (s UserService) GetOneUser(id uint) (user models.User, err error) {
	return user, s.repository.Find(&user, id).Error
}

// GetAllUser get all the user
func (s UserService) GetAllUser() (users []models.User, err error) {
	return users, s.repository.Find(&users).Error
}

// CreateUser call to create the user
func (s UserService) CreateUser(registerDto dto.RegisterDto) error {
	var exists bool
	err := s.repository.Model(&models.User{}).Select("count(*) > 0").Where("UserName = ?", registerDto.Username).Find(&exists).Error
	if err != nil {
		s.logger.Error(err)
		return errors.New("Người dùng đã tồn tại trên hệ thống")
	}

	userId, err := s.snowflake.GenerateID()
	if err != nil {
		return err
	}
	passHashStr, salt, err := s.passwordHash.GeneratePassword(registerDto.Password, s.env.SaltSize, s.env.HashAlgorithm)
	if err != nil {
		return err
	}

	user := models.User{
		ID:          userId,
		UserName:    registerDto.Username,
		UserCode:    "UC" + s.passwordHash.RandStringBytes(6),
		PhoneNumber: registerDto.Username,
		Type:        "USER",
		Status:      1,
	}
	var hashing models.HashingAlgorithms
	if err := s.repository.Where("algorithm_name = ?", s.env.HashAlgorithm).First(&hashing).Error; err != nil {
		return err
	}
	user.LoginData = models.UserLoginData{
		ID:                 userId,
		PasswordHash:       passHashStr,
		PasswordSalt:       salt,
		HashAlgorithm:      hashing,
		HashingAlgorithmId: hashing.ID,
	}

	return s.repository.Create(&user).Error
}

// UpdateUser updates the user
func (s UserService) UpdateUser(user models.User) error {
	return s.repository.Save(&user).Error
}

// DeleteUser deletes the user
func (s UserService) DeleteUser(id uint) error {
	return s.repository.Delete(&models.User{}, id).Error
}

func (s UserService) GetOneByUserName(username string) (user models.User, err error) {
	return user, s.repository.Where("user_name = ?", username).First(&user).Error
}

func (s UserService) LoginByUserName(username string, password string) (user *models.User, err error) {
	err = s.repository.Preload("LoginData").Preload("LoginData.HashAlgorithm").Where("user_name = ?", username).First(&user).Error

	passHash := s.passwordHash.VerifyPassword(password, user.LoginData.PasswordHash, user.LoginData.PasswordSalt, user.LoginData.HashAlgorithm.AlgorithmName)

	if !passHash {
		return nil, errors.New("Sai mật khẩu")
	}

	return user, err
}
