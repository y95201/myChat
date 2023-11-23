<!--
 * @Description: 
 * @Author: Y95201
 * @Date: 2022-12-20 10:27:23
 * @LastEditors: y95201 957612196@qq.com
 * @LastEditTime: 2023-07-28 13:34:23
-->
聊天初始
|-- conf #配置文件
|   |-- config.go
|   `-- config.go.env
|-- controller
|   |-- ImageController.go
|   `-- IndexController.go
|-- main.go
|-- models
|   |-- message.go
|   |-- mysql.go
|   `-- user.go
|-- routes
|   `-- route.go
|-- services # 简单逻辑处理服务层
|   |-- helper
|   |   `-- helper.go
|   |-- img_kr
|   |   `-- imgKr.go
|   |-- message_service
|   |   `-- message.go
|   |-- session
|   |   `-- session.go
|   |-- user_service
|   |   `-- user.go
|   `-- validator
|       `-- validator.go
|-- sql
|   `-- go_gin_chat.sql
|-- static #静态文件 js 、css 、image 目录
|-- views
|   |-- index.html
|   |-- login.html
|   |-- private_chat.html
|   `-- room.html
`-- ws websocket 服务端主要逻辑
    |-- ServeInterface.go 
    |-- go_ws
    |   `-- serve.go # websocket服务端处理代码
    |-- primary
    |   `-- start.go # 为了兼容新旧版 websocket服务端 的调用策略
    |-- serve.go # 初版websocket服务端逻辑代码，可以忽略
    `-- ws_test #本地测试代码
        |-- exec.go
        `-- mock_ws_client_coon.go

框架说明引用
github.com/gin-gonic/gin
gorm.io/driver/mysql
gorm.io/gorm           
github.com/gravityblast/fresh    热更服务
github.com/valyala/fasthttp     http框架
github.com/spf13/viper      配置文件
github.com/gorilla/websocket    websocket框架
github.com/patrickmn/go-cache  cache

//日志
logfile, err := os.OpenFile("D:/phpstudy_pro/WWW/GoDome/dome/log/test.log", os.O_RDWR|os.O_CREATE, 0666)
if err != nil {
     fmt.Printf("%s\r\n", err.Error())
     os.Exit(-1)
}
defer logfile.Close()
logger := log.New(logfile, "\r\n", log.Ldate|log.Ltime|log.Llongfile)
logger.Println("hello")
logger.Fatal("test")

go mod init myChat
