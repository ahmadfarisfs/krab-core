package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v1 *echo.Group) {
	accountGroup := v1.Group("/account")
	accountGroup.POST("/", h.RegisterAccount)
	accountGroup.GET("/:id", h.ViewAccountSummary)
	accountGroup.GET("/list", h.ListAccount)
	accountGroup.GET("/mutation", h.ViewMutation)

	trxGroup := v1.Group("/transaction")
	trxGroup.POST("/", h.CreateTransaction)
	trxGroup.GET("/:id", h.ViewTransactionDetails)
	trxGroup.GET("/list", h.ListTransaction)
}
