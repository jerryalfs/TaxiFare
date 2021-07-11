package taxifare

type UseCase interface {
	GetData(param Param) (res Response, err error)
}