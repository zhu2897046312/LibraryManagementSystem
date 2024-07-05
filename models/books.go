package models

import (
	"time"
	"LibraryManagementSystem/utils"
	"gorm.io/gorm"
)

// Book represents a book in the library
type Book struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	BookName      string    `gorm:"type:varchar(100);not null" json:"book_name"`
	BookID        string    `gorm:"type:varchar(50);unique;not null" json:"book_id"`
	BookAuthor    string    `gorm:"type:varchar(100);not null" json:"book_author"`
	BookPublisher string    `gorm:"type:varchar(100);not null" json:"book_publisher"`
	BookType      string    `gorm:"type:varchar(50);not null" json:"book_type"`
	PublishTime   time.Time `gorm:"type:datetime;not null" json:"publish_time"`
	BookPrice     float64   `gorm:"type:decimal(10,2);not null" json:"book_price"`
	IsBorrowed    bool      `gorm:"type:boolean;not null" json:"is_borrowed"`
	InboundTime   time.Time `gorm:"type:datetime;not null" json:"inbound_time"`
}

// BookType represents the type of a book
type BookType struct {
	ID           uint    `gorm:"primaryKey" json:"id"`
	Type         string  `gorm:"type:varchar(50);unique;not null" json:"type"`
	BorrowedDays int64   `gorm:"type:int;not null" json:"borrowed_days"`
	FineAmount   float64 `gorm:"type:decimal(10,2);not null" json:"fine_amount"`
}

func init(){
	err := utils.DB_MySQL.AutoMigrate(&Book{})
	if err!= nil {
        panic(err)
    }
	err = utils.DB_MySQL.AutoMigrate(&BookType{})
	if err!= nil {
        panic(err)
    }
}

func (u *Book) Insert(employee *Book) *gorm.DB {
	return utils.DB_MySQL.Model(&Book{}).Create(employee)
}

func (u *Book) FindByID(user *Book) *gorm.DB {
	return utils.DB_MySQL.Model(&Book{}).Where("book_id = ?", user.BookID).Find(&user)
}

func (u *Book) Update(employee *Book) *gorm.DB {
	return utils.DB_MySQL.Model(&Book{}).Where("book_id =?", employee.BookID).Updates(&employee)
}

func (u *Book) Delete(employee *Book) *gorm.DB {
	return utils.DB_MySQL.Model(&Book{}).Where("book_id = ?", employee.BookID).Delete(&Book{})
}

func (u *Book) GetAll() ([]Book, *gorm.DB){
	var borrowers []Book
    return borrowers, utils.DB_MySQL.Model(&Book{}).Find(&borrowers)
}

func (u *BookType) Insert(employee *BookType) *gorm.DB {
	return utils.DB_MySQL.Model(&BookType{}).Create(employee)
}

func (u *BookType) Update(employee *BookType) *gorm.DB {
	return utils.DB_MySQL.Model(&BookType{}).Where("id = ?", employee.ID).Updates(&employee)
}

func (u *BookType) Delete(user_name string) *gorm.DB {
	return utils.DB_MySQL.Model(&BookType{}).Where("user_id = ?", user_name).Delete(&BookType{})
}

func (u *BookType) GetAll() ([]BookType, *gorm.DB){
	var borrowers []BookType
    return borrowers, utils.DB_MySQL.Model(&BookType{}).Find(&borrowers)
}

func (u *BookType) PageQuery(page int, pageSize int) ([]BookType, *gorm.DB) {
	employees := make([]BookType, 0)
	var total int64

	utils.DB_MySQL.Model(&BookType{}).Count(&total)
	offset := (page - 1) * pageSize
	query := utils.DB_MySQL.Model(&BookType{}).Limit(pageSize).Offset(offset).Find(&employees)
	return employees, query
}

func (u *Book) PageQuery(page int, pageSize int) ([]Book, *gorm.DB) {
	employees := make([]Book, 0)
	var total int64

	utils.DB_MySQL.Model(&Book{}).Count(&total)
	offset := (page - 1) * pageSize
	query := utils.DB_MySQL.Model(&Book{}).Limit(pageSize).Offset(offset).Find(&employees)
	return employees, query
}


