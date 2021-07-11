package common

type ResultStopWatch struct {
	Hour                   int
	Minute                 int
	Second                 int
	Millisecond            int
	TotalTimeInSecond      int
	TotalTimeInMinute      int
	TotalTimeInHour        int
	TotalTimeInMilliSecond int64
}

type ResultTimeAndMileage struct {
	Time    string
	Mileage string
}
