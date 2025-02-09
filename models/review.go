// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

const TableNameReview = "review"

// Review mapped from table <review>
type Review struct {
	ID                   int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ReviewID             string    `gorm:"column:reviewId" json:"reviewId"`
	UserName             string    `gorm:"column:userName" json:"userName"`
	UserImage            string    `gorm:"column:userImage" json:"userImage"`
	Content              string    `gorm:"column:content" json:"content"`
	Score                int32     `gorm:"column:score" json:"score"`
	ThumbsUpCount        int32     `gorm:"column:thumbsUpCount" json:"thumbsUpCount"`
	ReviewCreatedVersion string    `gorm:"column:reviewCreatedVersion" json:"reviewCreatedVersion"`
	At                   string `gorm:"column:at" json:"at"`
	ReplyContent         string    `gorm:"column:replyContent" json:"replyContent"`
	RepliedAt            string `gorm:"column:repliedAt" json:"repliedAt"`
	AppVersion           string    `gorm:"column:appVersion" json:"appVersion"`
	AppId                string    `gorm:"column:appId" json:"appId"`
}

// TableName Review's table name
func (*Review) TableName() string {
	return TableNameReview
}
