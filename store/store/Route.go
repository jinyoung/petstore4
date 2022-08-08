package store

import (
	"github.com/labstack/echo"
)

func RouteInit() *echo.Echo {
	e := echo.New()
	storeOrder := &StoreOrder{}
	e.GET("/storeOrders", storeOrder.Get)
	e.GET("/storeOrders/:id", storeOrder.GetbyId)
	e.POST("/storeOrders", storeOrder.Persist)
	e.PUT("/storeOrders/:id", storeOrder.Put)
	e.DELETE("/storeOrders/:id", storeOrder.Remove)
	e.PUT("/finishCook", func(c echo.Context) error {
		return c.String(http.StatusOK, "finishCook")
	})
	e.PUT("/accept", func(c echo.Context) error {
		return c.String(http.StatusOK, "accept")
	})
	e.PUT("/reject", func(c echo.Context) error {
		return c.String(http.StatusOK, "reject")
	})
	e.PUT("/startCook", func(c echo.Context) error {
		return c.String(http.StatusOK, "startCook")
	})
	menu := &Menu{}
	e.GET("/menus", menu.Get)
	e.GET("/menus/:id", menu.GetbyId)
	e.POST("/menus", menu.Persist)
	e.PUT("/menus/:id", menu.Put)
	e.DELETE("/menus/:id", menu.Remove)
	e.PUT("/remove", func(c echo.Context) error {
		return c.String(http.StatusOK, "메뉴삭제")
	})
	e.PUT("/set-price", func(c echo.Context) error {
		return c.String(http.StatusOK, "changePrice")
	})
	topFood := &TopFood{}
	e.GET("/topFoods", topFood.Get)
	e.GET("/topFoods/:id", topFood.GetbyId)
	e.POST("/topFoods", topFood.Persist)
	e.PUT("/topFoods/:id", topFood.Put)
	e.DELETE("/topFoods/:id", topFood.Remove)
	return e
}
