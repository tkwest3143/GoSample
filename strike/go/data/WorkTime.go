package data

import (
	"time"

	"gorm.io/gorm"
)

// WorkTimeテーブルエンティティ情報
type WorkTime struct {
	gorm.Model
	StartTime time.Time //投稿日
	EndTime   time.Time //投稿日
	RestTime  float32   //投稿日
	UserID    string
}
