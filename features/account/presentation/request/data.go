package request

import (
	"movie-api/features/account"
)

type ReqAccount struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type ReqWatchlist struct {
	AccountID int `json:"account_id"`
	MovieID   int `json:"movies_id"`
}

func ToAccountCore(reqData ReqAccount) account.AccountCore {
	return account.AccountCore{
		Username: reqData.Username,
		Password: reqData.Password,
		Email:    reqData.Email,
	}
}

func ToWatchlistCore(reqData ReqWatchlist) account.WatchlistCore {
	return account.WatchlistCore{
		AccountID: reqData.AccountID,
		MoviesID:  reqData.MovieID,
	}
}
