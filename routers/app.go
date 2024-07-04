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
    //insert
	r.POST("/admin/register",service.Register)
    r.POST("/admin/BookInsert",service.BookInsert)
    r.POST("/admin/BorrowerInsert",service.BorrowerInsert)
    r.POST("/admin/BorrowingInsert",service.BorrowingInsert)
    //update
    r.POST("/admin/UserUpdate",service.UserUpdate)
    r.POST("/admin/BookUpdate",service.BookUpdate)
    r.POST("/admin/BorrowerUpdate",service.BorrowerUpdate)
    r.POST("/admin/BorrowingUpdate",service.BorrowingUpdate)
    //delete
    r.POST("/admin/UserDelete",service.UserDelete)
    r.POST("/admin/BookDelete",service.BookDelete)
    r.POST("/admin/BorrowerDelete",service.BorrowerDelete)
    r.POST("/admin/BorrowingDelete",service.BorrowingDelete)
    //query
    r.GET("/admin/BorrowerPageQuery",service.BorrowerPageQuery)
    r.GET("/admin/BookPageQuery",service.BooksPageQuery)
    r.GET("/admin/BorrowingPageQuery",service.BorrowingPageQuery)
    r.GET("/admin/UserPageQuery",service.UserPageQuery)
	return r
}