package models

type PosProduct struct {
}

func (PosProduct) TableName() string {
	return TableNameKiccProduct
}
