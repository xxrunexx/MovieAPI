package account

type AccountCore struct {
	ID         int
	Username   string
	Password   string
	Email      string
	Watchlists []WatchlistCore
}

type WatchlistCore struct {
	ID        int
	AccountID int
	MoviesID  int
}

// Untuk layer data
type Data interface {
	InsertAccount(accData AccountCore) error
	InsertWatchlist(wlData WatchlistCore) error
}

// Untuk layer business
type Business interface {
	CreateAccount(accData AccountCore) error
	CreateWatchlist(wlData WatchlistCore) error
}
