package cache

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

//type gCache struct {
//	users gcache.Cache //该对象可以缓存任何类型数据
//}
//
//const (
//	cacheSize = 1_000_000
//	cacheTTL  = 1 * time.Hour // default expiration
//)
//设置缓存
func SaveAuthCache(info interface{}) {

	c := cache.New(5*time.Minute, 10*time.Minute)
	//设置缓存，使用默认过期时间

	c.Set("uid", info, cache.DefaultExpiration)
}

//获取缓存
func GetSessionUserInfo(info string) {
	c := cache.New(5*time.Minute, 10*time.Minute)
	var foo string
	if x, found := c.Get(info); found {
		foo = x.(string)
		fmt.Println(foo)
	}

}

// 清除缓存
func ClearAuthCache() {
	c := cache.New(5*time.Minute, 10*time.Minute)
	c.Delete("uid")
}

//func HasSession(c *gin.Context) bool {
//	session := sessions.Default(c)
//	if sessionValue := session.Get("uid"); sessionValue == nil {
//		return false
//	}
//	return true
//}
//
//func AuthSessionMiddle() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		session := sessions.Default(c)
//		sessionValue := session.Get("uid")
//		if sessionValue == nil {
//			c.Redirect(http.StatusFound, "/")
//			return
//		}
//
//		uidInt, _ := strconv.Atoi(sessionValue.(string))
//
//		if uidInt <= 0 {
//			c.Redirect(http.StatusFound, "/")
//			return
//		}
//
//		// 设置简单的变量
//		c.Set("uid", sessionValue)
//
//		c.Next()
//		return
//	}
//}
