package models

type Product struct {
}

func (Product) TableName() string {
	return TableNameMasterProduct
}
