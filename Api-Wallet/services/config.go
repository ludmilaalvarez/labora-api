package services

import "Api-Wallet/db"

var LogHandler LogService
var WalletHandler WalletService

func Init() {
	WalletHandler = WalletService{
		DbHandlers: &PostgresWallet{Db: db.Db},
	}
	LogHandler = LogService{
		DbHandlers: &PostgresLog{Db: db.Db},
	}
}
