package models

type StoreInfoFull struct {
	Store     `gorm:"embedded"`
	StoreKicc `gorm:"embedded"`
}

type StoreInfoShort struct {
	StoreIdx    int32  `gorm:"primaryKey;column:store_idx;polymorphic:Owner;polymorphicValue:master" json:"store_idx"`
	StoreStatus int8   `gorm:"column:store_status;" json:"store_status"`
	StoreName   string `gorm:"column:store_name" json:"store_name"`
}
