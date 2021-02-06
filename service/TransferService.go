package service

import "ajocard/status-indicator/dto"

type TransferService interface {

	CheckTransferServiceStatus() (dto.IndicatorDto, error)

}