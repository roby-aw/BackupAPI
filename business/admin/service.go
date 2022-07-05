package admin

import (
	"api-redeem-point/business/customermitra"
	"api-redeem-point/utils"
	"errors"

	"github.com/go-playground/validator/v10"
)

type Repository interface {
	GetAdminByID(id int) (*Admin, error)
	Dashboard() (*int, error)
	TransactionPending(pagination utils.Pagination) ([]*TransactionPending, error)
	InsertAdmin(admin *RegisterAdmin) (*RegisterAdmin, error)
	AcceptTransaction(idtransaction string) error
	RenewAdmin(id int, admin *Admin) (*Admin, error)
	LoginAdmin(Auth *AuthLogin) (*ResponseLogin, error)
	GetCustomers(pagination utils.Pagination) ([]*customermitra.Customers, error)
	GetHistoryCustomers(pagination utils.Pagination) ([]CustomerHistory, error)
	DeleteCustomer(id int) error
	TransactionDate() ([]TransactionDate, error)
	TransactionByDate(startdate string, enddate string) ([]TransactionDate, error)
	UpdateCustomer(data UpdateCustomer) (*UpdateCustomer, error)
	UpdateCustomerPoint(id int, point int) (*int, error)
	GetProduct() ([]StockProduct, error)
	UpdateStock(id int, stock int) (*StockProduct, error)
	GetTransactionMonthDay() ([]TransactionMonth, error)
	HistoryStore(pagination utils.Pagination, name string) ([]HistoryStore, error)
	DeleteStore(id int) error
	GetStore(pagination utils.Pagination, name string) ([]*customermitra.Store, error)
	UpdateStore(store UpdateStore) (*UpdateStore, error)
}

type Service interface {
	FindAdminByID(id int) (*Admin, error)
	Dashboard() (*Dashboard, error)
	TransactionPending(pagination utils.Pagination) ([]*TransactionPending, error)
	CreateAdmin(admin *RegisterAdmin) (*RegisterAdmin, error)
	ApproveTransaction(idtransaction string) error
	UpdateAdmin(id int, admin *Admin) (*Admin, error)
	LoginAdmin(Auth *AuthLogin) (*ResponseLogin, error)
	FindCustomers(pagination utils.Pagination) ([]*customermitra.Customers, error)
	FindHistoryCustomers(pagination utils.Pagination) ([]CustomerHistory, error)
	DeleteCustomer(id int) error
	TransactionDate() ([]TransactionDate, error)
	TransactionByDate(startdate string, enddate string) ([]TransactionDate, error)
	UpdateCustomer(data UpdateCustomer) (*UpdateCustomer, error)
	UpdateCustomerPoint(id int, point int) (*int, error)
	FindProduct() ([]StockProduct, error)
	UpdateStock(id int, stock int) (*StockProduct, error)
	GetTransactionMonthDay() ([]TransactionMonth, error)
	HistoryStore(pagination utils.Pagination, name string) ([]HistoryStore, error)
	DeleteStore(id int) error
	GetStore(pagination utils.Pagination, name string) ([]*customermitra.Store, error)
	UpdateStore(store UpdateStore) (*UpdateStore, error)
}

type service struct {
	repository Repository
	validate   *validator.Validate
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
		validate:   validator.New(),
	}
}

func (s *service) FindAdminByID(id int) (*Admin, error) {
	return s.repository.GetAdminByID(id)
}

func (s *service) Dashboard() (*Dashboard, error) {
	transMonts, err := s.repository.GetTransactionMonthDay()
	if err != nil {
		return nil, err
	}
	Stock, err := s.repository.GetProduct()
	today, err := s.repository.Dashboard()
	Dashboard := Dashboard{
		Today: *today,
		Stock: Stock,
		Month: transMonts,
	}
	return &Dashboard, nil
}

func (s *service) TransactionPending(pagination utils.Pagination) ([]*TransactionPending, error) {
	return s.repository.TransactionPending(pagination)
}

func (s *service) CreateAdmin(admin *RegisterAdmin) (*RegisterAdmin, error) {
	err := s.validate.Struct(admin)
	if err != nil {
		return nil, err
	}
	admin, err = s.repository.InsertAdmin(admin)
	return admin, err
}

func (s *service) ApproveTransaction(idtransaction string) error {
	return s.repository.AcceptTransaction(idtransaction)
}

func (s *service) UpdateAdmin(id int, admin *Admin) (*Admin, error) {
	return s.repository.RenewAdmin(id, admin)
}

func (s *service) LoginAdmin(Auth *AuthLogin) (*ResponseLogin, error) {
	err := s.validate.Struct(Auth)
	if err != nil {
		return nil, err
	}
	tokens, err := s.repository.LoginAdmin(Auth)
	return tokens, err
}

func (s *service) FindCustomers(pagination utils.Pagination) ([]*customermitra.Customers, error) {
	return s.repository.GetCustomers(pagination)
}

func (s *service) FindHistoryCustomers(pagination utils.Pagination) ([]CustomerHistory, error) {
	return s.repository.GetHistoryCustomers(pagination)
}

func (s *service) DeleteCustomer(id int) error {
	if id == 0 {
		err := errors.New("Masukkan id customer")
		return err
	}
	return s.repository.DeleteCustomer(id)
}

func (s *service) TransactionDate() ([]TransactionDate, error) {
	return s.repository.TransactionDate()
}

func (s *service) TransactionByDate(startdate string, enddate string) ([]TransactionDate, error) {
	return s.repository.TransactionByDate(startdate, enddate)
}

func (s *service) UpdateCustomer(data UpdateCustomer) (*UpdateCustomer, error) {
	err := s.validate.Struct(data)
	if err != nil {
		return nil, err
	}
	return s.repository.UpdateCustomer(data)
}

func (s *service) UpdateCustomerPoint(id int, point int) (*int, error) {
	return s.repository.UpdateCustomerPoint(id, point)
}

func (s *service) FindProduct() ([]StockProduct, error) {
	return s.repository.GetProduct()
}

func (s *service) UpdateStock(id int, stock int) (*StockProduct, error) {
	return s.repository.UpdateStock(id, stock)
}

func (s *service) GetTransactionMonthDay() ([]TransactionMonth, error) {
	return s.repository.GetTransactionMonthDay()
}

func (s *service) HistoryStore(pagination utils.Pagination, name string) ([]HistoryStore, error) {
	return s.repository.HistoryStore(pagination, name)
}

func (s *service) DeleteStore(id int) error {
	if id == 0 {
		err := errors.New("Masukkan id customer")
		return err
	}
	return s.repository.DeleteStore(id)
}

func (s *service) GetStore(pagination utils.Pagination, name string) ([]*customermitra.Store, error) {
	return s.repository.GetStore(pagination, name)
}

func (s *service) UpdateStore(store UpdateStore) (*UpdateStore, error) {
	err := s.validate.Struct(store)
	if err != nil {
		return nil, err
	}
	return s.repository.UpdateStore(store)
}
