package data

import (
	"movie-api/features/account"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Username string
	Password string
	Email    string
}

type Watchlist struct {
	gorm.Model
	AccountID int
	MovieID   int
}

func FromAccountCore(Data account.AccountCore) Account {
	return Account{
		Username: Data.Username,
		Password: Data.Password,
		Email:    Data.Email,
	}
}

func FromWatchlistCore(Data account.WatchlistCore) Watchlist {
	return Watchlist{
		AccountID: Data.AccountID,
		MovieID:   Data.MoviesID,
	}
}
