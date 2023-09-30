package entity

import "github.com/dedicio/sisgares-transactions-service/pkg/utils"

const (
	OrderStatusPending = "pending"
	OrderStatusPaid    = "paid"
	OrderStatusCancel  = "cancel"
)

const (
	PaymentMethodCreditCard = "credit_card"
	PaymentMethodDebitCard  = "debit_card"
	PaymentMethodPix        = "pix"
)

type OrderItem struct {
	ProductID string
	Quantity  int64
	Price     float64
}

type Order struct {
	ID            string
	Items         []OrderItem
	Discount      float64
	Status        string
	PaymentMethod string
}

type OrderRepository interface {
	Store(order Order) error
	FindAll() ([]Order, error)
	FindByID(id string) (Order, error)
	UpdateStatus(id string, status string) error
}

func (o *Order) TotalPrice() float64 {
	var totalPrice float64

	for _, item := range o.Items {
		totalPrice += item.Price * float64(item.Quantity)
	}

	return totalPrice - o.Discount
}

func NewOrder(
	items []OrderItem,
	discount float64,
	paymentMethod string,
) *Order {
	id := utils.NewUUID()
	return &Order{
		ID:            id,
		Items:         items,
		Discount:      discount,
		Status:        OrderStatusPending,
		PaymentMethod: paymentMethod,
	}
}
