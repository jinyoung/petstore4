package store

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo"
)

func (self *TopFood) Get(c echo.Context) error {
	repository := TopFoodRepository()
	entities := repository.GetList()
	return c.JSON(http.StatusOK, entities)
}

func (self *TopFood) GetbyId(c echo.Context) error{
	repository := TopFoodRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	self = repository.GetID(id)

	return c.JSON(http.StatusOK, self)
}

func (self *TopFood) Persist(c echo.Context) error{
	repository := TopFoodRepository()
	params := make(map[string] string)
	
	c.Bind(&params)
	ObjectMapping(self, params)

	err := repository.save(self)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}else{
		return c.JSON(http.StatusOK, self)
	}
}

func (self *TopFood) Put(c echo.Context) error{
	repository := TopFoodRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	params := make(map[string] string)
	
	c.Bind(&params)

	err := repository.Update(id, params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	} else {
		entity := repository.GetID(id)
		return c.JSON(http.StatusOK, entity)
	}
}

func (self *TopFood) Remove(c echo.Context) error{
	repository := TopFoodRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	self = repository.GetID(id)

	err := repository.Delete(self)

	return c.JSON(http.StatusOK, err)
}