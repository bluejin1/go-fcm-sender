package models

import (
	"gorm.io/datatypes"
)

type AppOrderProduct struct {
	No    string `json:"no"`
	Name  string `json:"name"`
	Cnt   string `json:"cnt"`
	Price string `json:"price"`
}

type AlarmOrderData struct {
	AppType  string         `json:"type"`
	Id       string         `json:"id"`
	Name     string         `json:"name"`
	Amount   string         `json:"amount"`
	Products datatypes.JSON `json:"products"`
}

type AlarmRecipeData struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Icon string `json:"icon"`
}

type OrderProduct struct {
	No    string `json:"no"`
	Name  string `json:"name"`
	Cnt   string `json:"cnt"`
	Price string `price:"price"`
}
