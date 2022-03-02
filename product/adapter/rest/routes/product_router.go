package routes

import (
	"github.com/labstack/echo/v4"
	"product/adapter/rest/controllers/interfaces"
)

func ProductRouter(e *echo.Echo, Icontroller interfaces.IproductController) *echo.Echo {
	e.POST("/product", func(context echo.Context) error { return Icontroller.CreateProduct(context) })

	return e
}
