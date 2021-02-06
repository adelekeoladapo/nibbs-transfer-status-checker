package controller

import (
	"ajocard/status-indicator/service"
	"ajocard/status-indicator/service/impl"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

var transferService service.TransferService = impl.NibbsTransferService{}

func CheckTransferStatus(c echo.Context) error {
	response, err := transferService.CheckTransferServiceStatus()
	if err != nil {
		fmt.Println("An error occurred. ", err.Error())
		return c.String(http.StatusOK, err.Error())
	}
	return c.JSON(http.StatusOK, response)
}
