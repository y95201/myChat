/*
 * @Description:
 * @Author: Y95201
 * @Date: 2022-12-19 10:09:37
 * @LastEditors: Y95201
 * @LastEditTime: 2022-12-20 10:49:21
 */
package main

import (
	"bytes"
	"log"
	"myChat/conf"
	"myChat/models"
	"myChat/routes"
	"myChat/views"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigType("json") // 设置配置文件的类型
	if err := viper.ReadConfig(bytes.NewBuffer(conf.AppJsonConfig)); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 未找到配置文件；如果需要，忽略错误
			log.Println("no such config file")
		} else {
			//已找到配置文件，但产生了另一个错误
			log.Println("read config error")
		}
		log.Fatal(err) // 读取配置文件失败致命错误
	}
	//
	models.InitDB()
}
func main() {
	//设置 release模式
	gin.SetMode(gin.ReleaseMode)
	//设置 debug模式
	//gin.SetMode(gin.DebugMode)
	//启动热更服务
	port := viper.GetString(`app.port`)
	//加载路由
	router := routes.InitRoute()
	//加载模板文件
	router.SetHTMLTemplate(views.GoTpl)
	//启动ws服务
	//go_ws.CleanOfflineConn()
	log.Println("监听端口", "http://127.0.0.1:"+port)
	http.ListenAndServe(":"+port, router)
}
