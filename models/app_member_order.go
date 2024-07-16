package models

type AppMemberOrder struct {
	log_idx  int    `gorm:"primaryKey;column:log_idx;autoIncrement:true"`
	Name     string `gorm:"column:beast_id"`
	Email    string
	Password []byte
}
