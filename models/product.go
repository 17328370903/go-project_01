package models

import "gorm.io/gorm"

type Product struct {
	Name       string  `json:"name" gorm:"comment:商品名称;default:''" `
	Price      float64 `json:"price" comment:"商品价格"`
	CategoryId uint    `json:"category" comment:"商品分类"`
	BrandId    uint    `json:"brand" comment:"商品品牌"`
	Status     uint8   `json:"status" comment:"商品状态"`
	gorm.Model
}
