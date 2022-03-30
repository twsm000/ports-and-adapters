package ports

type ArithmeticPort interface {
	Add(a, b int32) (int32, error)
	Sub(a, b int32) (int32, error)
	Multi(a, b int32) (int32, error)
	Div(a, b int32) (int32, error)
}
