package customermitra

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type History_Transaction struct {
	ID                 uint      `json:"id"`
	CreatedAt          time.Time `json:"createdat"`
	UpdatedAt          time.Time `json:"updatedat"`
	ID_Transaction     string    `json:"id_transaction"`
	Customer_id        int       `json:"customer_id"`
	Customers          Customers `json:"customers" gorm:"foreignkey:ID;references:Customer_id"`
	Store_id           int       `json:"store_id"`
	Store              Store     `json:"store" gorm:"foreignkey:ID;references:Store_id"`
	Transaction_type   string    `json:"transaction_type"`
	Bank_Provider      string    `json:"bank_provider"`
	Nomor              string    `json:"nomor"`
	Poin_Account       int       `json:"poin_account"`
	Poin_Redeem        int       `json:"poin_redeem"`
	Amount             int       `json:"amount"`
	Description        string    `json:"description"`
	Status_Transaction string    `json:"status_transaction"`
	Status_Poin        string    `json:"status_poin"`
}

type Store struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Store    string `json:"store"`
	Alamat   string `json:"alamat"`
}

type RegisterStore struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Store    string `json:"store" validate:"required"`
	Alamat   string `json:"alamat" validate:"required"`
}

type AuthStore struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ResponseLoginStore struct {
	Store Store  `json:"store"`
	Token string `json:"token"`
}

type Customers struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	Email     string `json:"email"`
	Fullname  string `json:"fullname"`
	Password  string `json:"password"`
	No_hp     string `json:"no_hp"`
	Poin      int    `json:"poin"`
	Pin       int    `json:"pin"`
}
type UpdateCustomers struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Password string `json:"password"`
	No_hp    string `json:"no_hp"`
	Poin     int    `json:"poin"`
	Pin      int    `json:"pin"`
}

type StockProduct struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdat"`
	Product   string    `json:"product" gorm:"size:100"`
	Balance   int       `json:"balance" gorm:"size:100"`
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

type RegisterCustomer struct {
	Email    string `json:"email" validate:"required,email"`
	Fullname string `json:"fullname" validate:"required"`
	Password string `json:"password" validate:"required"`
	No_hp    string `json:"no_hp" validate:"required,numeric"`
	Pin      int    `json:"pin" validate:"required,lt=5,numeric"`
}

type Login struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	Password string `json:"password"`
	Poin     int    `json:"poin"`
	Pin      int    `json:"pin"`
}

type ResponseLogin struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Password string `json:"password"`
	No_hp    string `json:"no_hp"`
	Poin     int    `json:"poin"`
	Pin      int    `json:"pin"`
	Token    string `json:"token"`
}

type AuthLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type DetailHistory struct {
	ID_Transaction     string    `json:"transaction_id"`
	Transaction_type   string    `json:"transaction_type"`
	CreatedAt          time.Time `json:"createdat"`
	Bank_Provider      string    `json:"bank_provider"`
	Nomor              string    `json:"nomor"`
	Amount             int       `json:"amount"`
	Poin_account       int       `json:"poin_account"`
	Poin_redeem        int       `json:"poin_redeem"`
	Description        string    `json:"description"`
	Status_Transaction string    `json:"status_transaction"`
}
type History struct {
	ID                 int       `json:"id"`
	ID_Transaction     string    `json:"id_transaction"`
	Transaction_type   string    `json:"transaction_type"`
	CreatedAt          time.Time `json:"createdat"`
	Status_Transaction string    `json:"status_transaction"`
}

type RedeemPulsaData struct {
	Customer_id   int    `json:"customer_id" validate:"required"`
	Bank_Provider string `json:"bank_provider" validate:"required"`
	Nomor         string `json:"nomor" validate:"required,numeric"`
	Poin_redeem   int    `json:"poin_redeem" validate:"required"`
	Amount        int    `json:"amount" validate:"required"`
}

type UpdateCustomer struct {
	ID    int    `json:"id" validate:"required"`
	Name  string `json:"name"`
	Email string `json:"email"`
	No_hp string `json:"no_hp"`
}

type Bank struct {
	BankCode    string `json:"bankcode"`
	No_rekening string `json:"no_rekening"`
	Amount      int    `json:"amount"`
}

type Disbursement struct {
	UserID                  string `json:"user_id"`
	ExternalID              string `json:"external_id"`
	Amount                  int    `json:"amount"`
	BankCode                string `json:"bank_code"`
	AccountHolderName       string `json:"account_holder_name"`
	DisbursementDescription string `json:"disbursement_description"`
	Status                  string `json:"status"`
	ID                      string `json:"id"`
}

type TransactionBank struct {
	ID                uint64 `gorm:"primaryKey"`
	ID_Transaction    string `json:"id_transaction"`
	ID_User           string `json:"id_user"`
	Jenis_transaction string `json:"jenis_transaction" validate:"required"`
	Nama_bank         string `json:"nama_bank" validate:"required"`
	AN_Bank           string `json:"AN_Bank" validate:"required"`
	No_rekening       string `json:"no_rekening" validate:"required"`
	Amount            int    `json:"amount"`
	Status            string `json:"status"`
}

type InputTransactionBankEmoney struct {
	Customer_id   int    `json:"customer_id" validate:"required,numeric"`
	Bank_Provider string `json:"bank_provider" validate:"required"`
	AN_Rekening   string `json:"an_rekening" validate:"required"`
	Nomor         string `json:"nomor" validate:"required,numeric"`
	Amount        int    `json:"amount" validate:"required,numeric"`
	Poin_redeem   int    `json:"poin_redeem" validate:"required,numeric"`
}

type InputPoin struct {
	Customer_id int `json:"customer_id" validate:"required"`
	Store_id    int `json:"store_id" validate:"required"`
	Amount      int `json:"amount" validate:"required"`
}

type Claims struct {
	ID       int
	Email    string
	Customer bool
	jwt.StandardClaims
}

type ClaimsMitra struct {
	ID    int
	Email string
	Store bool
	jwt.StandardClaims
}
