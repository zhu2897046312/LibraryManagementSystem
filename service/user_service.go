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

func BookInsert(c *gin.Context){
	user := models.Book{}
	//接收json数据
	if err := c.ShouldBindJSON(&user); err != nil{
		c.JSON(300, gin.H{
            "error": err.Error(),
        })
        return
	}
	//查询数据库
	db := user.FindByID(&user)

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
			"message": "成功",
		})
		return
	}else{
		//找到了
		c.JSON(300, gin.H{
			"message": "已存在",
		})
		return
	}
}

func BorrowingInsert(c *gin.Context){
	user := models.Borrowing{}
	//接收json数据
	if err := c.ShouldBindJSON(&user); err != nil{
		c.JSON(300, gin.H{
            "error": err.Error(),
        })
        return
	}
	//查询数据库
	db := user.FindByID(&user)

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
			"message": "成功",
		})
		return
	}else{
		//找到了
		c.JSON(300, gin.H{
			"message": "已存在",
		})
		return
	}
}

func BorrowerInsert(c *gin.Context){
	user := models.Borrower{}
	//接收json数据
	if err := c.ShouldBindJSON(&user); err != nil{
		c.JSON(300, gin.H{
            "error": err.Error(),
        })
        return
	}
	fmt.Printf("user: %v\n", user)
	//查询数据库
	db := user.FindByID(&user)

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
			"message": "成功",
		})
		return
	}else{
		//找到了
		c.JSON(300, gin.H{
			"message": "已存在",
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

func BorrowingPageQuery(c *gin.Context){
	//接收数据
	page := c.Query("page")
	pageSize := c.Query("page_size")

	page_, _ := strconv.Atoi(page)
	pageSize_, _ :=strconv.Atoi(pageSize)

	//TODO：token校验

	// 分页查询
	employee := models.Borrowing{}
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

func BooksPageQuery(c *gin.Context){
	//接收数据
	page := c.Query("page")
	pageSize := c.Query("page_size")

	page_, _ := strconv.Atoi(page)
	pageSize_, _ :=strconv.Atoi(pageSize)

	//TODO：token校验

	// 分页查询
	employee := models.Book{}
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
func UserPageQuery(c *gin.Context){
	//接收数据
	page := c.Query("page")
	pageSize := c.Query("page_size")

	page_, _ := strconv.Atoi(page)
	pageSize_, _ :=strconv.Atoi(pageSize)

	//TODO：token校验

	// 分页查询
	employee := models.User{}
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

func UserUpdate(c *gin.Context){
	var user_json models.User
	//接收json数据
	if err := c.ShouldBindJSON(&user_json); err != nil{
		c.JSON(300, gin.H{
            "error": err.Error(),
        })
        return
	}
    user := models.User{}
	user_update := user_json
    db := user.FindByName(&user_json)
    // 检查是否找到要update的用户
	if db.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "未找到要删除的用户",
			"data":    nil,
		})
		return
	}

    db = user.Update(&user_update)
    if db.Error!= nil{
        c.JSON(300, gin.H{
            "error": db.Error.Error(),
        })
        return
    }
	c.JSON(200, gin.H{
		"msg": "Updated",
	})
}

func UserDelete(c *gin.Context){
	var user_json models.User
	//接收json数据
	if err := c.ShouldBindJSON(&user_json); err != nil{
		c.JSON(300, gin.H{
            "error": err.Error(),
        })
        return
	}

	user := models.User{}
	db := user.FindByName(&user_json)
    if db.Error!= nil{
        c.JSON(http.StatusBadRequest, gin.H{
            "code":    http.StatusBadRequest,
            "message": "查询数据失败",
            "data":    nil,
        })
        return
    }
	// 检查是否找到要删除的用户
	if db.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "未找到要删除的用户",
			"data":    nil,
		})
		return
	}

	// 执行删除操作
	db = user.Delete(&user_json)
	if db.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": db.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "Deleted",
	})
}

func BookUpdate(c *gin.Context){
	var user_json models.Book
	//接收json数据
	if err := c.ShouldBindJSON(&user_json); err != nil{
		c.JSON(300, gin.H{
            "error": err.Error(),
        })
        return
	}
    user := models.Book{}
	user_update := user_json
    db := user.FindByID(&user_json)
    // 检查是否找到要update的用户
	if db.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "未找到要删除的用户",
			"data":    nil,
		})
		return
	}

    db = user.Update(&user_update)
    if db.Error!= nil{
        c.JSON(300, gin.H{
            "error": db.Error.Error(),
        })
        return
    }
	c.JSON(200, gin.H{
		"msg": "Updated",
	})
}

func BookDelete(c *gin.Context){
	var user_json models.Book
	//接收json数据
	if err := c.ShouldBindJSON(&user_json); err != nil{
		c.JSON(300, gin.H{
            "error": err.Error(),
        })
        return
	}

	user := models.Book{}
	db := user.FindByID(&user_json)
    if db.Error!= nil{
        c.JSON(http.StatusBadRequest, gin.H{
            "code":    http.StatusBadRequest,
            "message": "查询数据失败",
            "data":    nil,
        })
        return
    }
	// 检查是否找到要删除的用户
	if db.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "未找到要删除的用户",
			"data":    nil,
		})
		return
	}

	// 执行删除操作
	db = user.Delete(&user_json)
	if db.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": db.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "Deleted",
	})
}

func BorrowerUpdate(c *gin.Context){
	var user_json models.Borrower
	//接收json数据
	if err := c.ShouldBindJSON(&user_json); err != nil{
		c.JSON(300, gin.H{
            "error": err.Error(),
        })
        return
	}
    user := models.Borrower{}
	user_update := user_json
    db := user.FindByID(&user_json)
    // 检查是否找到要update的用户
	if db.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "未找到要删除的用户",
			"data":    nil,
		})
		return
	}

    db = user.Update(&user_update)
    if db.Error!= nil{
        c.JSON(300, gin.H{
            "error": db.Error.Error(),
        })
        return
    }
	c.JSON(200, gin.H{
		"msg": "Updated",
	})
}

func BorrowerDelete(c *gin.Context){
	var user_json models.Borrower
	//接收json数据
	if err := c.ShouldBindJSON(&user_json); err != nil{
		c.JSON(300, gin.H{
            "error": err.Error(),
        })
        return
	}

	user := models.Borrower{}
	db := user.FindByID(&user_json)
    if db.Error!= nil{
        c.JSON(http.StatusBadRequest, gin.H{
            "code":    http.StatusBadRequest,
            "message": "查询数据失败",
            "data":    nil,
        })
        return
    }
	// 检查是否找到要删除的用户
	if db.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "未找到要删除的用户",
			"data":    nil,
		})
		return
	}

	// 执行删除操作
	db = user.Delete(&user_json)
	if db.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": db.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "Deleted",
	})
}

func BorrowingUpdate(c *gin.Context){
	var user_json models.Borrowing
	//接收json数据
	if err := c.ShouldBindJSON(&user_json); err != nil{
		c.JSON(300, gin.H{
            "error": err.Error(),
        })
        return
	}
    user := models.Borrowing{}
	user_update := user_json
    db := user.FindByID(&user_json)
    // 检查是否找到要update的用户
	if db.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "未找到要删除的用户",
			"data":    nil,
		})
		return
	}

    db = user.Update(&user_update)
    if db.Error!= nil{
        c.JSON(300, gin.H{
            "error": db.Error.Error(),
        })
        return
    }
	c.JSON(200, gin.H{
		"msg": "Updated",
	})
}

func BorrowingDelete(c *gin.Context){
	var user_json models.Borrowing
	//接收json数据
	if err := c.ShouldBindJSON(&user_json); err != nil{
		c.JSON(300, gin.H{
            "error": err.Error(),
        })
        return
	}

	user := models.Borrowing{}
	db := user.FindByID(&user_json)
    if db.Error!= nil{
        c.JSON(http.StatusBadRequest, gin.H{
            "code":    http.StatusBadRequest,
            "message": "查询数据失败",
            "data":    nil,
        })
        return
    }
	// 检查是否找到要删除的用户
	if db.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "未找到要删除的用户",
			"data":    nil,
		})
		return
	}

	// 执行删除操作
	db = user.Delete(&user_json)
	if db.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": db.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "Deleted",
	})
}

