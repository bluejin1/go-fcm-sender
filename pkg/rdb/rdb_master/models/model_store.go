package models

type Store struct {
	StoreIdx         int64  `gorm:"primaryKey;column:store_idx;polymorphic:Owner;polymorphicValue:master" json:"store_idx"`
	StoreStatus      int8   `gorm:"column:store_status;" json:"store_status"`
	StoreName        string `gorm:"column:store_name" json:"store_name"`
	StoreGetAppOrder int8   `gorm:"column:store_get_app_order;" json:"store_get_app_order"`
	StoreAlarmRecipe int8   `gorm:"column:store_alarm_recipe" json:"store_alarm_recipe"`
	StoreAlarmOrder  int8   `gorm:"column:store_alarm_order" json:"store_alarm_order"`
}

func (Store) TableName() string {
	return TableNameMasterStore
}
