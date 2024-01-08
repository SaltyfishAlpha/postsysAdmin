package model

type UploadApiInput struct {
	Company     string `gorm:"type:char(50)" form:"company" json:"company"`                    // 物流公司
	Cash_On     bool   `gorm:"type:boolean" form:"cash_on" json:"cash_on"`                     // 快递是否到付
	Pre_Freight uint   `gorm:"type:uint" form:"pre_freight" json:"pre_freight"`     // 快递垫付运费
	Tracing_Num string `gorm:"type:char(21)" form:"tracing_num" json:"tracing_num"` // 快递单号
	Registred   bool   `gorm:"type:boolean" form:"registered" json:"registered"`
	// 快递是否挂号信
	Receiver_Phone string `gorm:"type:char(12)" form:"receiver_phone" json:"receiver_phone"`
	// 收件人手机号  11位手机号，1位'\0'
	//Shelf_Num    string    `gorm:"type:char(5)" form:"shelf_num" json:"shelf_num"`   // 货架号
	//Deliver_time time.Time `gorm:"datetime" form:"deliver_time" json:"deliver_time"` // 快递入库时间
	// 这两项需要生成
	Rejected bool `gorm:"boolean" form:"rejected" json:"rejected"` // 快递是否拒收

}

type Exp_List_Return struct {
	Exp_List []Express_In
}

// AllocateInput: null

type AllocatedApiInput struct {
	Tracing_Num string `gorm:"type:char(21)" form:"tracing_num" json:"tracing_num"` // 快递单号
	Shelf_Num   string `gorm:"type:char(5)" form:"shelf_num" json:"shelf_num"`                 // 货架号
}

type QueryApiInput struct {
	Receiver_Phone string `gorm:"type:char(12)" form:"receiver_phone" json:"receiver_phone"`
}

type RemoveApiInput struct {
	Tracing_Num    string `gorm:"type:char(21)" form:"tracing_num" json:"tracing_num"` // 快递单号
	Receiver_Phone string `gorm:"type:char(12)" form:"receiver_phone" json:"receiver_phone"`
}

// 取出

type InfoApiInput struct {
	Company          string  `gorm:"type:varchar(50)" form:"company" json:"company"`               // 物流公司
	Sender_Name      string  `gorm:"type:char(21)" form:"sender_name" json:"sender_name"`          // 寄件人姓名
	Sender_Phone     string  `gorm:"type:char(12)" form:"sender_phone" json:"sender_phone"`        // 寄件人手机号
	Sender_Id        string  `gorm:"type:char(19)" form:"sender_id" json:"sender_id"`              // 寄件人身份证件号
	Receiver_Name    string  `gorm:"type:char(21)" form:"receiver_name" json:"receiver_name"`      // 收件人姓名
	Receiver_Phone   string  `gorm:"type:char(12)" form:"receiver_phone" json:"receiver_phone"`    // 收件人手机号
	Receiver_Addr    string  `gorm:"type:char(21)" form:"receiver_addr" json:"receiver_addr"`      // 收件人地址
	Object_Type      uint    `gorm:"type:uint" from:"object_type" json:"object_type"`              // 物品类型 [list indexes]
	Object_Weight    float32 `gorm:"type:float" form:"object_weights" json:"object_weights"`       // 物品重量
	Parcel_Insurance bool    `gorm:"type:boolean" form:"parcel_insurance" json:"parcel_insurance"` // 是否保价
	Object_Value     float32 `gorm:"type:float" form:"object_value" json:"object_value"`           // 物品价值
	Cash_On          bool    `gorm:"type:boolean" form:"cash_on" json:"cash_on"`                   // 是否到付
	//Tracing_Num      string  `gorm:"type:char(21);primaryKey" form:"tracing_num" json:"tracing_num"` // 快递单号
}

type Info_List_Return struct {
	Info_List []Senders
}

// 寄件

//type GetinfoApiInput: null

type UpdateApiInput struct {
	Sender_Id   string `gorm:"type:char(19)" form:"sender_id" json:"sender_id"`                // 寄件人身份证件号
	Tracing_Num string `gorm:"type:char(21)" form:"tracing_num" json:"tracing_num"` // 快递单号
}
