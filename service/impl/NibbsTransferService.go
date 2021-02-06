package impl

import "ajocard/status-indicator/dto"

type NibbsTransferService struct {}

func (nibbs NibbsTransferService) CheckTransferServiceStatus() (dto.IndicatorDto, error) {
	return dto.IndicatorDto{"ACTIVE", 1.5}, nil
}

