package repository

import "errors"

type ProductRepository struct {
	products []Product
}

func (r *ProductRepository) GetAll() ([]Product, error) {
	return r.products, nil
}

func (r *ProductRepository) GetById(id int) (Product, error) {
	for _, product := range r.products {
		if product.ID == id {
			return product, nil
		}
	}
	return Product{}, errors.New("product not found")
}

func (r *ProductRepository) Create(product Product) (Product, error) {
	r.products = append(r.products, product)
	return product, nil
}

func (r *ProductRepository) Update(id int, product Product) (Product, error) {
	for i, p := range r.products {
		if p.ID == id {
			r.products[i] = product
			return product, nil
		}
	}
	return Product{}, errors.New("product not found")
}

func (r *ProductRepository) Delete(id int) error {
	for i, product := range r.products {
		if product.ID == id {
			r.products = append(r.products[:i], r.products[i+1:]...)
			return nil
		}
	}
	return errors.New("product not found")
}
