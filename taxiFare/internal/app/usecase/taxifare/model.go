package taxifare

type Param struct {
	TimeAndMileage string
	Name           string
}

type Response struct {
	TotalFare            int32
	TotalDistanceInMeter float64
	Name                 string
}
