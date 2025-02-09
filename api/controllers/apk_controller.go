package controllers

import (
	"github.com/dipeshdulal/clean-gin/domains"
	"github.com/dipeshdulal/clean-gin/dto"
	"github.com/dipeshdulal/clean-gin/lib"
	"github.com/dipeshdulal/clean-gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// ApkController data type
type ApkController struct {
	service domains.ApkService
	logger  lib.Logger
}

func NewApkController(apkService domains.ApkService, logger lib.Logger) ApkController {
	return ApkController{
		service: apkService,
		logger:  logger,
	}
}

func (u ApkController) CreateApk(c *gin.Context) {
	var req = dto.ApkDto{}
	if err := c.BindJSON(&req); err != nil {
		u.logger.Error("Tham số sai định dạng", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Tham số sai định dạng",
		})
		return
	}
	apk, err := u.service.CreateApk(req)
	if err != nil {
		u.logger.Error(err)
	}
	c.JSON(200, apk)
}

func (u ApkController) UpdateCategoryApk(c *gin.Context) {
	var req = dto.ApkDto{}
	if err := c.BindJSON(&req); err != nil {
		u.logger.Error("Tham số sai định dạng", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Tham số sai định dạng",
		})
		return
	}
	u.service.UpdateCategoryApk(req)
	c.JSON(200, gin.H{})
}

func (u ApkController) GetApk(c *gin.Context) {
	appId, _ := c.GetQuery("appId")
	apk, err := u.service.GetApk(appId)
	if err != nil {
		u.logger.Error(err)
	}
	c.JSON(200, apk)
}

func (u ApkController) GetFeaturedApk(c *gin.Context) {
	page := c.GetInt("page")
	size := c.GetInt("size")
	categorySlug, _ := c.GetQuery("slug")
	apk, err := u.service.FeaturedApk(dto.FilterApkDto{}, page, size, categorySlug)
	if err != nil {
		u.logger.Error(err)
	}
	c.JSON(200, apk)
}

func (u ApkController) GetFeaturedEducationApk(c *gin.Context) {
	page := c.GetInt("page")
	size := c.GetInt("size")
	categorySlug, _ := c.GetQuery("slug")
	apk, err := u.service.FeaturedEducationApk(dto.FilterApkDto{}, page, size, categorySlug)
	if err != nil {
		u.logger.Error(err)
	}
	c.JSON(200, apk)
}

func (u ApkController) SearchApk(c *gin.Context) {
	page := c.GetInt("page")
	size := c.GetInt("size")
	keyWord, _ := c.GetQuery("u")
	sortBy, _ := c.GetQuery("s")
	apk, err := u.service.SearchApk(keyWord, sortBy, page, size)
	if err != nil {
		u.logger.Error(err)
	}
	c.JSON(200, apk)
}

func (u ApkController) ClickDownload(c *gin.Context) {
	versionIdStr, _ := c.GetQuery("versionId")
	versionId, _ := strconv.ParseInt(versionIdStr, 10, 64)
	err := u.service.ClickDownload(versionId)
	if err != nil {
		u.logger.Error(err)
	}
	c.JSON(200, gin.H{})
}

func (u ApkController) GetSimilarApk(c *gin.Context) {
	page := c.GetInt("page")
	size := c.GetInt("size")

	req := dto.FilterApkDto{}
	if err := c.BindJSON(&req); err != nil {
		u.logger.Error("Tham số sai định dạng")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Tham số sai định dạng",
		})
		return
	}

	apk, err := u.service.SimilarApk(req, page, size)
	if err != nil {
		u.logger.Error(err)
	}
	c.JSON(200, apk)
}

func (u ApkController) GetSimilarDevelopApk(c *gin.Context) {
	page := c.GetInt("page")
	size := c.GetInt("size")

	req := dto.FilterApkDto{}
	if err := c.BindJSON(&req); err != nil {
		u.logger.Error("Tham số sai định dạng")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Tham số sai định dạng",
		})
		return
	}

	apk, err := u.service.SimilarDevelopApk(req, page, size)
	if err != nil {
		u.logger.Error(err)
	}
	c.JSON(200, apk)
}

func (u ApkController) CreateApkVersion(c *gin.Context) {

	req := []models.ApkVersion{}
	if err := c.BindJSON(&req); err != nil {
		u.logger.Error("Tham số sai định dạng")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Tham số sai định dạng",
		})
		return
	}

	err := u.service.CreateApkVersion(req)
	if err != nil {
		u.logger.Error(err)
		c.Error(err)
		return
	}
	c.JSON(200, nil)
}

func (u ApkController) GetApkVersion(c *gin.Context) {
	appId, _ := c.GetQuery("appId")

	rlt, err := u.service.GetApkVersion(appId)
	if err != nil {
		u.logger.Error(err)
		c.JSON(400, gin.H{})
		return
	}
	c.JSON(200, rlt)
}

func (u ApkController) GetApkInCategory(c *gin.Context) {
	page := c.GetInt("page")
	size := c.GetInt("size")
	categorySlug, _ := c.GetQuery("slug")
	sortBy, _ := c.GetQuery("sortBy")
	println(sortBy)

	rlt, err := u.service.GetApkInCategory(categorySlug, sortBy, page, size)
	if err != nil {
		u.logger.Error(err)
		c.JSON(400, gin.H{})
		return
	}
	c.JSON(200, rlt)
}

func (u ApkController) GetAllCategory(c *gin.Context) {
	page := c.GetInt("page")
	size := c.GetInt("size")
	categorySlug, _ := c.GetQuery("slug")
	sortBy, _ := c.GetQuery("sortBy")
	println(sortBy)

	rlt, err := u.service.GetAllCategory(categorySlug, sortBy, page, size)
	if err != nil {
		u.logger.Error(err)
		c.JSON(400, gin.H{})
		return
	}
	c.JSON(200, rlt)
}

func (u ApkController) GetReviews(c *gin.Context) {
	page := c.GetInt("page")
	size := c.GetInt("size")
	appId, _ := c.GetQuery("appId")
	apk, err := u.service.GetReviews(appId, page, size)
	if err != nil {
		u.logger.Error(err)
	}
	c.JSON(200, apk)
}

func (u ApkController) CreateReview(c *gin.Context) {
	req := make([]models.Review, 0)
	if err := c.BindJSON(&req); err != nil {
		u.logger.Error("Tham số sai định dạng")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Tham số sai định dạng",
		})
		return
	}
	err := u.service.CreateReviews(req)
	if err != nil {
		u.logger.Error(err)
	}
	c.JSON(200, nil)
}

func (com ApkController) GetFeed(c *gin.Context) {
	ads, _ := com.service.GetFeed()

	c.JSON(http.StatusOK, ads)
}

func (com ApkController) ClickFeed(c *gin.Context) {
	id, _ := c.GetQuery("id")
	i, err := strconv.ParseInt(id, 10, 64)
	err = com.service.ClickFeed(i)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (com ApkController) GetAds(c *gin.Context) {
	ads, _ := com.service.GetAds()

	c.JSON(http.StatusOK, ads)
}

func (com ApkController) ClickAds(c *gin.Context) {
	id, _ := c.GetQuery("id")
	i, err := strconv.ParseInt(id, 10, 64)
	err = com.service.ClickAds(i)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
