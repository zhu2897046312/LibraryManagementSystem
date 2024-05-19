package models

import (
	"time"
	//"gorm.io/driver/sqlite"
	"LibraryManagementSystem/utils"
	"gorm.io/gorm"
)

// Borrower represents a borrower in the library system
type Borrower struct {
	ID                   uint      `gorm:"primaryKey" json:"id"`
	BorrowerID           string    `gorm:"type:varchar(50);unique;not null" json:"borrower_id"`
	BorrowerName         string    `gorm:"type:varchar(100);not null" json:"borrower_name"`
	BorrowerAddress      string    `gorm:"type:varchar(255)" json:"borrower_address"`
	BorrowerPhone        string    `gorm:"type:varchar(20)" json:"borrower_phone"`
	BorrowerType         string    `gorm:"type:varchar(50)" json:"borrower_type"`
	BorrowerBooksNumbers int       `gorm:"type:int" json:"borrower_books_numbers"`
	IdentityCardID       string    `gorm:"type:varchar(50)" json:"identity_card_id"`
	IssuanceDate         time.Time `gorm:"type:datetime" json:"issuance_date"`
	IsLoss               bool      `gorm:"type:boolean" json:"is_loss"`
	Gender               int       `gorm:"type:int" json:"gender"`
	Unit                 string    `gorm:"type:varchar(100)" json:"unit"`
}

func (u *Borrower) Insert(employee *Borrower) *gorm.DB {
	return utils.DB_MySQL.Model(&Borrower{}).Create(employee)
}

func (u *Borrower) Update(employee *Borrower) *gorm.DB {
	return utils.DB_MySQL.Model(&Borrower{}).Where("user_id =?", employee.BorrowerID).Updates(&employee)
}

func (u *Borrower) Delete(user_name string) *gorm.DB {
	return utils.DB_MySQL.Model(&Borrower{}).Where("user_id =?", user_name).Delete(&Borrower{})
}

