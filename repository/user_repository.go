package repository

import "errors"

// Определяем структуры для наших сущностей
type User struct {
	ID   int
	Name string
}

type Product struct {
	ID    int
	Name  string
	Price float64
}

type Order struct {
	ID        int
	ProductID int
	UserID    int
	Quantity  int
}

// Реализация репозитория для User
type UserRepository struct {
	users []User
}

func (r *UserRepository) GetAll() ([]User, error) {
	return r.users, nil
}

func (r *UserRepository) GetById(id int) (User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}
	return User{}, errors.New("user not found")
}

func (r *UserRepository) Create(user User) (User, error) {
	r.users = append(r.users, user)
	return user, nil
}

func (r *UserRepository) Update(id int, user User) (User, error) {
	for i, u := range r.users {
		if u.ID == id {
			r.users[i] = user
			return user, nil
		}
	}
	return User{}, errors.New("user not found")
}

func (r *UserRepository) Delete(id int) error {
	for i, user := range r.users {
		if user.ID == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}
