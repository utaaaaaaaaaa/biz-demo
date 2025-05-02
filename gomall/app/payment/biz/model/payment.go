package model

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type PaymentLog struct {
	gorm.Model
	UserId        int32     `json:"user_id"`
	OrderId       string    `json:"order_id"`
	TransactionId string    `json:"transaction_id"`
	PayAt         time.Time `json:"pay_at"`
	Amount        float32   `json:"amount"`
}

func (p PaymentLog) TableName() string {
	return "payment_log"
}

func CreatePaymentLog(db *gorm.DB, ctx context.Context, payment *PaymentLog) (err error) {
	return db.WithContext(ctx).Model(PaymentLog{}).Create(payment).Error
}
