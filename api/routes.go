package api

import (
	"api-redeem-point/api/admin"
	"api-redeem-point/api/customermitra"
	"api-redeem-point/api/middleware"

	//auth "api-redeem-point/api/middleware"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	AdminControlller        *admin.Controller
	CustomerMitraController *customermitra.Controller
}

func RegistrationPath(e *echo.Echo, controller Controller) {
	c := e.Group("/v1")
	c.POST("/customer/register", controller.CustomerMitraController.Register)
	c.POST("/customer", controller.CustomerMitraController.Login)
	c.PUT("/customer", controller.CustomerMitraController.UpdateCustomer)
	c.GET("/customer/:id", controller.CustomerMitraController.FindCustomersByID)
	c.GET("/history", controller.CustomerMitraController.HistoryCustomer, middleware.CustomerSetupAuthenticationJWT())
	c.GET("/dethistory/:idtransaction", controller.CustomerMitraController.DetailHistoryCustomer, middleware.CustomerSetupAuthenticationJWT())
	c.POST("/pulsa", controller.CustomerMitraController.OrderPulsa, middleware.CustomerSetupAuthenticationJWT())
	c.POST("/paketdata", controller.CustomerMitraController.OrderPaketData, middleware.CustomerSetupAuthenticationJWT())
	c.POST("/cashout", controller.CustomerMitraController.OrderCashout, middleware.CustomerSetupAuthenticationJWT())
	c.POST("/emoney", controller.CustomerMitraController.OrderEmoney, middleware.CustomerSetupAuthenticationJWT())
	c.POST("/callback", controller.CustomerMitraController.CallbackXendit)
	//admin
	g := c.Group("/admin")
	g.POST("/login", controller.AdminControlller.LoginAdmin)
	g.POST("", controller.AdminControlller.CreateAdmin)
	g.GET("", controller.AdminControlller.Dashboard)
	g.GET("/:id", controller.AdminControlller.FindAdminByID)
	g.GET("/transaction/pending", controller.AdminControlller.TransactionPending)
	g.POST("/approve/:idtransaction", controller.AdminControlller.ApproveTransaction)
	g.GET("/history", controller.AdminControlller.FindHistoryCustomers)
	g.GET("/transaction", controller.AdminControlller.TransactionByDate)
	g.GET("/customer", controller.AdminControlller.FindCustomers)
	g.PUT("/customer", controller.AdminControlller.UpdateCustomer)
	g.DELETE("/customer", controller.AdminControlller.DeleteCustomer)
	g.PUT("/customer/point", controller.AdminControlller.UpdateCustomerPoint)
	g.GET("/stock", controller.AdminControlller.StockProduct)
	g.PUT("/stock", controller.AdminControlller.UpdateStock)
	g.GET("/historystore", controller.AdminControlller.HistoryStore)
	g.DELETE("/store", controller.AdminControlller.DeleteStore)
	g.GET("/store", controller.AdminControlller.GetStore)
	g.PUT("/store", controller.AdminControlller.UpdateStore)
	s := c.Group("/store")
	s.POST("", controller.CustomerMitraController.RegisterStore)
	s.POST("/login", controller.CustomerMitraController.LoginStore)
	s.POST("/poin", controller.CustomerMitraController.InputPoinStore, middleware.StoreSetupAuthenticationJWT())
}
