package handler

import (
	"taxiFare/internal/app/repository/taxiFare"
	"taxiFare/internal/app/usecase/taxifare"
)

type calculatePriceUseCase struct {
	taxiFareRedisRepo taxiFare.Repository
	detailFareRepo    taxiFare.Repository
}

func New(taxiFareRepo taxiFare.Repository, detailFareRepo taxiFare.Repository) taxifare.UseCase {
	return &calculatePriceUseCase{taxiFareRedisRepo: taxiFareRepo, detailFareRepo: detailFareRepo}
}
