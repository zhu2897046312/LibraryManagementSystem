package service

import (
    "LibraryManagementSystem/models"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context){
	user := models.User{}
	//接收json数据
	if err := c.ShouldBindJSON(&user); err != nil{
		c.JSON(300, gin.H{
            "error": err.Error(),
        })
        return
	}
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
		c.JSON(200, gin.H{
            "error": err.Error(),
        })
        return
	}

	//查询数据库
	db := user.FindByName(&user)

	//没找到
	if db.RowsAffected == 0{
		//插入数据库
		db = user.Insert(&user)
		if db.Error!= nil{
			c.JSON(200, gin.H{
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
		c.JSON(200, gin.H{
			"message": "用户已存在",
		})
		return
	}
}