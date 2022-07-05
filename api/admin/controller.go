package admin

import (
	"api-redeem-point/business/admin"
	adminBusiness "api-redeem-point/business/admin"
	"api-redeem-point/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service adminBusiness.Service
}

func NewController(service adminBusiness.Service) *Controller {
	return &Controller{
		service: service,
	}
}

// Create godoc
// @Summary Dashboard Admin
// @description Dashboard for admin
// @tags Admin
// @Accept json
// @Produce json
// @Success 200	{object} response.Result
// @Failure 400 {object} response.Error
// @Router /admin/{id} [get]
func (Controller *Controller) FindAdminByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result, err := Controller.service.FindAdminByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success get admin by id",
		"result":   result,
	})
}

func (Controller *Controller) Dashboard(c echo.Context) error {
	result, err := Controller.service.Dashboard()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success get dashboard",
		"result":   result,
	})
}

// Create godoc
// @Summary Transaction Pending
// @description Transaction Pending for Admin
// @tags Admin
// @Accept json
// @Produce json
// @Success 200	{object} response.Result
// @Failure 400 {object} response.Error
// @Router /admin/transaction/pending [get]
func (Controller *Controller) TransactionPending(c echo.Context) error {
	pagination := utils.GeneratePagination(c.Request().URL.Query())
	result, err := Controller.service.TransactionPending(pagination)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":       200,
		"messages":   "success get transaction pending",
		"pagination": pagination,
		"result":     result,
	})
}

// Create godoc
// @Summary Create admin
// @description create admin with data
// @tags Admin
// @Accept json
// @Produce json
// @Param admin body admin.RegisterAdmin true "admin"
// @Success 200	{object} response.Result
// @Failure 400 {object} response.Error
// @Router /admin [post]
func (Controller *Controller) CreateAdmin(c echo.Context) error {
	admin := adminBusiness.RegisterAdmin{}
	c.Bind(&admin)
	admins, err := Controller.service.CreateAdmin(&admin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success create data",
		"data":     admins,
	})
}

// Create godoc
// @Summary Login admin
// @description Login admin
// @tags Admin
// @Accept json
// @Produce json
// @Param admin body admin.AuthLogin true "admin"
// @Success 200	{object} response.Result
// @Failure 400 {object} response.Error
// @Router /admin/login [post]
func (Controller *Controller) LoginAdmin(c echo.Context) error {
	var request adminBusiness.AuthLogin
	c.Bind(&request)
	result, err := Controller.service.LoginAdmin(&request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "success login",
		"result":  result,
	})
}

func (Controller *Controller) UpdateAdmin(c echo.Context) error {
	var admin *admin.Admin
	id, _ := strconv.Atoi(c.Param("id"))
	c.Bind(&admin)
	admin, err := Controller.service.UpdateAdmin(id, admin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success update data admin",
		"data":     admin,
	})
}

// Create godoc
// @Summary Approve Transaction
// @description Approve Transaction
// @tags Admin
// @Accept json
// @Produce json
// @Param transactionid path string true "transaction_id"
// @Success 200	{object} response.Result
// @Failure 400 {object} response.Error
// @Router /admin/approve/{transactionid} [post]
func (Controller *Controller) ApproveTransaction(c echo.Context) error {
	transactionid := c.Param("idtransaction")
	err := Controller.service.ApproveTransaction(transactionid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success approve transaction",
	})
}

// Create godoc
// @Summary Find customers
// @description Find customers for admin
// @tags Admin
// @Accept json
// @Produce json
// @Param limit query int true "pagination"
// @Param page query int true "pagination"
// @Success 200	{object} response.Result
// @Failure 400 {object} response.Error
// @Router /admin/customer [get]
func (Controller *Controller) FindCustomers(c echo.Context) error {
	pagination := utils.GeneratePagination(c.Request().URL.Query())
	result, err := Controller.service.FindCustomers(pagination)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":       200,
		"messages":   "success get customers",
		"pagination": pagination,
		"result":     result,
	})
}

// Create godoc
// @Summary Find History Customers
// @description Find History Customers for admin
// @tags Admin
// @Accept json
// @Produce json
// @Param limit query int true "pagination"
// @Param page query int true "pagination"
// @Success 200	{object} response.Result
// @Failure 400 {object} response.Error
// @Router /admin/history [get]
func (Controller *Controller) FindHistoryCustomers(c echo.Context) error {
	pagination := utils.GeneratePagination(c.Request().URL.Query())
	result, err := Controller.service.FindHistoryCustomers(pagination)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":       200,
		"messages":   "success get customers",
		"pagination": pagination,
		"result":     result,
	})
}

// Create godoc
// @Summary Delete Customers
// @description Delete Customers for admin
// @tags Admin
// @Accept json
// @Produce json
// @Param id query int true "id customer"
// @Success 200	{object} response.Result
// @Failure 400 {object} response.Error
// @Router /admin/customer [delete]
func (Controller *Controller) DeleteCustomer(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))
	err := Controller.service.DeleteCustomer(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success delete customers",
		"result":   id,
	})
}

// Create godoc
// @Summary Transaction By Date
// @description Transaction By Date for admin
// @tags Admin
// @Accept json
// @Produce json
// @Param start query string false "2022-01-01"
// @Param end query string false "2022-12-31"
// @Success 200	{object} response.Result
// @Failure 400 {object} response.Error
// @Router /admin/transaction [get]
func (Controller *Controller) TransactionByDate(c echo.Context) error {
	start := c.QueryParam("start")
	end := c.QueryParam("end")
	if start == "" {
		start = "2022-01-01"
	}
	if end == "" {
		end = time.Now().Format("2006-01-02")
	}
	result, err := Controller.service.TransactionByDate(start, end)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success get transaction by date",
		"result":   result,
	})
}

// Create godoc
// @Summary Transaction By Date
// @description Transaction By Date for admin
// @tags Admin
// @Accept json
// @Produce json
// @Param admin body admin.UpdateCustomer true "admin"
// @Success 200	{object} response.Result
// @Failure 400 {object} response.Error
// @Router /admin/customer [put]
func (Controller *Controller) UpdateCustomer(c echo.Context) error {
	var req adminBusiness.UpdateCustomer
	c.Bind(&req)
	result, err := Controller.service.UpdateCustomer(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success update customer",
		"result":   result,
	})

}

// Create godoc
// @Summary Update customer point
// @description Update customer point for admin
// @tags Admin
// @Accept json
// @Produce json
// @Param id query int true "id customer"
// @Param point query int true "point customer"
// @Success 200	{object} response.Result
// @Failure 400 {object} response.Error
// @Router /admin/customer/point [put]
func (Controller *Controller) UpdateCustomerPoint(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))
	point, _ := strconv.Atoi(c.QueryParam("point"))
	result, err := Controller.service.UpdateCustomerPoint(id, point)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success add point",
		"result":   result,
	})
}

// Create godoc
// @Summary Stock Product
// @description Stock Product for admin
// @tags Admin
// @Accept json
// @Produce json
// @Success 200	{object} response.Result
// @Failure 400 {object} response.Error
// @Router /admin/stock [get]
func (Controller *Controller) StockProduct(c echo.Context) error {
	result, err := Controller.service.FindProduct()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success get product",
		"result":   result,
	})
}

// Create godoc
// @Summary Update Stock Product
// @description Update Stock Product for admin
// @tags Admin
// @Accept json
// @Produce json
// @Param id query int true "id product"
// @Param balance query int true "balance product"
// @Success 200	{object} response.Result
// @Failure 400 {object} response.Error
// @Router /admin/stock [put]
func (Controller *Controller) UpdateStock(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))
	balance, _ := strconv.Atoi(c.QueryParam("balance"))
	result, err := Controller.service.UpdateStock(id, balance)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success update product",
		"result":   result,
	})
}

// Create godoc
// @Summary History Store
// @description History Store for admin
// @tags Admin
// @Accept json
// @Produce json
// @Param limit query int true "pagination"
// @Param page query int true "pagination"
// @Success 200	{object} response.Result
// @Failure 400 {object} response.Error
// @Router /admin/historystore [get]
func (Controller *Controller) HistoryStore(c echo.Context) error {
	name := c.QueryParam("name")
	pagination := utils.GeneratePagination(c.Request().URL.Query())
	result, err := Controller.service.HistoryStore(pagination, name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":       200,
		"messages":   "success get history store",
		"pagination": pagination,
		"result":     result,
	})
}

// Create godoc
// @Summary Delete Store
// @description Delete Store for admin
// @tags Admin
// @Accept json
// @Produce json
// @Param id query int true "id store"
// @Success 200	{object} response.Result
// @Failure 400 {object} response.Error
// @Router /admin/store [delete]
func (Controller *Controller) DeleteStore(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))
	err := Controller.service.DeleteStore(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success delete store",
		"result":   id,
	})
}

func (Controller *Controller) GetStore(c echo.Context) error {
	name := c.QueryParam("name")
	pagination := utils.GeneratePagination(c.Request().URL.Query())
	result, err := Controller.service.GetStore(pagination, name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success get store",
		"result":   result,
	})
}

func (Controller *Controller) UpdateStore(c echo.Context) error {
	var Store adminBusiness.UpdateStore
	c.Bind(&Store)
	result, err := Controller.service.UpdateStore(Store)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success update store",
		"result":   result,
	})
}
