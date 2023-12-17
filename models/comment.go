package models

const TableNameComment = "comment"

// Comment mapped from table <comment>
type Comment struct {
	ID      int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Content string `gorm:"column:content" json:"content"`
	ApkID   string `gorm:"column:apk_id" json:"apk_id"`
}

// TableName Comment's table name
func (*Comment) TableName() string {
	return TableNameComment
}
