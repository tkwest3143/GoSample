package data

import (
	"time"

	"gorm.io/gorm"
)

// WorkTimeSettingテーブルエンティティ情報
type WorkTimeSetting struct {
	gorm.Model
	ID        string
	UserId    string    //投稿者
	StartTime time.Time //投稿日
	EndTime   time.Time //投稿日
	RestTime  float32   //投稿日
}
