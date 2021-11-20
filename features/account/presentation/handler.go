package presentation

import (
	// Import Echo
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	// Import files
	"movie-api/features/account"
	"movie-api/features/account/presentation/request"
)

type AccountHandler struct {
	accountBusiness account.Business
}

func NewHandlerAccount(accountBusiness account.Business) *AccountHandler {
	return &AccountHandler{accountBusiness}
}

func (accHandler *AccountHandler) CreateAccount(c echo.Context) error {
	newAccount := request.ReqAccount{}

	if err := c.Bind(&newAccount); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Bad Request",
		})
	}

	if err := accHandler.accountBusiness.CreateAccount(request.ToAccountCore(newAccount)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Bad Request",
		})
	}

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "Success",
		"data":    newAccount,
	})
}

func (wlHandler *AccountHandler) CreateWatchlist(c echo.Context) error {
	newWatchlist := request.ReqWatchlist{}
	id, _ := strconv.Atoi(c.Param("id"))
	newWatchlist.AccountID = id
	if err := c.Bind(&newWatchlist); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Bad Request",
		})
	}

	if err := wlHandler.accountBusiness.CreateWatchlist(request.ToWatchlistCore(newWatchlist)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Bad Request",
		})
	}

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "Success",
		"data":    newWatchlist,
	})
}
