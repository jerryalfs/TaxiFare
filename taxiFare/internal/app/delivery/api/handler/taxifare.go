package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mholt/binding"
	"taxiFare/common/sanitize"
	"taxiFare/internal/app/usecase/taxifare"
)

func (a *apiDelivery) GetTaxiFareHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	success := true
	param, err := mapTaxiFareParam(req)
	if err != nil {
		sanitize.RenderError(w, err, 400)
		return
	}
	data, err := a.taxiFareUseCase.GetData(param)
	if err != nil {
		sanitize.RenderError(w, err, 400)
		return
	}
	response := ResponseTaxiFare{
		TotalFare:            data.TotalFare,
		TotalDistanceInMeter: data.TotalDistanceInMeter,
		Name:                 data.Name,
		Success:              success,
		Status:               "200",
	}
	sanitize.Render(w, response, req.FormValue("callback"))
	return
}

func mapTaxiFareParam(req *http.Request) (result taxifare.Param, err error) {
	var param Param
	tmpErr := binding.Bind(req, &param)
	if err = sanitize.BindError(tmpErr); err != nil {
		return
	}
	result = taxifare.Param{
		TimeAndMileage: param.TimeMileage,
		Name:           param.Name,
	}
	return
}
