package model

type Payment struct {
	ID            int64  `gorm:"primary_key;not_null;auto_increment"`
	PaymentName   string `json:"payment_name"` //支付
	PaymentSID    string `json:"payment_sid"`
	PaymentStatus bool   `json:"payment_status"` // 支付通道状态 true为生产
	PaymentImage  string `json:"payment_image"`  //支付图片logo
}
