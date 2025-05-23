package shop

import (
	// "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Id              int     `gorm:"primary_key" json:"id"`
	Title           string  `json:"title"`
	Price           int     `json:"price"`
	DiscountPercent float32 `json:"discount_percent"`
	DiscountedPrice int     `json:"discount_price"`
	//FKs Brand
	BrandID int   `json:"brand_id"`
	Brand   Brand `grom:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:BrandID"`
	//Fks Category
	CategoryID int      `json:"category_id"`
	Category   Category `grom:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:CategoryID"`
}
type Brand struct {
	gorm.Model
	Id    int    `gorm:"primary_key" json:"id"`
	Title string `json:"title"`
}
type Category struct {
	gorm.Model
	Id     int    `gorm:"primary_key" json:"id"`
	Title  string `json:"title"`
	Parent int    `json:"parent"`
}
