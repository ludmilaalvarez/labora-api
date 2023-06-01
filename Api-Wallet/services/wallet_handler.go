package services

import (
	"Api-Wallet/models"
)

type WalletService struct { //objeto para poder utilizar las operaciones CRUD de solicitud
	DbHandlers models.DBHandlerWallet
}

func (s *WalletService) CrearWallet(Datos *models.Datos_Solicitados) (int, error) {
	return s.DbHandlers.CrearWallet(Datos)
}

func (s *WalletService) StatusWallet(Dni string) (models.Wallets, error) {
	return s.DbHandlers.StatusWallet(Dni)
}
