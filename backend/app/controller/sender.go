package controller

import (
	"Express/app/response"
	"Express/model"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// api/info
func BeginSend(c echo.Context) (err error) {
	data := new(model.InfoApiInput)
	if err = c.Bind(data); err != nil {
		return response.SendResponse(c, 404, "Bind failed")
	}
	exp := new(model.Senders)

	exp.Company = data.Company
	exp.Sender_Name = data.Sender_Name
	exp.Sender_Phone = data.Sender_Phone
	exp.Sender_Id = data.Sender_Id
	exp.Receiver_Name = data.Receiver_Name
	exp.Receiver_Phone = data.Receiver_Phone
	exp.Object_Type = data.Object_Type
	exp.Object_Weight = data.Object_Weight
	exp.Parcel_Insurance = data.Parcel_Insurance
	exp.Object_Value = data.Object_Value
	exp.Cash_On = data.Cash_On

	exp.Tracing_Num = "x"

	err = model.DB.Debug().Create(&exp).Error
	if err != nil {
		logrus.Fatal(err)
		return response.SendResponse(c, 404, "express create failed")
	}
	return response.SendResponse(c, 200, "succeed")

}

// api/getinfo
func GetInfo(c echo.Context) (err error) {
	exp := new(model.Senders)
	err = model.DB.Debug().Where("tracing_num = ?", "x").First(&exp).Error
	if err != nil {
		return response.SendResponse(c, 404, "found no express in repository")
	}
	return response.SendResponse(c, 200, "found", exp)
}

// api/updateinfo
func UpdateInfos(c echo.Context) (err error) {
	data := new(model.UpdateApiInput)
	if err = c.Bind(data); err != nil {
		return response.SendResponse(c, 404, "Bind failed")
	}
	exp := new(model.Senders)

	exp.Sender_Id = data.Sender_Id
	err = model.DB.Debug().Where("sender_id = ?", data.Sender_Id).First(&exp).Error
	if err != nil {
		logrus.Panic(err)
		return response.SendResponse(c, 404, "query failed")
	}
	exp.Tracing_Num = data.Tracing_Num
	err = model.DB.Debug().Where("sender_id = ?", data.Sender_Id).Updates(&exp).Error
	if err != nil {
		logrus.Fatal(err)
		return response.SendResponse(c, 404, "update err")
	}
	return response.SendResponse(c, 200, "tracing_num allocate successfully")

}
