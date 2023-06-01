package models

type DBHandlerWallet interface {
	CrearWallet(Datos *Datos_Solicitados) (int, error)
	StatusWallet(dni string) (Wallets, error)
}

type DBHandlerLog interface {
	CrearSolicitud(Datos *Datos_Solicitados) (string, error)
}
