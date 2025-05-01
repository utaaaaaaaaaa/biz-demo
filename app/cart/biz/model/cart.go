package model

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"log"
)

// 便于自动生成表
// err = DB.AutoMigrate(&model.Cart{})
type Cart struct {
	gorm.Model
	UserId    int32 `gorm:"type:int(11);not null;index:index_user_id"`
	ProductId int32 `gorm:"type:int(11);not null;"`
	Qty       int32 `gorm:"type:int(11);not null;"`
}

func (Cart) TableName() string {
	return "cart"
}

func AddItem(ctx context.Context, db *gorm.DB, item *Cart) error {
	var row Cart
	err := db.WithContext(ctx).Model(&Cart{}).Where(&Cart{UserId: item.UserId, ProductId: item.ProductId}).First(&row).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("查询cart时出错:", err)
		return err
	}
	if row.ID > 0 {
		return db.WithContext(ctx).Model(&Cart{}).Where(&Cart{UserId: item.UserId, ProductId: item.ProductId}).
			UpdateColumn("qty", gorm.Expr("qty+?", item.Qty)).Error
	}
	err = db.WithContext(ctx).Create(item).Error
	if err != nil {
		log.Println("创建cart时出错:", err)
	}
	return err
}

func EmptyCart(ctx context.Context, db *gorm.DB, userId int32) error {
	if userId == 0 {
		return errors.New("user id is required")
	}
	return db.WithContext(ctx).Delete(&Cart{}, "user_id = ?", userId).Error
}

func GetCart(ctx context.Context, db *gorm.DB, userId int32) ([]*Cart, error) {
	var carts []*Cart
	err := db.WithContext(ctx).Model(&Cart{}).Where(&Cart{UserId: userId}).Find(&carts).Error
	return carts, err
}
