package repository

import "gorm.io/gorm"

type History_Transaction struct {
	gorm.Model
	ID_Transaction     string   `gorm:"primaryKey;autoIncrement:false"`
	Customer_id        int      `json:"customer_id"`
	Customers          Customer `json:"customers" gorm:"foreignkey:ID;references:Customer_id"`
	Store_id           int      `json:"store_id"`
	Store              Store    `json:"store" gorm:"foreignkey:ID;references:Store_id"`
	Transaction_type   string   `json:"transaction_type"`
	Bank_Provider      string   `json:"bank_provider" gorm:"size:255"`
	Nomor              string   `json:"nomor" gorm:"size:20"`
	Poin_Account       int      `json:"poin_account"`
	Poin_Redeem        int      `json:"poin_redeem"`
	Amount             int      `json:"amount"`
	Description        string   `json:"description" gorm:"size:255"`
	Status_Transaction string   `json:"status_transaction" gorm:"size:255"`
	Status_Poin        string   `json:"status_poin" gorm:"size:10"`
}

type Store struct {
	gorm.Model
	Email    string `json:"email" gorm:"size:255"`
	Password string `json:"password" gorm:"size:255"`
	Store    string `json:"store" gorm:"size:255"`
	Alamat   string `json:"alamat" gorm:"size:255"`
}

type Customer struct {
	gorm.Model
	Email    string `json:"email" gorm:"primaryKey;autoIncrement:false"`
	Fullname string `json:"fullname" gorm:"size:255"`
	Password string `json:"password" gorm:"size:255"`
	No_hp    string `json:"no_hp" gorm:"size:80"`
	Poin     int    `json:"poin" gorm:"size:50"`
	Pin      int    `json:"pin" gorm:"size:50"`
}

type StockProduct struct {
	gorm.Model
	Product string `json:"product" gorm:"size:100"`
	Balance int    `json:"balance" gorm:"size:100"`
}

type Admin struct {
	gorm.Model
	Email    string `json:"email" gorm:"size:255"`
	Fullname string `json:"fullname" gorm:"size:255"`
	Password string `json:"password" gorm:"size:255"`
	No_hp    string `json:"no_hp" gorm:"size:80"`
}
