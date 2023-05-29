package services

import (
	"Api-Wallet/models"
)

type LogService struct { //objeto para poder utilizar las operaciones CRUD de solicitud
	DbHandlers models.DBHandlerLog
}

func (s *LogService) CrearSolicitud(Datos *models.Datos_Solicitados) string {
	return s.DbHandlers.CrearSolicitud(Datos)
}
