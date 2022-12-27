package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"myChat/controller"
	"myChat/static"
	"net/http"
	//"myChat/ws/primary"
)

func InitRoute() *gin.Engine {
	//router := gin.Default()
	router := gin.New()

	if viper.GetString(`debug_mod`) == "false" {
		// 线上模式 打包用
		router.StaticFS("/static", http.FS(static.EmbedStatic))
	} else {
		// 开发模式 避免修改静态资源需要重启服务
		router.StaticFS("/static", http.Dir("static"))
	}

	sr := router.Group("/index")
	{
		sr.GET("/index", controller.Index)          //聊天页面
		sr.POST("/Login", controller.Login)         //聊天登录
		sr.POST("/quickadd", controller.Quickadd)   //快捷语录添加
		sr.POST("/quicklist", controller.Quicklist) //快捷语录列表
		sr.POST("/quickDell", controller.QuickDell) //快捷语录删除
		//sr.POST("/orderlist", controller.Orderlist) //订单列表
	}
	return router
}
