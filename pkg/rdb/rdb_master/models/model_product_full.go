package models

type ProductInfoFull struct {
	Product    `gorm:"embedded"`
	PosProduct `gorm:"embedded"`
}
