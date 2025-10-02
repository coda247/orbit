package GrpcOrder

import (
	orbit "github.com/coda247/orbit"
	"github.com/google/uuid"
)

func (r *OrderKey) ToOrderKey() *orbit.OrderKey {
	uuid, _ := uuid.FromBytes(r.Uuid)

	return &orbit.OrderKey{
		ID:        r.Id,
		UUID:      uuid,
		Symbol:    r.Symbol.ToSymbol(),
		Side:      orbit.OrderSide(r.Side),
		Price:     r.Price.ToDecimal(),
		StopPrice: r.StopPrice.ToDecimal(),
		Fake:      r.Fake,
		CreatedAt: r.CreatedAt.AsTime(),
	}
}

func (r *Order) ToOrder() *orbit.Order {
	uuid, _ := uuid.FromBytes(r.Uuid)

	return &orbit.Order{
		ID:             r.Id,
		UUID:           uuid,
		Symbol:         r.Symbol.ToSymbol(),
		MemberID:       r.MemberId,
		Side:           orbit.OrderSide(r.Side),
		Type:           orbit.OrderType(r.Type),
		Price:          r.Price.ToDecimal(),
		StopPrice:      r.StopPrice.ToDecimal(),
		Quantity:       r.Quantity.ToDecimal(),
		FilledQuantity: r.FilledQuantity.ToDecimal(),
		Cancelled:      r.Cancelled,
		Fake:           r.Fake,
		CreatedAt:      r.CreatedAt.AsTime(),
	}
}
