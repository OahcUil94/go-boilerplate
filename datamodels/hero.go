package datamodels

import "time"

type Hero struct {
	ID 			int64 		`json:"id"        gorm:"type:bigint;pq_comment:英雄的编号"`
	Name 		string 		`json:"name"      gorm:"type:varchar(50);not null;pq_comment:英雄的名字"`
	RealName 	string 		`json:"realName"  gorm:"type:varchar(50);not null;pq_comment:英雄本人的名字"`
	CreatedAt 	*time.Time 	`json:"createdAt" gorm:"type:timestamptz;not null;default:now();pq_comment:英雄的创建时间"`
}

func (h Hero) TableName() string {
	return "t_heroes"
}
