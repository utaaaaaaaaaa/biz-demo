package model

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	userId    int32 `gorm:"type:int(11);not null;index:index_user_id"`
	productId int32 `gorm:"type:int(11);not null;"`
	Qty       int32 `gorm:"type:int(11);not null;"`
}

func (Cart) TableName() string {
	return "cart"
}
