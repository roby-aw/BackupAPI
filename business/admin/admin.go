package admin

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AdminSwagger struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterAdmin struct {
	Email    string `json:"email" validate:"required,email"`
	Fullname string `json:"fullname" validate:"required"`
	Password string `json:"password" validate:"required"`
	No_hp    string `json:"no_hp" validate:"required"`
}

type Admin struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
	Email     string    `json:"email" validate:"required,email"`
	Fullname  string    `json:"fullname" validate:"required"`
	Password  string    `json:"password" validate:"required"`
	No_hp     string    `json:"no_hp" validate:"required"`
}

type CustomerHistory struct {
	Customer_id        int       `json:"customer_id"`
	Customers          Customers `json:"customers" gorm:"foreignkey:ID;references:Customer_id"`
	Description        string    `json:"description"`
	Nomor              string    `json:"nomor"`
	CreatedAt          time.Time `json:"createdat"`
	Status_Transaction string    `json:"status_transaction"`
	Poin_redeem        int       `json:"poin_redeem"`
}

type AuthLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type TransactionPending struct {
	ID_Transaction     string    `json:"id_transaction"`
	Nomor              string    `json:"nomor"`
	Customer_id        int       `json:"customer_id"`
	Customers          Customers `json:"customer" gorm:"foreignkey:ID;references:Customer_id"`
	Description        string    `json:"description"`
	Status_transaction string    `json:"status_transaction"`
}

type Customers struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Password string `json:"password"`
}

type ResponseLogin struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
	No_hp    string `json:"no_hp"`
	Token    string `json:"token"`
}

type InputAdminToken struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Kota struct {
	ID                 int    `json:"id"`
	Rajaongkir_city_id int    `json:"rajaongkir_city_id" validate:"required"`
	Kota_Nama          string `json:"kota_nama" validate:"required"`
	Postal_code        int    `json:"postal_code" validate:"required"`
	Tipe               string `json:"tipe" validate:"required"`
	Province_ID        int    `json:"province_id" validate:"required"`
}

type TransactionDate struct {
	ID               int       `json:"id"`
	Transaction_type string    `json:"transaction_type"`
	CreatedAt        time.Time `json:"createdat"`
}
type StockProduct struct {
	ID      int    `json:"id"`
	Product string `json:"product" gorm:"size:100"`
	Balance int    `json:"balance" gorm:"size:100"`
}

type Auth struct {
	Token string
}

type Claims struct {
	ID    int
	Email string
	Role  string
	jwt.StandardClaims
}

type TransactionMonth struct {
	Day   string `json:"Day"`
	Count int    `json:"count"`
}

type Dashboard struct {
	Today int                `json:"today"`
	Stock []StockProduct     `json:"stock"`
	Month []TransactionMonth `json:"month"`
}

type HistoryStore struct {
	Store_id    int       `json:"store_id"`
	Store       Store     `json:"store" gorm:"foreignkey:ID;references:Store_id"`
	Customer_id int       `json:"customer_id"`
	Customers   Customers `json:"customer" gorm:"foreignkey:ID;references:Customer_id"`
	CreatedAt   time.Time `json:"createdat"`
	Poin_Redeem int       `json:"poin_redeem"`
	Amount      int       `json:"amount"`
}

type Store struct {
	ID     uint   `json:"id"`
	Email  string `json:"email"`
	Store  string `json:"store"`
	Alamat string `json:"alamat"`
}

type UpdateCustomer struct {
	ID       int    `json:"id" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Fullname string `json:"fullname" validate:"required"`
	Password string `json:"password"`
	No_hp    string `json:"no_hp"`
	Pin      int    `json:"pin"`
}

type UpdateStore struct {
	ID       uint   `json:"id" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
	Store    string `json:"store" validate:"required"`
	Alamat   string `json:"alamat" validate:"required"`
}
