package entity

type OrderRepositoryInterface interface {
	Save(order *Order) error
	Find() ([]Order, error)
	// GetTotal() (int, error)
}
