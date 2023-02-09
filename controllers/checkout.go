package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/moody/helpers"
	"github.com/moody/models"
)

func CheckoutGetById(c echo.Context) error {
	m := models.Checkout{}
	res := m.GetById(helpers.SetContext(c), c.Param("id"), c.QueryParams())
	return helpers.Response(c, 200, res)
}

func CheckoutGetPaginated(c echo.Context) error {
	m := models.Checkout{}
	res := m.GetPaginated(helpers.SetContext(c), c.QueryParams())
	return helpers.Response(c, 200, res)
}

func CheckoutUpdateById(c echo.Context) error {
	o := new(models.Checkout)
	if err := c.Bind(o); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	o.ID = helpers.Convert(c.Param("id")).String()
	res := o.UpdateById(helpers.SetContext(c))
	return helpers.Response(c, 200, res)
}

func CheckoutPartialUpdateById(c echo.Context) error {
	return UserUpdateById(c)
}

func CheckoutDeleteById(c echo.Context) error {
	o := new(models.Checkin)
	if err := c.Bind(o); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	o.ID = helpers.Convert(c.Param("id")).String()
	res := o.DeleteById(helpers.SetContext(c))
	return helpers.Response(c, 200, res)
}

func CheckOutCreate(c echo.Context) error {
	o := new(models.Checkout)
	if err := c.Bind(o); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	res := o.Create(helpers.SetContext(c))

	return helpers.Response(c, 201, res)
}
