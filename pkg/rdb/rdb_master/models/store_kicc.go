package models

type StorePos struct {
	StorePosIdx int64 `gorm:"primaryKey;column:store_kicc_idx;polymorphic:Owner;polymorphicValue:master" json:"store_kicc_idx"`
}

func (StorePos) TableName() string {
	return TableNameMasterStoreKicc
}
