package factory

import (
	"movie-api/features/account/business"
	"movie-api/features/account/data"
	"movie-api/features/account/presentation"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	db, err := gorm.Open(mysql.Open("root:admin@/moviein?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db
	DB.AutoMigrate(&data.Account{}, &data.Watchlist{})

}

type presenter struct {
	AccountPresentation presentation.AccountHandler
}

func Init() presenter {
	accountData := data.NewMySqlAccount(DB)
	accountBusiness := business.NewBusinessAccount(accountData)
	return presenter{
		AccountPresentation: *presentation.NewHandlerAccount(accountBusiness),
	}
}
