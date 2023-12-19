package models

import "time"

const TableNameApkVersion = "apk_version"

// ApkVersion mapped from table <apk_version>
type ApkVersion struct {
	VersionID       int64     `gorm:"column:versionId;primaryKey;autoIncrement:true" json:"versionId"`
	AppID           string    `gorm:"column:appId;not null" json:"appId"`
	RequiresAndroid string    `gorm:"column:requiresAndroid" json:"requiresAndroid"`
	Architecture    string    `gorm:"column:architecture" json:"architecture"`
	Signature       string    `gorm:"column:signature" json:"signature"`
	Permissions     string    `gorm:"column:permissions" json:"permissions"`
	Version         string    `gorm:"column:version" json:"version"`
	VersionCode     string    `gorm:"column:versionCode" json:"versionCode"`
	DownloadLink    string    `gorm:"column:downloadLink" json:"downloadLink"`
	ApkPureLink     string    `gorm:"column:apkPureLink" json:"apkPureLink"`
	AmountStorage   string    `gorm:"column:amountStorage" json:"amountStorage"`
	ApkType         string    `gorm:"column:apkType" json:"apkType"`
	UpdateOn        string    `gorm:"column:updateOn" json:"updateOn"`
	CreateAt        time.Time `gorm:"column:createAt" json:"createAt"`
}

// TableName ApkVersion's table name
func (*ApkVersion) TableName() string {
	return TableNameApkVersion
}
