package handler

import "github.com/labstack/echo/v4"

type createAccountRequest struct {
	Name string `json:"name" validate:"required"`
}

func (ca *createAccountRequest) bind(c echo.Context) error {
	if err := c.Bind(ca); err != nil {
		return err
	}
	if err := c.Validate(ca); err != nil {
		return err
	}
	return nil
}
