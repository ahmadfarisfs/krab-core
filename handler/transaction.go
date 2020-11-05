package handler

import "github.com/labstack/echo/v4"

type ListTransactionRequest struct {
}

//ViewTransactionDetails see details transaction by trx_id
func (h *Handler) ViewTransactionDetails(c echo.Context) error {
	return nil
}

//CreateTransaction create new transaction between accounts
func (h *Handler) CreateTransaction(c echo.Context) error {
	return nil
}

//ListTransaction create new transaction between accounts
func (h *Handler) ListTransaction(c echo.Context) error {
	return nil
}
