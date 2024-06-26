package routers

import (
	"github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "time"
	"LibraryManagementSystem/service"
)

func Router() *gin.Engine{
	r := gin.Default()

	//cors 配置			--- 跨域  不知道为什么要开vpn才能正确传递数据
	r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }))

	r.POST("/admin/login",service.Login)
	r.POST("/admin/register",service.Register)
    r.GET("/admin/BorrowerPageQuery",service.BorrowerPageQuery)
	return r
}