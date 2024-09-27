package usecase

import (
	"context"
	"github.com/devfullcycle/20-CleanArch/internal/entity"
)

type ListOrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (lo *ListOrderUseCase) Execute(ctx context.Context) ([]ListOrderOutputDTO, error) {
	orders, err := lo.OrderRepository.Find()
	if err != nil {
		return nil, err
	}

	var output []ListOrderOutputDTO
	for _, order := range orders {
		dto := ListOrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
		output = append(output, dto)
	}
	return output, nil
}
