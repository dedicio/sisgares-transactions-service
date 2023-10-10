package entity

import "github.com/dedicio/sisgares-transactions-service/pkg/utils"

const (
	OrderStatusOpen    = "open"
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
	ID        string
	OrderID   string
	ProductID string
	Quantity  int64
	Price     float64
}

type Order struct {
	ID            string
	Items         []*OrderItem
	Discount      float64
	Status        string
	PaymentMethod string
	CompanyId     string
	CreatedAt     string
	UpdatedAt     string
}

type OrderRepository interface {
	Create(order *Order) error
	FindAll(companyID string) ([]*Order, error)
	FindByID(id string) (*Order, error)
	UpdateStatus(id string, status string) error
	FindAllOrderItemsByOrderId(orderId string) ([]*OrderItem, error)
}

func (o *Order) TotalPrice() float64 {
	var totalPrice float64

	for _, item := range o.Items {
		totalPrice += item.Price * float64(item.Quantity)
	}

	return totalPrice - o.Discount
}

func NewOrder(
	items []*OrderItem,
	discount float64,
	paymentMethod string,
	companyId string,
) *Order {
	id := utils.NewUUID()
	return &Order{
		ID:            id,
		Items:         items,
		Discount:      discount,
		Status:        OrderStatusPending,
		PaymentMethod: paymentMethod,
		CompanyId:     companyId,
		CreatedAt:     utils.Now(),
		UpdatedAt:     utils.Now(),
	}
}
