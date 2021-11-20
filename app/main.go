package main

import (
	"net/http"
	"strconv"

	// "github.com/labstack/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
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

var DB *gorm.DB

func initDB() {
	var err error
	db, err := gorm.Open(mysql.Open("root:admin@/moviein?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db
	DB.AutoMigrate(&Account{})
}

func main() {
	// Initiate DB
	initDB()

	// Initiate Echo
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	// Routing
	e.POST("/account", AddAccount)
	e.GET("/account", GetAllAccount)
	e.PUT("/account/:id", UpdateAccount)

	// Starting the server
	e.Start(":8585")
}

func AddAccount(c echo.Context) error {
	newAccount := Account{}
	// c.Bind(&newAccount)

	if err := c.Bind(&newAccount); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Bad Request",
		})
	}

	if err := DB.Create(&newAccount).Error; err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	// data = append(data, newAccount)
	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "hope all feeling well",
		"data":    newAccount,
	})
}

func GetAllAccount(c echo.Context) error {
	var accounts []Account

	if err := DB.Find(&accounts).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "hope all feeling well",
		"data":    accounts,
	})
}

func UpdateAccount(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	reqAccount := Account{}
	c.Bind(&reqAccount)

	// var dataAccount Account
	// reqAccount.ID = uint(id)

	// DB.Where("id = ?", id).First(&dataAccount)

	// if reqAccount.Username != "" {
	// 	dataAccount.Username = reqAccount.Username
	// }

	if err := DB.Where("id = ?", id).Updates(&reqAccount).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// SAVE FOR LATER
	// var dataAccount Account
	var dataAccount ResponseAccount // COBA COBA
	DB.Raw("SELECT * FROM accounts WHERE id = ?", id).Find(&dataAccount)
	// reqAccount.ID = uint(id)

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "hope all feeling well",
		"data":    dataAccount,
	})
}

// func getAccount(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))

// 	if id <= len(data) && id > 0 {
// 		return c.JSON(http.StatusOK, data[id-1])
// 	} else {
// 		return c.JSON(http.StatusOK, data)
// 	}
// }
