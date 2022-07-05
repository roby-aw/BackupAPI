package admin_test

import (
	"api-redeem-point/business/admin"
	"api-redeem-point/business/customermitra"
	"api-redeem-point/utils"
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"
)

var service admin.Service
var admin1, admin2, admin3, updateadmin admin.Admin
var customer1, customer2, customer3 customermitra.Customers
var store1, store2, store3 customermitra.Store
var history1, history2, history3 customermitra.History_Transaction
var InsertAdmin admin.Admin
var insertSpec, updateSpec, failedSpec, errorspec admin.RegisterAdmin
var stockProduct1, stockProduct2 admin.StockProduct
var TransactionMonth1, TransactionMonth2 admin.TransactionMonth
var pagination utils.Pagination
var loginadmin admin.AuthLogin
var stock1, stock2 admin.StockProduct
var historystore1, historystore2 customermitra.History_Transaction

var errorFindID int

var errorInsert error = errors.New("error on insert")
var errorFind error = errors.New("error on find")

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestGetAdminByID(t *testing.T) {
	t.Run("Expect found the result", func(t *testing.T) {
		result, _ := service.FindAdminByID(int(admin1.ID))
		if !reflect.DeepEqual(*result, admin1) {
			t.Error("Expect content has to be equal with content1", result, admin1)
		}
	})
	t.Run("Expect not found the result", func(t *testing.T) {
		result, err := service.FindAdminByID(int(100))
		if err != nil {
			t.Error("Expect error is nil. Error", err)
		} else if result != nil {
			t.Error("Expect result must be not found (nil)")
		}
	})
}
func TestInsertAdmin(t *testing.T) {
	t.Run("Expect found the result", func(t *testing.T) {
		result, err := service.CreateAdmin(&insertSpec)
		if err != nil {
			t.Error("Cannot insert admin")
		}
		if result.Email != insertSpec.Email {
			t.Error("Expect email admin is equal to email insert admin")
		}
		NewAdmin, _ := service.FindAdminByID(4)
		if NewAdmin == nil {
			t.Error("expect admins is not nil after inserted")
			t.FailNow()
		}
	})
}

func TestDashboard(t *testing.T) {
	t.Run("Expect found the result", func(t *testing.T) {
		result, _ := service.Dashboard()
		if result == nil {
			t.Error("expect result is nil")
		}
		if result.Month == nil {
			t.Error("expect data month is nil")
		}
		if result.Stock == nil {
			t.Error("expect stock is nil")
		}
		if len(result.Stock) != 2 {
			t.Error("expect lenght stock is 2", len(result.Stock))
		}
	})
}

func TestGetCustomers(t *testing.T) {
	t.Run("Expect found the result", func(t *testing.T) {
		result, _ := service.FindCustomers(pagination)
		if len(result) != 3 {
			t.Error("Expect found lenght customers is 3")
		}
		if result[0].ID != 1 {
			t.Error("Expect found id 1")
		}
	})
}

func TestDeleteCustomers(t *testing.T) {
	t.Run("Expect delete customer2", func(t *testing.T) {
		err := service.DeleteCustomer(int(customer2.ID))
		if err != nil {
			t.Error("error delete")
		}
		result, _ := service.FindCustomers(pagination)
		fmt.Println(result)
		if len(result) != 2 {
			t.Error("len customer must be 3")
		}
	})
}

func TestGetHistoryCustomers(t *testing.T) {
	t.Run("Expect get result history", func(t *testing.T) {
		result, _ := service.FindHistoryCustomers(pagination)
		if len(result) != 3 {
			t.Error("len history must be 3")
		}
		if result[1].Status_Transaction != history2.Status_Transaction {
			t.Error("Expect get description history 2")
		}
	})
}

func TestGetTransactionPending(t *testing.T) {
	t.Run("Expect get transaction pending", func(t *testing.T) {
		result, _ := service.TransactionPending(pagination)
		if len(result) != 2 {
			t.Error("len transaction pending must be 2")
		}
	})
}

func TestAcceptTransaction(t *testing.T) {
	t.Run("Expect can accept transaction pending", func(t *testing.T) {
		err := service.ApproveTransaction(history2.ID_Transaction)
		if err != nil {
			t.Error("error ", err)
		}
		result, _ := service.TransactionPending(pagination)
		if len(result) != 1 {
			t.Error("transaction not updated to completed")
		}
	})
}

func TestLoginAdmin(t *testing.T) {
	t.Run("Expect can login admin", func(t *testing.T) {
		result, _ := service.LoginAdmin(&loginadmin)
		if result == nil {
			t.Error("cannot login because no data")
		}
		if result.Email != loginadmin.Email {
			t.Error("error")
		}
	})
}

func TestGetStore(t *testing.T) {
	t.Run("Expect can get all store", func(t *testing.T) {
		result, _ := service.GetStore(pagination, "")
		if len(result) != 3 {
			t.Error("Expect len result is 3")
		}
		if result[store2.ID-1].Email != store2.Email {
			t.Error("Expect email store 2")
		}
	})
}

func setup() {
	admin1.ID = 1
	admin1.Email = "testemail@gmail.com"
	admin1.Fullname = "testname"
	admin1.Password = "testpassword"
	admin1.No_hp = "08565895685"

	admin2.ID = 2
	admin2.Email = "testemail2@gmail.com"
	admin2.Fullname = "testname2"
	admin2.Password = "testpassword"
	admin2.No_hp = "08565897685"

	admin3.ID = 3
	admin3.Email = "testemail3@gmail.com"
	admin3.Fullname = "testname3"
	admin3.Password = "testpassword"
	admin3.No_hp = "08565897665"

	stockProduct1.ID = 1
	stockProduct1.Product = "PulsaAndPaketData"
	stockProduct1.Balance = 500000

	stockProduct2.ID = 2
	stockProduct2.Product = "CashoutAndEmoney"
	stockProduct2.Balance = 500000

	TransactionMonth1.Day = "01"
	TransactionMonth2.Count = 5

	TransactionMonth1.Day = "02"
	TransactionMonth2.Count = 10

	customer1.ID = 1
	customer1.Email = "testcustomer1@gmail.com"
	customer1.Fullname = "testcustomer1"
	customer1.Password = "testpassword1"
	customer1.Pin = 1234
	customer1.Poin = 500000

	customer2.ID = 2
	customer2.Email = "testcustomer2@gmail.com"
	customer2.Fullname = "testcustomer2"
	customer2.Password = "testpassword2"
	customer2.Pin = 9876
	customer2.Poin = 30000

	customer3.ID = 3
	customer3.Email = "testcustomer3@gmail.com"
	customer3.Fullname = "testcustomer3"
	customer3.Password = "testpassword3"
	customer3.Pin = 6366
	customer3.Poin = 1000

	store1.ID = 1
	store1.Email = "store1@gmail.com"
	store1.Alamat = "jl. store 1"
	store1.Store = "store1"
	store1.Password = "passwordstore1"

	store2.ID = 2
	store2.Email = "store2@gmail.com"
	store2.Alamat = "jl. store 2"
	store2.Store = "store2"
	store2.Password = "passwordstore2"

	store3.ID = 3
	store3.Email = "store3@gmail.com"
	store3.Alamat = "jl. store 3"
	store3.Store = "store3"
	store3.Password = "passwordstore3"

	history1.ID = 1
	history1.ID_Transaction = "T12345"
	history1.Customer_id = 1
	history1.Transaction_type = "Redeem Pulsa"
	history1.Bank_Provider = "TELKOMSEL"
	history1.Nomor = "085696865698"
	history1.Poin_Account = 10000
	history1.Poin_Redeem = 10000
	history1.Amount = 10000
	history1.Description = "TELKOMSEL - 10000"
	history1.Status_Transaction = "COMPLETE"
	history1.Status_Poin = "OUT"

	history2.ID = 2
	history2.ID_Transaction = "T123565"
	history2.Customer_id = 1
	history2.Transaction_type = "Redeem Pulsa"
	history2.Bank_Provider = "TELKOMSEL"
	history2.Nomor = "08569686546"
	history2.Poin_Account = 10000
	history2.Poin_Redeem = 10000
	history2.Amount = 10000
	history2.Description = "TELKOMSEL - 10000"
	history2.Status_Transaction = "PENDING"
	history2.Status_Poin = "OUT"

	history3.ID = 3
	history3.ID_Transaction = "T136768"
	history3.Customer_id = 1
	history3.Transaction_type = "Redeem Pulsa"
	history3.Bank_Provider = "TELKOMSEL"
	history3.Nomor = "08569686546"
	history3.Poin_Account = 10000
	history3.Poin_Redeem = 10000
	history3.Amount = 10000
	history3.Description = "TELKOMSEL - 10000"
	history3.Status_Transaction = "PENDING"
	history3.Status_Poin = "OUT"

	pagination.Limit = 1000
	pagination.Page = 1

	repo := newInMemoryRepository()
	service = admin.NewService(&repo)

	insertSpec.Email = "insertadmin@gmail.com"
	insertSpec.Fullname = "insertfullname"
	insertSpec.No_hp = "0854696963"
	insertSpec.Password = "insertpassword"

	loginadmin.Email = customer1.Email
	loginadmin.Password = customer1.Password
}

type inMemoryRepository struct {
	Admin            map[int]admin.Admin
	AllAdmin         []admin.Admin
	Product          map[int]admin.StockProduct
	AllProduct       []admin.StockProduct
	TransactionMonth []admin.TransactionMonth
	Customer         map[int]customermitra.Customers
	AllCustomer      []customermitra.Customers
	History          map[string]customermitra.History_Transaction
	AllHistory       []customermitra.History_Transaction
	Store            map[int]customermitra.Store
	AllStore         []customermitra.Store
}

func newInMemoryRepository() inMemoryRepository {
	var repo inMemoryRepository
	repo.Admin = make(map[int]admin.Admin)
	repo.Admin[int(admin1.ID)] = admin1
	repo.Admin[int(admin2.ID)] = admin2
	repo.Admin[int(admin3.ID)] = admin3

	repo.Product = make(map[int]admin.StockProduct)
	repo.Product[stockProduct1.ID] = stockProduct1
	repo.Product[stockProduct2.ID] = stockProduct2

	repo.AllAdmin = []admin.Admin{}
	repo.AllAdmin = append(repo.AllAdmin, admin1)
	repo.AllAdmin = append(repo.AllAdmin, admin2)
	repo.AllAdmin = append(repo.AllAdmin, admin3)

	repo.AllProduct = []admin.StockProduct{}
	repo.AllProduct = append(repo.AllProduct, stockProduct1)
	repo.AllProduct = append(repo.AllProduct, stockProduct2)

	repo.TransactionMonth = []admin.TransactionMonth{}
	repo.TransactionMonth = append(repo.TransactionMonth, TransactionMonth1)
	repo.TransactionMonth = append(repo.TransactionMonth, TransactionMonth2)

	repo.Customer = make(map[int]customermitra.Customers)
	repo.Customer[int(customer1.ID)] = customer1
	repo.Customer[int(customer2.ID)] = customer2
	repo.Customer[int(customer3.ID)] = customer3

	repo.AllCustomer = []customermitra.Customers{}
	repo.AllCustomer = append(repo.AllCustomer, customer1)
	repo.AllCustomer = append(repo.AllCustomer, customer2)
	repo.AllCustomer = append(repo.AllCustomer, customer3)

	repo.History = make(map[string]customermitra.History_Transaction)
	repo.History[history1.ID_Transaction] = history1
	repo.History[history2.ID_Transaction] = history2
	repo.History[history3.ID_Transaction] = history3

	repo.AllHistory = []customermitra.History_Transaction{}
	repo.AllHistory = append(repo.AllHistory, history1)
	repo.AllHistory = append(repo.AllHistory, history2)
	repo.AllHistory = append(repo.AllHistory, history3)

	repo.Store = make(map[int]customermitra.Store)
	repo.Store[int(store1.ID)] = store1
	repo.Store[int(store2.ID)] = store2
	repo.Store[int(store3.ID)] = store3

	repo.AllStore = []customermitra.Store{}
	repo.AllStore = append(repo.AllStore, store1)
	repo.AllStore = append(repo.AllStore, store2)
	repo.AllStore = append(repo.AllStore, store3)

	return repo
}

func (repo *inMemoryRepository) GetAdminByID(id int) (*admin.Admin, error) {
	if id == errorFindID {
		return nil, errorFind
	}

	admin, ok := repo.Admin[id]
	if !ok {
		return nil, nil
	}
	return &admin, nil
}

func (repo *inMemoryRepository) InsertAdmin(admins *admin.RegisterAdmin) (*admin.RegisterAdmin, error) {
	if admins.Fullname == errorspec.Fullname {
		return nil, errorInsert
	}
	adminInsert := admin.Admin{
		ID:       4,
		Email:    admins.Email,
		Fullname: admins.Fullname,
		Password: admins.Password,
	}
	repo.AllAdmin = append(repo.AllAdmin, adminInsert)
	repo.Admin[int(adminInsert.ID)] = adminInsert

	return admins, nil
}

func (repo *inMemoryRepository) Dashboard() (*int, error) {
	today := len(repo.AllHistory)
	return &today, nil
}

func (repo *inMemoryRepository) GetProduct() ([]admin.StockProduct, error) {
	stock := repo.AllProduct
	return stock, nil
}

func (repo *inMemoryRepository) GetTransactionMonthDay() ([]admin.TransactionMonth, error) {
	month := repo.TransactionMonth
	return month, nil
}

func (repo *inMemoryRepository) TransactionPending(pagination utils.Pagination) ([]*admin.TransactionPending, error) {
	data := repo.AllHistory
	var history []*admin.TransactionPending
	for _, v := range data {
		if v.Status_Transaction == "PENDING" {
			var tmpHistory admin.TransactionPending
			tmpHistory.ID_Transaction = v.ID_Transaction
			tmpHistory.Nomor = v.Nomor
			tmpHistory.Customer_id = v.Customer_id
			tmpHistory.Description = v.Description
			tmpHistory.Status_transaction = v.Status_Transaction
			history = append(history, &tmpHistory)
		}
	}
	return history, nil
}

func (repo *inMemoryRepository) AcceptTransaction(idtransaction string) error {
	if repo.History[idtransaction].Status_Transaction == "PENDING" {
		if history, ok := repo.History[idtransaction]; ok {
			history.Status_Transaction = "COMPLETED"
			repo.History[idtransaction] = history
			for _, v := range repo.AllHistory {
				if v.ID_Transaction == idtransaction {
					repo.AllHistory[v.ID-1].Status_Transaction = "COMPLETED"
				}
			}
			return nil
		}
	}
	return errors.New("status transaction not pending")
}
func (repo *inMemoryRepository) LoginAdmin(Auth *admin.AuthLogin) (*admin.ResponseLogin, error) {
	var data admin.ResponseLogin
	for _, v := range repo.AllCustomer {
		if v.Email == Auth.Email {
			if v.Password == Auth.Password {
				data.ID = int(v.ID)
				data.Email = v.Email
				data.Fullname = v.Fullname
				data.No_hp = v.No_hp
				data.Password = v.Password
				data.Token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MywiRW1haWwiOiJ0ZXN0MUBnbWFpbC5jb20iLCJDdXN0b21lciI6dHJ1ZSwiZXhwIjoxNjU2ODcxMzEzfQ.a9O_RMxG7iJR4tMdBZmL6JY2lZKsiQv3rSegnhv1C00"
			}
		}
	}
	var dataFix *admin.ResponseLogin
	dataFix = &data
	return dataFix, nil
}
func (repo *inMemoryRepository) RenewAdmin(id int, admin *admin.Admin) (*admin.Admin, error) {
	return nil, nil
}
func (repo *inMemoryRepository) GetCustomers(pagination utils.Pagination) ([]*customermitra.Customers, error) {
	customers := repo.AllCustomer
	var tmpcustomer []*customermitra.Customers
	for _, v := range customers {
		var tmp customermitra.Customers
		tmp = v
		tmpcustomer = append(tmpcustomer, &tmp)
	}
	return tmpcustomer, nil
}
func (repo *inMemoryRepository) GetHistoryCustomers(pagination utils.Pagination) ([]admin.CustomerHistory, error) {
	history := repo.AllHistory
	var data []admin.CustomerHistory
	for _, v := range history {
		var tmpHistory admin.CustomerHistory
		tmpHistory.Customer_id = v.Customer_id
		tmpHistory.Description = v.Description
		tmpHistory.Nomor = v.Nomor
		tmpHistory.Status_Transaction = v.Status_Transaction
		tmpHistory.Poin_redeem = v.Poin_Redeem
		data = append(data, tmpHistory)
	}
	return data, nil
}
func (repo *inMemoryRepository) DeleteCustomer(id int) error {
	id = id - 1
	repo.AllCustomer = append(repo.AllCustomer[:id], repo.AllCustomer[id+1:]...)
	return nil
}
func (repo *inMemoryRepository) TransactionDate() ([]admin.TransactionDate, error) {
	return nil, nil
}
func (repo *inMemoryRepository) TransactionByDate(startdate string, enddate string) ([]admin.TransactionDate, error) {
	return nil, nil
}
func (repo *inMemoryRepository) UpdateCustomer(data admin.UpdateCustomer) (*admin.UpdateCustomer, error) {
	return nil, nil
}
func (repo *inMemoryRepository) UpdateCustomerPoint(id int, point int) (*int, error) {
	return nil, nil
}
func (repo *inMemoryRepository) UpdateStock(id int, stock int) (*admin.StockProduct, error) {
	return nil, nil
}
func (repo *inMemoryRepository) HistoryStore(pagination utils.Pagination, name string) ([]admin.HistoryStore, error) {
	return nil, nil
}
func (repo *inMemoryRepository) DeleteStore(id int) error {
	return nil
}
func (repo *inMemoryRepository) GetStore(pagination utils.Pagination, name string) ([]*customermitra.Store, error) {
	data := repo.AllStore
	var tmpStore []*customermitra.Store
	for _, v := range data {
		var tmpData customermitra.Store
		tmpData = v
		tmpStore = append(tmpStore, &tmpData)
	}
	return tmpStore, nil
}
func (repo *inMemoryRepository) UpdateStore(store admin.UpdateStore) (*admin.UpdateStore, error) {
	return nil, nil
}
