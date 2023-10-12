package repositories

type Repository interface {
	Ping() error
	Transaction(txFunc func(Repository) error) error
}
