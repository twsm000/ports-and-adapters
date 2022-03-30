package ports

type APIPort interface {
	GetAdd(a, b int32) (int32, error)
	GetSub(a, b int32) (int32, error)
	GetMulti(a, b int32) (int32, error)
	GetDiv(a, b int32) (int32, error)
}
