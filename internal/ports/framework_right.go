package ports

type DBPort interface {
	Close()
	AddHistory(answer int32, operation string) error
}
