package model

import "time"

type Express_In struct {
	Company     string `gorm:"type:char(50)" form:"company" json:"company"`                    // 物流公司
	Cash_On     bool   `gorm:"type:boolean" form:"cash_on" json:"cash_on"`                     // 快递是否到付
	Pre_Freight uint   `gorm:"type:uint" form:"pre_freight" json:"pre_freight"`     // 快递垫付运费
	Tracing_Num string `gorm:"type:char(21)" form:"tracing_num" json:"tracing_num"` // 快递单号
	Registred   bool   `gorm:"type:boolean" form:"registered" json:"registered"`
	// 快递是否挂号信
	Receiver_Phone string `gorm:"type:char(12)" form:"receiver_phone" json:"receiver_phone"`
	// 收件人手机号  11位手机号，1位'\0'
	Shelf_Num    string    `gorm:"type:char(5)" form:"shelf_num" json:"shelf_num"`   // 货架号
	Deliver_time time.Time `gorm:"datetime" form:"deliver_time" json:"deliver_time"` // 快递入库时间
	Rejected     bool      `gorm:"boolean" form:"rejected" json:"rejected"`          // 快递是否拒收
}

type Senders struct {
	Company          string  `gorm:"type:char(50)" form:"company" json:"company"`                    // 物流公司
	Sender_Name      string  `gorm:"type:char(21)" form:"sender_name" json:"sender_name"`            // 寄件人姓名
	Sender_Phone     string  `gorm:"type:char(12)" form:"sender_phone" json:"sender_phone"`          // 寄件人手机号
	Sender_Id        string  `gorm:"type:char(19)" form:"sender_id" json:"sender_id"`                // 寄件人身份证件号
	Receiver_Name    string  `gorm:"type:char(21)" form:"receiver_name" json:"receiver_name"`        // 收件人姓名
	Receiver_Phone   string  `gorm:"type:char(12)" form:"receiver_phone" json:"receiver_phone"`      // 收件人手机号
	Receiver_Addr    string  `gorm:"type:char(21)" form:"receiver_addr" json:"receiver_addr"`        // 收件人地址
	Object_Type      uint    `gorm:"type:uint" from:"object_type" json:"object_type"`                // 物品类型 [list indexes]
	Object_Weight    float32 `gorm:"type:float" form:"object_weights" json:"object_weights"`         // 物品重量
	Parcel_Insurance bool    `gorm:"type:boolean" form:"parcel_insurance" json:"parcel_insurance"`   // 是否保价
	Object_Value     float32 `gorm:"type:float" form:"object_value" json:"object_value"`             // 物品价值
	Cash_On          bool    `gorm:"type:boolean" form:"cash_on" json:"cash_on"`                     // 是否到付
	Tracing_Num      string  `gorm:"type:char(21)" form:"tracing_num" json:"tracing_num"` // 快递单号
}

// TODO:  add new model
type Url struct {
	Id int `gorm:"type:uint;primaryKey autoincrement" form:"id" json:"id"`
	//UserId     int       `gorm:"type:varchar(10) column:user_id" form:"user_id" json:"user_id"`
	Origin     string    `gorm:"type:varchar(200)" form:"origin" json:"origin"`
	Short      string    `gorm:"type:varchar(40)" form:"short" json:"short"` //?
	Comment    string    `gorm:"type:varchar(100)" form:"comment" json:"comment"`
	StartTime  time.Time `gorm:"type:datetime;autoCreateTime" json:"start-time"`
	ExpireTime time.Time `gorm:"type:datetime" json:"expire-time"` //2022-01-01T08:00:00+08:00
	Enable     string    `gorm:"type:varchar(40)" json:"enable"`
}

// var counturl int
type Users struct {
	Id    int    `gorm:"type:uint;primaryKey autoincrement=1000" form:"id" json:"id"`
	Name  string `gorm:"type:varchar(20)" form:"name" json:"name"`
	Email string `gorm:"type:varchar(50)" form:"email" json:"email"`
	Pwd   string `gorm:"type:varchar(90)" form:"pwd" json:"pwd"`
	SecQ  string `gorm:"type:varchar(100)" form:"secq" json:"secq"`
	SecA  string `gorm:"type:varchar(100)" form:"seca" json:"seca"`
	//IPpub      string    `gorm:"type:varchar(100)"form:"ippub" json:"ippub"`
	//IPpri      string    `gorm:"type:varchar(100)"form:"ippri" json:"ippri"`
	LatestTime time.Time `gorm:"type:datetime" form:"expire_time" json:"expire_time"`
	// url[]	?
	// Id: auto-increasement.
	// name,email,pwd,secq,seca mustn't contain special characters
}
