package controller

// 入库
import (
	"Express/app/response"
	"Express/model"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"time"
)

// api/upload
func UploadExpress(c echo.Context) (err error) {
	data := new(model.UploadApiInput)
	if err = c.Bind(data); err != nil {
		logrus.Error(err)
		return response.SendResponse(c, 404, "Bind failed")
	}

	exp := new(model.Express_In)

	exp.Company = data.Company
	exp.Cash_On = data.Cash_On
	exp.Pre_Freight = data.Pre_Freight
	exp.Tracing_Num = data.Tracing_Num
	exp.Registred = data.Registred
	exp.Receiver_Phone = data.Receiver_Phone
	exp.Rejected = data.Rejected
	exp.Shelf_Num = "-1"
	exp.Deliver_time = time.Now()

	err = model.DB.Debug().Create(&exp).Error
	if err != nil {
		return response.SendResponse(c, 404, "express create error in database")

	}
	return response.SendResponse(c, 200, "succeed")
}

// api/allocate
func Allocate(c echo.Context) (err error) {
	// query for "-1"
	exp := new(model.Express_In)
	exp.Shelf_Num = "-1"
	err = model.DB.Where("shelf_num = ?", exp.Shelf_Num).First(&exp).Error
	if err != nil {
		return response.SendResponse(c, 404, "find error")
	}
	//fmt.Println(len(exp_s))
	return response.SendResponse(c, 200, "succeed", exp)
}

// api/allocated
func Allocated(c echo.Context) (err error) {
	data := new(model.AllocatedApiInput)
	if err = c.Bind(data); err != nil {
		logrus.Error(err)
		return response.SendResponse(c, 404, "Bind failed")
	}
	exp := new(model.Express_In)
	exp.Tracing_Num = data.Tracing_Num
	err = model.DB.Debug().Where("tracing_num = ?", exp.Tracing_Num).First(&exp).Error
	if err != nil {
		logrus.Error(err)
		return response.SendResponse(c, 404, "find nothing")
	}
	exp.Shelf_Num = data.Shelf_Num
	err = model.DB.Debug().Updates(&exp).Error
	if err != nil {
		return response.SendResponse(c, 404, "shelf_num updating error")
	}
	return response.SendResponse(c, 200, "succeed")

}
