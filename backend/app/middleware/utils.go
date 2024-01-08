package middleware

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	// "github.com/labstack/echo/v4"
)

func (c CustomValidator) Validate(i interface{}) error {
	c.ToInit()
	return c.validate.Struct(i)
}

func (c CustomValidator) ToInit() {
	c.once.Do(func() {
		c.validate = validator.New()
	})
}

func Cors(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")                                                            // 允许访问所有域，可以换成具体url，注意仅具体url才能带cookie信息
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token") //header的类型
		w.Header().Add("Access-Control-Allow-Credentials", "true")                                                    //设置为true，允许ajax异步请求带cookie信息
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")                             //允许请求方法
		w.Header().Set("content-type", "application/json;charset=UTF-8")                                              //返回数据格式是json
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		f(w, r)
	}
}

// func Corls(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		method := c.Request.Method
// 		origin := c.Request.Header.Get("Origin")
// 		c.Header("Access-Control-Allow-Origin", origin)
// 		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id")
// 		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
// 		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
// 		c.Header("Access-Control-Allow-Credentials", "true")
 
// 		// 放行所有OPTIONS方法
// 		if method == "OPTIONS" {
// 			c.AbortWithStatus(http.StatusNoContent)
// 		}
// 		// 处理请求
// 		return next(c)
// 	}
// }




