package models

type PartnerMemberToken struct {
	AppToken        string `gorm:"column:app_token;" json:"app_token"`
	AppUuid         string `gorm:"column:app_uuid;" json:"app_uuid"`
	AppOs           string `gorm:"column:app_os;" json:"app_os"`
	StoreIdx        int32  `gorm:"column:store_idx" json:"store_idx"`
	AppTokenDate    string `gorm:"column:app_token_date" json:"app_token_date"`
	AllowBasicPush  int8   `gorm:"column:allow_basic_push;" json:"allow_basic_push"`
	AllowOrderPush  int8   `gorm:"column:allow_order_push;" json:"allow_order_push"`
	AllowRecipePush int8   `gorm:"column:allow_recipe_push;" json:"allow_recipe_push"`
	AllowEventPush  int8   `gorm:"column:allow_event_push;" json:"allow_event_push"`
	AllowQaPush     int8   `gorm:"column:allow_qa_push;" json:"allow_qa_push"`
}

func (PartnerMemberToken) TableName() string {
	return TableNamePartnerMemberToken
}
