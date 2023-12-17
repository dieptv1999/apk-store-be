package repository

import (
	"github.com/dipeshdulal/clean-gin/lib"
	"gorm.io/gorm"
)

// ApkRepository database structure
type ApkRepository struct {
	lib.Database
	logger lib.Logger
}

// NewApkRepository creates a new user repository
func NewApkRepository(db lib.Database, logger lib.Logger) ApkRepository {
	return ApkRepository{
		Database: db,
		logger:   logger,
	}
}

// WithTrx enables repository with transaction
func (r ApkRepository) WithTrx(trxHandle *gorm.DB) ApkRepository {
	if trxHandle == nil {
		r.logger.Error("Transaction Database not found in gin context. ")
		return r
	}
	r.Database.DB = trxHandle
	return r
}
