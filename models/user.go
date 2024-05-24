package models

import (
	"LibraryManagementSystem/utils"

	"gorm.io/gorm"
)

type User struct {
	UserName        string `gorm:"type:varchar(50);unique;not null" json:"user_name"`
	Password        string `gorm:"type:varchar(100);not null" json:"password"`
	IsAdministrator bool   `gorm:"type:boolean;not null" json:"is_administrator"`
}

func init(){
	err := utils.DB_MySQL.AutoMigrate(&User{})
	if err!= nil {
        panic(err)
    }
}

func (u *User) Insert(employee *User) *gorm.DB {
	return utils.DB_MySQL.Model(&User{}).Create(employee)
}

func (u *User) Update(employee *User) *gorm.DB {
	return utils.DB_MySQL.Model(&User{}).Where("user_name =?", employee.UserName).Updates(&employee)
}

func (u *User) Delete(employee *User) *gorm.DB {
	return utils.DB_MySQL.Model(&User{}).Where("user_name =?", employee.UserName).Delete(&User{})
}

func (u *User) FindByName(user *User) *gorm.DB {
	return utils.DB_MySQL.Model(&User{}).Where("user_name =?", user.UserName).Find(&user)
}

func (u *User) GetAll() ([]User, *gorm.DB){
	var borrowers []User
    return borrowers, utils.DB_MySQL.Model(&User{}).Find(&borrowers)
}