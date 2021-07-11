package taxiFare

// Repository is place to register every function that available to call from taxiFare repository
type Repository interface {
	GetDetailFare(param Param) (result Response, err error)
	GetData(param Param) (result []ResponseRedis, err error)
	StoreData(param Param) (err error)
}
