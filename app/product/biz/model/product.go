package model

import (
	"context"
	"gorm.io/gorm"
)

type Product struct {
	Base
	Name        string     `gorm:"column:name" json:"name"`
	Description string     `gorm:"column:description" json:"description"`
	Picture     string     `gorm:"column:picture" json:"picture"`
	Price       float32    `gorm:"column:price" json:"price"`
	Categories  []Category `gorm:"many2many:product_category" json:"category"`
}

func (product *Product) TableName() string {
	return "product"
}

type ProductQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func (p ProductQuery) GetProductById(productId int) (product Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).First(&product, productId).Error
	return product, err
}

func (p ProductQuery) SearchProducts(query string) (products []*Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).Where("name like ? or description like ?", "%"+query+"%", "%"+query+"%").Find(&products).Error
	return products, err
}

func NewProductQuery(ctx context.Context, db *gorm.DB) *ProductQuery {
	return &ProductQuery{
		ctx: ctx,
		db:  db,
	}
}
