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

func (u *Borrowing) Insert(employee *Borrowing) *gorm.DB {
	return utils.DB_MySQL.Model(&Borrowing{}).Create(employee)
}

func (u *Borrowing) Update(employee *Borrowing) *gorm.DB {
	return utils.DB_MySQL.Model(&Borrowing{}).Where("user_id =?", employee.ID).Updates(&employee)
}

func (u *Borrowing) Delete(user_name string) *gorm.DB {
	return utils.DB_MySQL.Model(&Borrowing{}).Where("user_id =?", user_name).Delete(&Borrowing{})
}