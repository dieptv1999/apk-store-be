package services

import (
	"github.com/dipeshdulal/clean-gin/domains"
	"github.com/dipeshdulal/clean-gin/dto"
	"github.com/dipeshdulal/clean-gin/lib"
	"github.com/dipeshdulal/clean-gin/models"
	"github.com/dipeshdulal/clean-gin/repository"
	"strconv"
	"strings"
	"time"
)

// ApkService service layer
type ApkService struct {
	logger     lib.Logger
	repository repository.ApkRepository
	snowflake  lib.Snowflake
	env        lib.Env
}

func (a ApkService) ClickDownload(apkVersionId int64) error {
	return a.repository.Exec("update apk_version set self_click_download = self_click_download + 1 where versionId = ?", apkVersionId).Error
}

func (a ApkService) SearchApk(keyWord string, sortBy string, page int, size int) ([]models.Apk, error) {
	orderBy := "realInstalls desc"
	switch sortBy {
	case "latest":
		orderBy = "updated desc"
	case "highestRated":
		orderBy = "score desc"
	}
	var apks []models.Apk

	err := a.repository.
		Model(&models.Apk{}).
		Order(orderBy).
		Offset(page*size).
		Limit(size).
		Find(&apks, "lower(title) like ?", "%"+strings.ToLower(keyWord)+"%").
		Error
	if err != nil {
		return make([]models.Apk, 0), err
	}

	a.repository.Save(&models.SearchHistory{
		SearchText: keyWord,
		CreatedAt:  time.Time{},
	})
	return apks, nil
}

func (a ApkService) FeaturedEducationApk(dto dto.FilterApkDto, page int, size int, categorySlug string) ([]models.Apk, error) {
	var ltApk []models.Apk
	var categories []models.Category
	err := a.repository.Find(&categories, "is_game = 0").Error
	if err != nil {
		return make([]models.Apk, 0), err
	}
	cateFilterStr := ""
	for i, c := range categories {
		if i == 0 {
			cateFilterStr = cateFilterStr + "categories like '%" + c.Name + "%' "
		} else {
			cateFilterStr = cateFilterStr + "or categories like '%" + c.Name + "%' "
		}
	}

	a.logger.Info(cateFilterStr)

	err = a.repository.Model(models.Apk{}).
		Offset(page*size).
		Limit(size).
		Order("realInstalls desc ").
		Find(&ltApk, "isHot = ? and ("+cateFilterStr+")", 1).Error

	if err != nil {
		return make([]models.Apk, 0), err
	}

	return ltApk, nil
}

func (a ApkService) GetAllCategory(categorySlug string, sortBy string, page int, size int) ([]models.Category, error) {
	var categories []models.Category

	err := a.repository.Find(&categories, "status = 1 and name like ?", "%"+categorySlug+"%").Error

	if err != nil {
		return make([]models.Category, 0), err
	}

	return categories, nil
}

func (a ApkService) UpdateCategoryApk(dto dto.ApkDto) {
	categoriesStr := ""
	if len(dto.Categories) > 0 {
		for _, v := range dto.Categories {
			categoriesStr = categoriesStr + "," + v.Name
		}
	}

	a.repository.Model(&models.Apk{}).Where("appId = ?", dto.AppID).Update("categories", categoriesStr)

	categories := make([]models.Category, 0)

	for _, c := range dto.Categories {
		categories = append(categories, models.Category{
			StoreID: c.ID,
			Name:    c.Name,
		})
	}

	a.repository.Save(&categories)
}

func (a ApkService) GetApkInCategory(categorySlug string, sortBy string, page int, size int) ([]models.Apk, error) {
	var apks []models.Apk
	orderBy := "realInstalls desc"
	switch sortBy {
	case "latest":
		orderBy = "updated desc"
	case "highestRated":
		orderBy = "score desc"
	}
	err := a.repository.Model(&models.Apk{}).
		Offset(page*size).
		Limit(size).
		Order(orderBy).
		Find(&apks, "lower(categories) like ?", "%"+strings.ToLower(categorySlug)+"%").Error

	if err != nil {
		return make([]models.Apk, 0), nil
	}
	return apks, nil
}

func (a ApkService) GetApkVersion(appId string) ([]models.ApkVersion, error) {
	var versions []models.ApkVersion
	err := a.repository.Find(&versions, "appId = ?", appId).Error
	if err != nil {
		return make([]models.ApkVersion, 0), nil
	}

	return versions, nil
}

func (a ApkService) CreateApkVersion(dto []models.ApkVersion) error {
	a.repository.Exec("update package set is_crawl_version = 1 where package_id = ?", dto[0].AppID)
	for _, v := range dto {
		v.SelfClickDownload = 0
	}
	return a.repository.Save(&dto).Error
}

func (a ApkService) SimilarDevelopApk(dto dto.FilterApkDto, page int, size int) ([]models.Apk, error) {
	var ltApk []models.Apk
	err := a.repository.Model(models.Apk{}).
		Offset(page*size).
		Limit(size).
		Order("realInstalls desc ").
		Find(&ltApk, "(developer = ? or developerId = ?) and appId <> ?", dto.Developer, dto.DeveloperId, dto.AppId).Error

	if err != nil {
		return make([]models.Apk, 0), err
	}

	return ltApk, nil
}

func (a ApkService) SimilarApk(dto dto.FilterApkDto, page int, size int) ([]models.Apk, error) {
	var ltApk []models.Apk
	err := a.repository.Model(models.Apk{}).
		Offset(page*size).
		Limit(size).
		Order("realInstalls desc ").
		Find(&ltApk, "categories like ? and appId <> ?", "%"+dto.GenreId+"%", dto.AppId).Error

	if err != nil {
		return make([]models.Apk, 0), err
	}

	return ltApk, nil
}

func (a ApkService) FeaturedApk(dto dto.FilterApkDto, page int, size int, categorySlug string) ([]models.Apk, error) {
	var ltApk []models.Apk
	err := a.repository.Model(models.Apk{}).
		Offset(page*size).
		Limit(size).
		Order("realInstalls desc ").
		Find(&ltApk, "isHot = ? and categories like ?", 1, "%"+categorySlug+"%").Error

	if err != nil {
		return make([]models.Apk, 0), err
	}

	return ltApk, nil
}

func (a ApkService) FilterApk(dto dto.FilterApkDto) ([]models.Apk, error) {
	//TODO implement me
	panic("implement me")
}

func (a ApkService) GetReviews(appId string, page int, size int) ([]models.Review, error) {
	var reviews []models.Review
	err := a.repository.
		Model(models.Review{}).
		Select("review.*").
		Offset(page*size).
		Limit(size).
		Find(&reviews, "appId = ?", appId).Error
	if err != nil {
		return make([]models.Review, 0), err
	}

	return reviews, nil
}

func (a ApkService) GetApk(appId string) (models.Apk, error) {
	var apk models.Apk
	err := a.repository.First(&apk, "appId = ?", appId).Error

	if err != nil {
		return models.Apk{}, err
	}

	return apk, nil
}

func (a ApkService) CreateReviews(dto []models.Review) error {
	err := a.repository.Save(&dto).Error
	return err
}

func (a ApkService) CreateApk(dto dto.ApkDto) (models.Apk, error) {
	histogramStr := ""
	categoriesStr := ""
	if len(dto.Histogram) > 0 {
		for _, v := range dto.Histogram {
			temp := strconv.Itoa(v)
			histogramStr = histogramStr + "," + temp
		}
	}

	if len(dto.Categories) > 0 {
		for _, v := range dto.Categories {
			temp := v.Name
			categoriesStr = categoriesStr + "," + temp
		}
	}

	apk := models.Apk{
		ID:                       dto.ID,
		Title:                    dto.Title,
		Description:              dto.Description,
		DescriptionHTML:          dto.DescriptionHTML,
		Summary:                  dto.Summary,
		Installs:                 dto.Installs,
		MinInstalls:              dto.MinInstalls,
		RealInstalls:             dto.RealInstalls,
		Score:                    dto.Score,
		Ratings:                  dto.Ratings,
		Reviews:                  dto.Reviews,
		Histogram:                histogramStr,
		Price:                    dto.Price,
		Free:                     dto.Free,
		Currency:                 dto.Currency,
		Sale:                     dto.Sale,
		SaleTime:                 dto.SaleTime,
		OriginalPrice:            dto.OriginalPrice,
		InAppProductPrice:        dto.InAppProductPrice,
		Developer:                dto.DeveloperID,
		DeveloperID:              dto.Developer,
		DeveloperEmail:           dto.DeveloperEmail,
		DeveloperWebsite:         dto.DeveloperWebsite,
		DeveloperAddress:         dto.DeveloperAddress,
		PrivacyPolicy:            dto.PrivacyPolicy,
		Genre:                    dto.GenreID,
		GenreID:                  dto.Genre,
		Icon:                     dto.Icon,
		HeaderImage:              dto.HeaderImage,
		Screenshots:              strings.Join(dto.Screenshots, ","),
		Video:                    dto.Video,
		VideoImage:               dto.VideoImage,
		ContentRating:            dto.ContentRating,
		ContentRatingDescription: dto.ContentRatingDescription,
		AdSupported:              dto.AdSupported,
		ContainsAds:              dto.ContainsAds,
		Released:                 dto.Released,
		Updated:                  time.Unix(dto.Updated, 0),
		Version:                  dto.Version,
		AppID:                    dto.AppID,
		URL:                      dto.URL,
		Categories:               categoriesStr,
	}

	err := a.repository.Save(&apk).Error
	if err != nil {
		return models.Apk{}, err
	}

	a.repository.Exec("update package set status = 1 where package_id = ?", dto.AppID)

	comments := make([]models.Comment, 0)
	for _, c := range dto.Comments {
		comments = append(comments, models.Comment{
			Content: c,
			ApkID:   dto.AppID,
		})
	}

	err = a.repository.Save(&comments).Error
	if err != nil {
		return models.Apk{}, err
	}

	categories := make([]models.Category, 0)

	for _, c := range dto.Categories {
		categories = append(categories, models.Category{
			StoreID: c.ID,
			Name:    c.Name,
		})
	}

	err = a.repository.Save(&categories).Error
	if err != nil {
		return models.Apk{}, err
	}

	return apk, nil
}

func (s ApkService) GetFeed() ([]models.Feed, error) {
	var ads []models.Feed
	err := s.repository.Find(&ads).Error
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	return ads, nil
}

func (s ApkService) ClickFeed(id int64) error {
	err := s.repository.Exec("update feed set click = click + 1 where id = ?", id).Error

	if err != nil {
		s.logger.Error(err)
		return err
	}

	return nil
}

func (s ApkService) ClickAds(id int64) error {
	err := s.repository.Exec("update ads set click = click + 1 where id = ?", id).Error

	if err != nil {
		s.logger.Error(err)
		return err
	}

	return nil
}

func (s ApkService) GetAds() ([]models.Ads, error) {
	var ads []models.Ads
	err := s.repository.Find(&ads).Error
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	return ads, nil
}

// NewApkService creates a new ApkService
func NewApkService(env lib.Env, logger lib.Logger, repository repository.ApkRepository) domains.ApkService {
	return ApkService{
		logger:     logger,
		repository: repository,
		env:        env,
	}
}
