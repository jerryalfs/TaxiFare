package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/tokopedia/panics"
	"gopkg.in/tokopedia/grace.v1"
	"taxiFare/common/redis"
	"taxiFare/internal/app/delivery/api"
	apiHandler "taxiFare/internal/app/delivery/api/handler"
	"taxiFare/internal/app/repository/taxiFare"
	detailFareRepository "taxiFare/internal/app/repository/taxiFare/faredata"
	redisTaxiFareRepo "taxiFare/internal/app/repository/taxiFare/redis"
	"taxiFare/internal/app/usecase/taxifare"
	handlerTaxiFareUseCase "taxiFare/internal/app/usecase/taxifare/handler"
)

var (
	taxiFareRepo    taxiFare.Repository
	detailFareRepo  taxiFare.Repository
	taxiFareUseCase taxifare.UseCase
	APIDelivery     api.Delivery
)

func init() {
	initInternal()
}

func initInternal() {
	redisTF := redis.NewConnection(
		redis.ConnectionConfig{
			Address:         "localhost:6379",
			MaxActive:       10,
			IdleTimeout:     10,
			MaxIdle:         10,
			MaxConnLifeTime: 10,
		},
	)
	taxiFareRepo = redisTaxiFareRepo.New(redisTF)
	detailFareRepo = detailFareRepository.New()
	taxiFareUseCase = handlerTaxiFareUseCase.New(taxiFareRepo,detailFareRepo)
	APIDelivery = apiHandler.New(taxiFareUseCase)
}

func main() {
	router := httprouter.New()
	router.POST("/taxifare/detail", panics.CaptureHTTPRouterHandler(APIDelivery.GetTaxiFareHandler))
	grace.Serve(":9000", router)
}
