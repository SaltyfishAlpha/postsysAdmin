package controller

import (
	"Express/app/response"
	"Express/model"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// 寄件

// api/query
func QueryMine (c echo.Context) (err error) {
	data := new(model.QueryApiInput)
	if err = c.Bind(data); err != nil {
		return response.SendResponse(c, 404, "Bind failed",)
	}
	exp := new(model.Express_In)
	// exp.Receiver_Phone = data.Receiver_Phone
	err = model.DB.Debug().Where("receiver_phone = ?", data.Receiver_Phone).First(&exp).Error
	if err != nil {
		return response.SendResponse(c, 200, "No express found for you")
	}
	return response.SendResponse(c, 200, "find~", exp)
}



// api/remove
func PickUp (c echo.Context) (err error) {
	data := new(model.RemoveApiInput)
	if err = c.Bind(data); err != nil {
		logrus.Error(err)
		return response.SendResponse(c, 404, "Bind failed",)
	}
	exp := new(model.Express_In)
	exp.Tracing_Num = data.Tracing_Num
	//exp.Receiver_Phone = data.Receiver_Phone
	err = model.DB.Debug().Where("tracing_Num = ?", exp.Tracing_Num).First(&exp).Error
	if err != nil {
		return response.SendResponse(c, 404, "No Record Found",)
	}
	if exp.Receiver_Phone != data.Receiver_Phone {
		return response.SendResponse(c, 404, "Express not yours",)
	}
	err = model.DB.Debug().Where("tracing_num = ?", exp.Tracing_Num).Delete(&exp).Error
	if err != nil {
		logrus.Fatal(err)
		return response.SendResponse(c, 404, "delete (pick up) error")
	}
	return response.SendResponse(c, 200, "succeed",)
}







