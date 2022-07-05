package customermitra

import (
	customermitraBussiness "api-redeem-point/business/customermitra"
	"api-redeem-point/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service customermitraBussiness.Service
}

func NewController(service customermitraBussiness.Service) *Controller {
	return &Controller{
		service: service,
	}
}

// Create godoc
// @Summary Login
// @description Login Customer
// @tags Customer
// @Accept json
// @Produce json
// @Param Customer body customermitra.AuthLogin true "Customer"
// @Success 200	{object} response.Result
// @Failure 400 {object} response.Error
// @Router /customer [post]
func (Controller *Controller) Login(c echo.Context) error {
	var req customermitraBussiness.AuthLogin
	var err error
	c.Bind(&req)
	result, err := Controller.service.LoginCustomer(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success login",
		"result":   result,
	})
}

func (Controller *Controller) FindCustomersByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result, err := Controller.service.FindCustomersByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success get customer by id",
		"result":   result,
	})
}

// Create godoc
// @Summary Register
// @description Register customer
// @tags Customer
// @Accept json
// @Produce json
// @Param Registercustomer body customermitra.RegisterCustomer true "Register"
// @Success 200	{object} response.Result
// @Failure 400 {object} response.Error
// @Router /customer/register [post]
func (Controller *Controller) Register(c echo.Context) error {
	var req customermitraBussiness.RegisterCustomer
	c.Bind(&req)
	result, err := Controller.service.CreateCustomer(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success register",
		"result":   result,
	})
}

// Create godoc
// @Summary Updatecustomer
// @description Updatecustomer
// @tags Customer
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param Registercustomer body customermitra.UpdateCustomer true "customer"
// @Success 200 {object} customermitra.UpdateCustomer
// @Router /customer [put]
func (Controller *Controller) UpdateCustomer(c echo.Context) error {
	var req customermitraBussiness.UpdateCustomer
	c.Bind(&req)
	result, err := Controller.service.UpdateCustomer(&req)
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
// @Summary History Customer
// @description History Customer
// @tags Customer
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param limit query int false "pagination"
// @Param page query int false "pagination"
// @Param id query int true "id customer"
// @Success 200	{object} response.Result
// @Failure 400 {object} response.Error
// @Router /history [GET]
func (Controller *Controller) HistoryCustomer(c echo.Context) error {
	idcustomer, _ := strconv.Atoi(c.QueryParam("id"))
	pagination := utils.GeneratePagination(c.Request().URL.Query())
	result, err := Controller.service.HistoryCustomer(idcustomer, pagination)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":       200,
		"messages":   "success get history customer",
		"pagination": pagination,
		"result":     result,
	})
}

// Create godoc
// @Summary Detail history customer
// @description Detail history customer
// @tags Customer
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param idtransaction path string true "id transaction"
// @Success 200	{object} response.Result
// @Failure 400 {object} response.Error
// @Router /dethistory/{idtransaction} [GET]
func (Controller *Controller) DetailHistoryCustomer(c echo.Context) error {
	idcustomer := c.Param("idtransaction")
	result, err := Controller.service.DetailHistoryCustomer(idcustomer)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success get history customer",
		"result":   result,
	})
}

// Create godoc
// @Summary Redeem Emoney
// @description Redeem Emoney customer
// @tags Redeem
// @Accept json
// @Produce json
// @Param InputDataCashout body customermitra.InputTransactionBankEmoney true "inputdataemoney"
// @Success 200	{object} response.Result
// @Failure 400 {object} response.Error
// @Router /emoney [post]
func (Controller *Controller) OrderEmoney(c echo.Context) error {
	emoney := customermitraBussiness.InputTransactionBankEmoney{}
	c.Bind(&emoney)
	_, err := Controller.service.ToOrderEmoney(&emoney)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success order emoney",
	})
}

// Create godoc
// @Summary Redeem Cashout
// @description Redeem Cashout customer
// @tags Redeem
// @Accept json
// @Produce json
// @Param InputDataCashout body customermitra.InputTransactionBankEmoney true "inputdataemoney"
// @Success 200	{object} response.Result
// @Failure 400 {object} response.Error
// @Router /cashout [post]
func (Controller *Controller) OrderCashout(c echo.Context) error {
	req := customermitraBussiness.InputTransactionBankEmoney{}
	c.Bind(&req)
	_, err := Controller.service.RedeemBank(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success order cashout",
	})
}

// Create godoc
// @Summary Redeem Pulsa
// @description Redeem Pulsa customer
// @tags Redeem
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param InputDataPulsa body customermitra.RedeemPulsaData true "Input data pulsa"
// @Success 200	{object} response.Result
// @Failure 400 {object} response.Error
// @Router /pulsa [post]
func (Controller *Controller) OrderPulsa(c echo.Context) error {
	var req customermitraBussiness.RedeemPulsaData
	c.Bind(&req)
	err := Controller.service.RedeemPulsa(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success redeem pulsa",
	})
}

// Create godoc
// @Summary Redeem PaketData
// @description Redeem PaketData customer
// @tags Redeem
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param InputDataPaketData body customermitra.RedeemPulsaData true "Input data paket data"
// @Success 200	{object} response.Result
// @Failure 400 {object} response.Error
// @Router /paketdata [post]
func (Controller *Controller) OrderPaketData(c echo.Context) error {
	var req customermitraBussiness.RedeemPulsaData
	c.Bind(&req)
	err := Controller.service.RedeemPaketData(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success redeem paket data",
	})
}

func (Controller *Controller) CallbackXendit(c echo.Context) error {
	req := c.Request()
	decoder := json.NewDecoder(req.Body)
	disbursermentData := customermitraBussiness.Disbursement{}
	err := decoder.Decode(&disbursermentData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	defer req.Body.Close()
	disbursement, _ := json.Marshal(disbursermentData)
	var resbank customermitraBussiness.Disbursement
	json.Unmarshal(disbursement, &resbank)
	responseWriter := c.Response().Writer
	responseWriter.Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	responseWriter.WriteHeader(200)
	databank, err := Controller.service.GetCallback(&resbank)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success redeem emoney",
		"result":   databank,
	})
}

// Create godoc
// @Summary Register Store
// @description Register Store for Admin
// @tags Admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RegisterStore body customermitra.RegisterStore true "Register Store"
// @Success 200	{object} response.Result
// @Failure 400 {object} response.Error
// @Router /store [post]
func (Controller *Controller) RegisterStore(c echo.Context) error {
	var req customermitraBussiness.RegisterStore
	c.Bind(&req)
	result, err := Controller.service.CreateStore(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success create store",
		"result":   result,
	})
}

// Create godoc
// @Summary Login Store
// @description Register Store for Admin
// @tags Store
// @Accept json
// @Produce json
// @Param LoginStore body customermitra.AuthStore true "LoginStore"
// @Success 200	{object} response.Result
// @Failure 400 {object} response.Error
// @Router /store/login [post]
func (Controller *Controller) LoginStore(c echo.Context) error {
	var req customermitraBussiness.AuthStore
	c.Bind(&req)
	result, err := Controller.service.LoginStore(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success login store",
		"result":   result,
	})
}

// Create godoc
// @Summary Input Poin Store
// @description Input Poin Customer for Store
// @tags Store
// @Accept json
// @Produce json
// @Param InputPoinStore body customermitra.InputPoin true "InputPoinStore"
// @Success 200	{object} response.Result
// @Failure 400 {object} response.Error
// @Router /store/poin [post]
func (Controller *Controller) InputPoinStore(c echo.Context) error {
	var req customermitraBussiness.InputPoin
	c.Bind(&req)
	result, err := Controller.service.InputPoin(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success input poin",
		"Add poin": result,
	})
}
