package services

import (
	"errors"
	"fmt"
	"github.com/dipeshdulal/clean-gin/models"
	"github.com/dipeshdulal/clean-gin/repository"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dipeshdulal/clean-gin/domains"
	"github.com/dipeshdulal/clean-gin/lib"
)

// JWTAuthService service relating to authorization
type JWTAuthService struct {
	env        lib.Env
	logger     lib.Logger
	repository repository.UserRepository
}

// NewJWTAuthService creates a new auth service
func NewJWTAuthService(env lib.Env, logger lib.Logger, repository repository.UserRepository) domains.AuthService {
	return JWTAuthService{
		env:        env,
		logger:     logger,
		repository: repository,
	}
}

// Authorize authorizes the generated token
func (s JWTAuthService) Authorize(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.env.JWTSecret), nil
	})
	if token.Valid {
		return true, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return false, errors.New("token malformed")
		}
		if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return false, errors.New("token expired")
		}
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return false, fmt.Errorf("validate: invalid token")
	}

	exp, err := strconv.ParseInt(claims["exp"].(string), 10, 64)
	if err != nil {
		return false, errors.New("invalid token")
	}

	if time.Unix(exp, 0).Before(time.Now()) {
		return false, errors.New("token expired")
	}

	return false, errors.New("couldn't handle token")
}

// CreateToken creates jwt auth token
func (s JWTAuthService) CreateToken(user models.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": map[string]any{
			"id":    user.ID,
			"name":  user.UserName,
			"email": user.Email,
		},
		"exp": time.Now().Add(time.Duration(s.env.TokenTTL) * time.Second).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.env.JWTSecret))

	if err != nil {
		s.logger.Error("JWT validation failed: ", err)
	}

	return tokenString
}

// CreateRefreshToken creates jwt auth token
func (s JWTAuthService) CreateRefreshToken(user models.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Duration(s.env.RefreshTokenTokenTTL) * time.Second).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.env.JWTRefreshSecret))

	if err != nil {
		s.logger.Error("JWT validation failed: ", err)
	}

	return tokenString
}

func (s JWTAuthService) RefreshAccessToken(refreshToken string) (string, error) {
	token, err := jwt.Parse(refreshToken, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.env.JWTRefreshSecret), nil
	})
	if token.Valid {
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return "", fmt.Errorf("validate: invalid token")
		}

		// * verify thời gian hết hạn
		exp := int(claims["exp"].(float64))

		if time.Unix(int64(exp), 0).Before(time.Now()) {
			return "", errors.New("token expired")
		}

		// * verify user existed
		userId := int(claims["id"].(float64))

		var user models.User
		err := s.repository.Find(&user, userId).Error
		if err != nil {
			return "", err
		}

		return s.CreateToken(user), nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return "", errors.New("token malformed")
		}
		if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return "", errors.New("token expired")
		}
	}

	return "", errors.New("couldn't handle token")
}
