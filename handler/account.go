package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ahmadfarisfs/krab-core/utils"
	"github.com/labstack/echo/v4"
)

//RegisterAccount freate new accounts
func (h *Handler) RegisterAccount(c echo.Context) error {
	req := &createAccountRequest{}
	if err := req.bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.StandardResponse{Success: false, ErrorMessage: err.Error()})
	}
	ac, err := h.accountStore.CreateAccount(req.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.StandardResponse{Success: false, ErrorMessage: err.Error()})
	}
	return c.JSON(http.StatusOK, utils.StandardResponse{Success: true, Data: ac})
}

//ViewAccountSummary view current summary of an account
func (h *Handler) ViewAccountSummary(c echo.Context) error {
	accountID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.StandardResponse{Success: false, ErrorMessage: err.Error()})
	}
	ac, err := h.accountStore.GetAccountDetails(accountID)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, utils.StandardResponse{Success: false, ErrorMessage: err.Error()})
	}
	return c.JSON(http.StatusOK, utils.StandardResponse{Success: true, Data: ac})
}

//ViewMutation list mutation between dates with filter (list of account id) also with paging
func (h *Handler) ViewMutation(c echo.Context) error {
	return nil
}

//ListAccount list account with paging
func (h *Handler) ListAccount(c echo.Context) error {
	return nil
}
