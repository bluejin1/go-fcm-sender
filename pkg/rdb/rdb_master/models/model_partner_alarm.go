package models

type PartnerAlarm struct {
}

func (PartnerAlarm) TableName() string {
	return TableNameMasterPartnerAlarm
}
