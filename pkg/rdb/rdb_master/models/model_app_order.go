package models

type AppOrder struct {
	OrderNo    string `gorm:"unique;column:orderNo;" json:"orderNo"`
	OrderTitle string `gorm:"column:orderTitle" json:"orderTitle"`
	RegNo      string `gorm:"column:regNo;" json:"regNo"`
	MemberNo   int32  `gorm:"column:memberNo;" json:"memberNo"`
	StoreIdx   int32  `gorm:"column:storeIdx" json:"storeIdx"`
}

func (AppOrder) TableName() string {
	return TableNameMasterAppOrder
}
