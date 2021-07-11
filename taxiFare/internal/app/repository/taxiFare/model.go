package taxiFare

type Param struct {
	DataTimeAndMileage string
	Name               string
	Mileage            float64
}

type Response struct {
	FareCategory            CategoryFareStatus
	AdditionalDistanceFare  int
	Mileage                 float64
	TotalTimeInMinute       int
}

type ResponseRedis struct {
	TimeAndMileage    string
	Mileage           float64
	TotalTimeInMinute int
}
