package routers

import (
	"github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "time"
	"LibraryManagementSystem/service"
)

func Router() *gin.Engine{
	r := gin.Default()

	//cors 配置			--- 跨域
	r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:8080", "http://192.168.171.243:8080"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }))

	r.POST("/admin/login",service.Login)
	r.POST("/admin/register",service.Register)
	// r.GET("/admin/employee/page",service.PageQuery)
	// r.POST("/admin/employee/update",service.UpdateEmployee)

	// r.GET("/admin/category/page",service.PageQueryCategory)
	// r.POST("/admin/category/addCategory",service.AddCategory)
	// r.POST("/admin/category/deleteCategory",service.DeleteCategory)
	// r.POST("/admin/category/updateCategory",service.UpdateCategory)
	return r
}