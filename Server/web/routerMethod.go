package web

import "Server/controler"
import (
	"github.com/gin-gonic/gin"
)
/*
	定义路由处理函数
*/
var handle *controler.Handle   //逻辑处理对象

func init()  {
	handle = controler.NewHandle()
}

func IndexRouter(c *gin.Context) {
	if c.Request.Form == nil {		//获取所有请求参数名和值
		c.Request.ParseMultipartForm(32 << 20)
	}
	handle.Insert_ser(c.Request.Form)

	//c.HTML(http.StatusOK, "index.html", nil)		//页面跳转
}
