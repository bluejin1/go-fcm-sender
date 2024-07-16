package types

type KiccOrderProduct struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	Count        string `json:"count"`
	UnitPrice    string `json:"unitPrice"`
	UnitOptPrice string `json:"unitOptPrice"`
	TotPrice     string `json:"totPrice"`
	TotDisPrice  string `json:"totDisPrice"`
	AmountPrice  string `json:"amountPrice"`
}

type KiccOrderInfo struct {
	OrderIdx          string             `json:"order_idx"`
	IsCheck           string             `json:"is_check"`
	OrderNumber       string             `json:"order_number"`
	TransactionIdx    string             `json:"transaction_idx"`
	OrderStatus       int32              `json:"order_status"`
	IsSet             int32              `json:"is_set"`
	ClientId          string             `json:"client_id"`
	DeviceId          string             `json:"device_id"`
	MemberPhone       string             `json:"member_phone"`
	OrderTime         string             `json:"order_time"`
	PayTime           string             `json:"pay_time"`
	SetStamps         string             `json:"set_stamps"`
	SetPayPoint       string             `json:"set_pay_point"`
	PaymentAmount     int32              `json:"payment_amount"`
	PaymentType       string             `json:"payment_type"`
	CreateIp          string             `json:"create_ip"`
	LastTransactionId int64              `json:"last_transaction_id"`
	CancelOrderNumber string             `json:"cancel_order_number"`
	CancelReason      string             `json:"cancel_reason"`
	CancelDate        string             `json:"cancel_date"`
	CreateTime        string             `json:"create_time"`
	CreateDate        string             `json:"create_date"`
	ModifyDate        string             `json:"modify_date"`
	OrderProducts     []KiccOrderProduct `json:"order_products"`
}

type StoreChangeConfig struct {
	StoreIdx     int32  `json:"storeIdx"`
	PushCategory string `json:"pushCategory"`
	PushTitle    string `json:"pushTitle"`
	PushBody     string `json:"pushBody"`
	ConfigKey    string `json:"configKey"`
	ConfigVal    string `json:"configVal"`
	ModifyDate   string `json:"modifyDate"`
}

type PartnerFcmPushData struct {
	FcmType          string `json:"type" structs:"type"`
	Category         string `json:"category" structs:"category"`
	Title            string `json:"title" structs:"title"`
	Body             string `json:"body" structs:"body"`
	StoreNo          string `json:"store_no" structs:"store_no"`
	OrderNumber      string `json:"order_number" structs:"order_number"`
	OrderStatus      string `json:"order_status" structs:"order_status"`
	OrderReceiveType string `json:"order_receiveType" structs:"order_receiveType"`
	ChangeKey        string `json:"change_key" structs:"change_key"`
	ChangeVal        string `json:"change_val" structs:"change_val"`
	AddBody          string `json:"add_body" structs:"add_body"`
	IconType         string `json:"icon_type" structs:"icon_type"`
	UserIdx          string `json:"userIdx" structs:"userIdx"`
	Time             string `json:"time" structs:"time"`
}
