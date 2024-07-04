package models

import (
	"time"
	"LibraryManagementSystem/utils"
	"gorm.io/gorm"
)

// Borrowing represents a borrowing record in the library system
type Borrowing struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	BorrowerID   string    `gorm:"type:varchar(50);not null" json:"borrower_id"`
	BookID       string    `gorm:"type:varchar(50);not null" json:"book_id"`
	BorrowerName string    `gorm:"type:varchar(100);not null" json:"borrower_name"`
	BookName     string    `gorm:"type:varchar(100);not null" json:"book_name"`
	BorrowedDate time.Time `gorm:"type:datetime;not null" json:"borrowed_date"`
	FineAmount   float64   `gorm:"type:decimal(10,2);not null" json:"fine_amount"`
}

func init(){
	err := utils.DB_MySQL.AutoMigrate(&Borrowing{})
	if err!= nil {
        panic(err)
    }
}

func (u *Borrowing) Insert(employee *Borrowing) *gorm.DB {
	return utils.DB_MySQL.Model(&Borrowing{}).Create(employee)
}

func (u *Borrowing) FindByID(user *Borrowing) *gorm.DB {
	return utils.DB_MySQL.Model(&Borrowing{}).Where("borrower_id = ?", user.BorrowerID).Find(&user)
}

func (u *Borrowing) Update(employee *Borrowing) *gorm.DB {
	return utils.DB_MySQL.Model(&Borrowing{}).Where("borrower_id = ?", employee.BorrowerID).Updates(&employee)
}

func (u *Borrowing) Delete(employee *Borrowing) *gorm.DB {
	return utils.DB_MySQL.Model(&Borrowing{}).Where("borrower_id = ?", employee.BorrowerID).Delete(&Borrowing{})
}

func (u *Borrowing) GetAll() ([]Borrowing, *gorm.DB){
	var borrowers []Borrowing
    return borrowers, utils.DB_MySQL.Model(&Borrowing{}).Find(&borrowers)
}

func (u *Borrowing) PageQuery(page int, pageSize int) ([]Borrowing, *gorm.DB) {
	employees := make([]Borrowing, 0)
	var total int64

	utils.DB_MySQL.Model(&Borrowing{}).Count(&total)
	offset := (page - 1) * pageSize
	query := utils.DB_MySQL.Model(&Borrowing{}).Limit(pageSize).Offset(offset).Find(&employees)
	return employees, query
}