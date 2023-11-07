// order_repository.go
package repository

import "errors"

type OrderRepository struct {
	orders []Order
}

func (r *OrderRepository) GetAll() ([]Order, error) {
	return r.orders, nil
}

func (r *OrderRepository) GetById(id int) (Order, error) {
	for _, order := range r.orders {
		if order.ID == id {
			return order, nil
		}
	}
	return Order{}, errors.New("order not found")
}

func (r *OrderRepository) Create(order Order) (Order, error) {
	r.orders = append(r.orders, order)
	return order, nil
}

func (r *OrderRepository) Update(id int, order Order) (Order, error) {
	for i, o := range r.orders {
		if o.ID == id {
			r.orders[i] = order
			return order, nil
		}
	}
	return Order{}, errors.New("order not found")
}

func (r *OrderRepository) Delete(id int) error {
	for i, order := range r.orders {
		if order.ID == id {
			r.orders = append(r.orders[:i], r.orders[i+1:]...)
			return nil
		}
	}
	return errors.New("order not found")
}
