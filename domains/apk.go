package domains

import (
	"github.com/dipeshdulal/clean-gin/dto"
	"github.com/dipeshdulal/clean-gin/models"
)

type ApkService interface {
	CreateApk(dto dto.ApkDto) (models.Apk, error)
	UpdateCategoryApk(dto dto.ApkDto)
	GetApk(appId string) (models.Apk, error)
	GetReviews(appId string, page int, size int) ([]models.Review, error)
	CreateReviews(dto []models.Review) error
	FilterApk(dto dto.FilterApkDto) ([]models.Apk, error)
	FeaturedApk(dto dto.FilterApkDto, page int, size int, categorySlug string) ([]models.Apk, error)
	FeaturedEducationApk(dto dto.FilterApkDto, page int, size int, categorySlug string) ([]models.Apk, error)
	SimilarApk(dto dto.FilterApkDto, page int, size int) ([]models.Apk, error)
	SimilarDevelopApk(dto dto.FilterApkDto, page int, size int) ([]models.Apk, error)
	CreateApkVersion(dto []models.ApkVersion) error
	GetApkVersion(appId string) ([]models.ApkVersion, error)
	GetApkInCategory(categorySlug string, sortBy string, page int, size int) ([]models.Apk, error)
	GetAllCategory(categorySlug string, sortBy string, page int, size int) ([]models.Category, error)
	SearchApk(keyWord string, sortBy string, page int, size int) ([]models.Apk, error)
	ClickDownload(apkVersionId int64) error

	GetAds() ([]models.Ads, error)
	ClickAds(id int64) error
	GetFeed() ([]models.Feed, error)
	ClickFeed(id int64) error
}
