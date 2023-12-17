package controllers

import (
	"github.com/dipeshdulal/clean-gin/constants"
	"github.com/dipeshdulal/clean-gin/domains"
	"github.com/dipeshdulal/clean-gin/dto"
	"github.com/dipeshdulal/clean-gin/lib"
	"github.com/dipeshdulal/clean-gin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// JWTAuthController struct
type JWTAuthController struct {
	logger      lib.Logger
	service     domains.AuthService
	userService domains.UserService
}

// NewJWTAuthController creates new controller
func NewJWTAuthController(
	logger lib.Logger,
	service domains.AuthService,
	userService domains.UserService,
) JWTAuthController {
	return JWTAuthController{
		logger:      logger,
		service:     service,
		userService: userService,
	}
}

// SignIn signs in user
func (jwt JWTAuthController) SignIn(c *gin.Context) {
	jwt.logger.Info("SignIn route called")
	var req = dto.LoginDto{}
	if err := c.BindJSON(&req); err != nil {
		jwt.logger.Error("Tham số sai định dạng")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Tham số sai định dạng",
		})
		return
	}
	// Currently not checking for username and password
	// Can add the logic later if necessary.
	var user *models.User
	user, err := jwt.userService.LoginByUserName(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	token := jwt.service.CreateToken(*user)
	refreshToken := jwt.service.CreateRefreshToken(*user)
	c.JSON(200, gin.H{
		"message":       "logged in successfully",
		"access_token":  token,
		"refresh_token": refreshToken,
	})
}

// Register registers user
func (jwt JWTAuthController) Register(c *gin.Context) {
	user := dto.RegisterDto{}
	trxHandle := c.MustGet(constants.DBTransaction).(*gorm.DB)

	if err := c.ShouldBindJSON(&user); err != nil {
		jwt.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := jwt.userService.WithTrx(trxHandle).CreateUser(user); err != nil {
		jwt.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": "user created"})
}

func (jwt *JWTAuthController) RefreshAccessToken(ctx *gin.Context) {

	refreshToken, err := ctx.GetQuery("refresh_token")
	if !err {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Refresh Token không hợp lệ"})
	}

	accessToken, er := jwt.service.RefreshAccessToken(refreshToken)
	if er != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": er.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}
