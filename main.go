package main

import (
	"movie-api/factory"

	// "github.com/labstack/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

// Use for Request later
type Account struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// Use for Response Later
type ResponseAccount struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func main() {
	factory.InitDB()
	// Initiate Echo
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	presenter := factory.Init()

	// Routing
	e.POST("/account", presenter.AccountPresentation.CreateAccount)
	e.POST("/account/:id/watchlist", presenter.AccountPresentation.CreateWatchlist)
	// e.PUT("/account/:id", UpdateAccount)
	// e.GET("/account", GetAllAccount)
	// e.GET("/account/:id", GetAccountByID)

	// Starting the server
	e.Start(":8585")
}

// func GetAllAccount(c echo.Context) error {
// 	var accounts []Account

// 	if err := DB.Find(&accounts).Error; err != nil {
// 		return c.JSON(http.StatusBadRequest, err)
// 	}
// 	return c.JSON(http.StatusAccepted, map[string]interface{}{
// 		"message": "Success",
// 		"data":    accounts,
// 	})
// }

// func UpdateAccount(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))

// 	reqAccount := Account{}
// 	c.Bind(&reqAccount)

// var dataAccount Account
// reqAccount.ID = uint(id)

// DB.Where("id = ?", id).First(&dataAccount)

// if reqAccount.Username != "" {
// 	dataAccount.Username = reqAccount.Username
// }

// if err := DB.Where("id = ?", id).Updates(&reqAccount).Error; err != nil {
// 	return c.JSON(http.StatusBadRequest, err)
// }
// SAVE FOR LATER
// var dataAccount Account
// var dataAccount ResponseAccount // COBA COBA
// DB.Raw("SELECT * FROM accounts WHERE id = ?", id).Find(&dataAccount)
// DB.Where("id = ?", id).Find(&dataAccount)
// reqAccount.ID = uint(id)

// 	return c.JSON(http.StatusAccepted, map[string]interface{}{
// 		"message": "Success",
// 		"data":    dataAccount,
// 	})
// }

// func GetAccountByID(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	reqAccount := Account{}
// 	c.Bind(&reqAccount)

// 	if err := DB.Where("id = ?", id).First(&reqAccount); err != nil {
// 		return c.JSON(http.StatusNotFound, err)
// 	}

// 	var dataAccount ResponseAccount
// 	DB.Where("id = ?", id)
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "Success",
// 		"data":    dataAccount,
// 	})
// if id <= len(data) && id > 0 {
// 	return c.JSON(http.StatusOK, data[id-1])
// } else {
// 	return c.JSON(http.StatusOK, data)
// }
//}
