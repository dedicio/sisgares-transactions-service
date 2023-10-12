package repository

import (
	"database/sql"

	"github.com/dedicio/sisgares-transactions-service/internal/entity"
)

type OrderRepositoryPostgresql struct {
	db *sql.DB
}

func NewOrderRepositoryPostgresql(db *sql.DB) *OrderRepositoryPostgresql {
	return &OrderRepositoryPostgresql{
		db: db,
	}
}

func (pr *OrderRepositoryPostgresql) FindByID(id string) (*entity.Order, error) {
	var order entity.Order

	sqlOrderStatement := `
		SELECT
			id,
			discount,
			status,
			payment_method,
			created_at,
			updated_at
		FROM orders
		WHERE id = $1
			AND deleted_at IS NULL
	`
	err := pr.db.QueryRow(sqlOrderStatement, id).Scan(
		&order.ID,
		&order.Discount,
		&order.Status,
		&order.PaymentMethod,
	)

	if err != nil {
		return nil, err
	}

	sqlOrderItemStatement := `
		SELECT
			product_id,
			quantity,
			price
		FROM order_items
		WHERE order_id = $1
			AND deleted_at IS NULL
	`
	rows, err := pr.db.Query(sqlOrderItemStatement, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	orderItems, err := pr.FindAllOrderItemsByOrderId(order.ID)
	if err != nil {
		return nil, err
	}
	order.Items = orderItems

	return &order, nil
}

func (pr *OrderRepositoryPostgresql) FindAll(companyID string) ([]*entity.Order, error) {
	sqlOrderStatement := `
		SELECT
			id,
			discount,
			status,
			payment_method
		FROM orders 
		WHERE company_id = $1
			AND deleted_at IS NULL
	`
	rows, err := pr.db.Query(sqlOrderStatement, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*entity.Order
	for rows.Next() {
		var order entity.Order
		err := rows.Scan(
			&order.ID,
			&order.Discount,
			&order.Status,
			&order.PaymentMethod,
		)
		if err != nil {
			return nil, err
		}

		orderItems, err := pr.FindAllOrderItemsByOrderId(order.ID)
		if err != nil {
			return nil, err
		}
		order.Items = orderItems

		orders = append(orders, &order)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (pr *OrderRepositoryPostgresql) Create(order *entity.Order) error {
	sql := `
		INSERT INTO
			orders (
				id,
				discount,
				status,
				payment_method,
				company_id,
				created_at,
				updated_at
			)
		VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			NOW(),
			NOW()
		)
	`
	_, err := pr.db.Exec(
		sql,
		order.ID,
		order.Discount,
		order.Status,
		order.PaymentMethod,
		order.CompanyId,
	)

	if err != nil {
		return err
	}

	for _, item := range order.Items {
		sql := `
			INSERT INTO
				order_items (
					id,
					company_id,
					order_id,
					product_id,
					quantity,
					price,
					created_at,
					updated_at
				)
			VALUES (
				$1,
				$2,
				$3,
				$4,
				$5,
				$6,
				NOW(),
				NOW()
			)
		`
		_, err := pr.db.Exec(
			sql,
			item.ID,
			order.CompanyId,
			order.ID,
			item.ProductID,
			item.Quantity,
			item.Price,
		)

		if err != nil {
			return err
		}

	}

	return nil
}

func (pr *OrderRepositoryPostgresql) UpdateStatus(orderId string, status string) error {
	sql := `
		UPDATE orders
		SET
			status = $1,
			updated_at = NOW()
		WHERE
			id = $2
	`
	_, err := pr.db.Exec(
		sql,
		status,
		orderId,
	)

	if err != nil {
		return err
	}

	return nil
}

func (or *OrderRepositoryPostgresql) FindAllOrderItemsByOrderId(orderId string) ([]*entity.OrderItem, error) {
	sql := `
		SELECT
			product_id,
			quantity,
			price
		FROM order_items
		WHERE order_id = $1
			AND deleted_at IS NULL
	`
	rows, err := or.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orderItems []*entity.OrderItem
	for rows.Next() {
		var orderItem entity.OrderItem
		err := rows.Scan(
			&orderItem.ProductID,
			&orderItem.Quantity,
			&orderItem.Price,
		)
		if err != nil {
			return nil, err
		}
		orderItems = append(orderItems, &orderItem)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return orderItems, nil
}

func (or *OrderRepositoryPostgresql) CreateOrderItem(orderItem *entity.OrderItem) error {
	sql := `
		INSERT INTO
			order_items (
				id,
				order_id,
				product_id,
				quantity,
				price,
				created_at,
				updated_at
			)
		VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			NOW(),
			NOW()
		)
	`
	_, err := or.db.Exec(
		sql,
		orderItem.ID,
		orderItem.OrderID,
		orderItem.ProductID,
		orderItem.Quantity,
		orderItem.Price,
	)

	if err != nil {
		return err
	}

	return nil
}

func (or *OrderRepositoryPostgresql) DeleteOrderItem(orderItemId string) error {
	sql := `
		DELETE FROM order_items
		WHERE id = $1
	`
	_, err := or.db.Exec(
		sql,
		orderItemId,
	)

	if err != nil {
		return err
	}

	return nil
}
