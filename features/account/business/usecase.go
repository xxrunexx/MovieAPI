package business

import "movie-api/features/account"

type AccountBusiness struct {
	accountData account.Data
}

func NewBusinessAccount(accountData account.Data) account.Business {
	return &AccountBusiness{accountData}
}

func (accBusiness *AccountBusiness) CreateAccount(accData account.AccountCore) error {
	if err := accBusiness.accountData.InsertAccount(accData); err != nil {
		return err
	}
	return nil
}

func (wlBusiness *AccountBusiness) CreateWatchlist(wlData account.WatchlistCore) error {
	if err := wlBusiness.accountData.InsertWatchlist(wlData); err != nil {
		return err
	}
	return nil
}
