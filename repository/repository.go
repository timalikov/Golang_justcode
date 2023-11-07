package repository

type Repository[T any] interface {
	GetAll() ([]T, error)
	GetById(id int) (T, error)
	Create(entity T) (T, error)
	Update(id int, entity T) (T, error)
	Delete(id int) error
}
