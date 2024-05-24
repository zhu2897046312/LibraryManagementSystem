package service

import (
	"LibraryManagementSystem/models"
	"fmt"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context){
	var user models.User
	//接收json数据
	if err := c.ShouldBindJSON(&user); err != nil{
		c.JSON(300, gin.H{
            "error": err.Error(),
        })
        return
	}
	fmt.Printf("user.UserName: %v\n", user.UserName)
	fmt.Printf("user.Password: %v\n", user.Password)
	if user.UserName == "" || user.Password == ""{
		c.JSON(300, gin.H{
            "error": "用户名或密码不能为空",
        })
        return
	}

	//查询数据库
	usr_after_find := models.User{
		UserName: user.UserName,
	}
	db := user.FindByName(&usr_after_find)
	if db.Error!= nil{
		c.JSON(300, gin.H{
            "error": db.Error.Error(),
        })
        return
	}

	//密码比对
	if usr_after_find.Password != user.Password{
        c.JSON(300, gin.H{
            "error": "密码错误",
        })
        return
    }
    c.JSON(200, gin.H{
        "message": "登录成功",
		"data": usr_after_find.IsAdministrator,	//是否是管理员
    })
}

func Register(c *gin.Context){
	user := models.User{}
	//接收json数据
	if err := c.ShouldBindJSON(&user); err != nil{
		c.JSON(300, gin.H{
            "error": err.Error(),
        })
        return
	}
	fmt.Printf("user.UserName: %v\n", user.UserName)
	fmt.Printf("user.Password: %v\n", user.Password)
	fmt.Printf("user.IsAdministrator: %v\n", user.IsAdministrator)
	//查询数据库
	db := user.FindByName(&user)

	//没找到
	if db.RowsAffected == 0{
		//插入数据库
		db = user.Insert(&user)
		if db.Error!= nil{
			c.JSON(300, gin.H{
				"error": db.Error.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "注册成功",
		})
		return
	}else{
		//找到了
		c.JSON(300, gin.H{
			"message": "用户已存在",
		})
		return
	}
}

func BorrowerPageQuery(c *gin.Context){
	//接收数据
	page := c.Query("page")
	pageSize := c.Query("page_size")

	page_, _ := strconv.Atoi(page)
	pageSize_, _ :=strconv.Atoi(pageSize)

	//TODO：token校验

	// 分页查询
	employee := models.Borrower{}
	data, db := employee.PageQuery(page_,pageSize_)
	if db.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "分页查找失败",
			"data":    nil,
		})
		return
	}
	//返回数据
	c.JSON(http.StatusOK, gin.H{
        "code":    http.StatusOK,
        "message": "分页查找成功",
        "data":    data,
    })
}


