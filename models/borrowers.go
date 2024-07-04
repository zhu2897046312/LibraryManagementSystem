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
func init(){
	err := utils.DB_MySQL.AutoMigrate(&Borrower{})
	if err!= nil {
        panic(err)
    }
}

func (u *Borrower) Insert(employee *Borrower) *gorm.DB {
	return utils.DB_MySQL.Model(&Borrower{}).Create(employee)
}

func (u *Borrower) FindByID(user *Borrower) *gorm.DB {
	return utils.DB_MySQL.Model(&Borrower{}).Where("borrower_id = ?", user.BorrowerID).Find(&user)
}

func (u *Borrower) Update(employee *Borrower) *gorm.DB {
	return utils.DB_MySQL.Model(&Borrower{}).Where("borrower_id = ?", employee.BorrowerID).Updates(&employee)
}

func (u *Borrower) Delete(employee *Borrower) *gorm.DB {
	return utils.DB_MySQL.Model(&Borrower{}).Where("borrower_id = ?", employee.BorrowerID).Delete(&Borrower{})
}

func (u *Borrower) GetAll() ([]Borrower, *gorm.DB){
	var borrowers []Borrower
    return borrowers, utils.DB_MySQL.Model(&Borrower{}).Find(&borrowers)
}

func (u *Borrower) PageQuery(page int, pageSize int) ([]Borrower, *gorm.DB) {
	employees := make([]Borrower, 0)
	var total int64

	utils.DB_MySQL.Model(&Borrower{}).Count(&total)
	offset := (page - 1) * pageSize
	query := utils.DB_MySQL.Model(&Borrower{}).Limit(pageSize).Offset(offset).Find(&employees)
	return employees, query
}