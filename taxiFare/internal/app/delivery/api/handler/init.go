package handler

import (
	"taxiFare/internal/app/delivery/api"
	"taxiFare/internal/app/usecase/taxifare"
)

type apiDelivery struct {
	taxiFareUseCase taxifare.UseCase
}

func New(taxiFareUseCase taxifare.UseCase) api.Delivery {
	return &apiDelivery{taxiFareUseCase: taxiFareUseCase}
}
