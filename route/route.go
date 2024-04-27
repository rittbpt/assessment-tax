package route

import (
	"github.com/labstack/echo/v4"
	_"github.com/labstack/echo/v4/middleware"
	"github.com/rittbpt/assessment-tax/controller"
	// "os"
)


func TaxRoutes(e *echo.Echo , c *controller.TaxController) {
	taxRouter := e.Group("/tax")
	// taxRouter.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
	// 	if username == os.Getenv("ADMIN_USERNAME") && password == os.Getenv("ADMIN_PASSWORD") {
	// 		return true, nil
	// 	}
	// 	return false, nil
	// }))

	taxRouter.POST("/calculations", c.Cal)
}
