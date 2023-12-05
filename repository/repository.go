package repository

import (
	"studentgit.kata.academy/eldar/review_4_2/model"
	"studentgit.kata.academy/eldar/review_4_2/storage"
)

type Repositorer interface {
	Create(name, state string) (int, error)
	Delete(id int) error
	Update(id int, name, state string) error
	List() ([]*model.Cities, error)
}

type Repository struct {
	db storage.DAO
}

func New(db storage.DAO) Repositorer {
	return &Repository{db: db}
}

func (r *Repository) Create(name, state string) (int, error) {
	return r.db.Insert(name, state)
}

func (r *Repository) Delete(id int) error {
	return r.db.Delete(id)
}

func (r *Repository) Update(id int, name, state string) error {
	return r.db.Update(id, name, state)
}

func (r *Repository) List() ([]*model.Cities, error) {
	return r.db.List()
}
