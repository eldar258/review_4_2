package storage

import (
	"github.com/jmoiron/sqlx"
	"studentgit.kata.academy/eldar/review_4_2/model"
)

const list = `SELECT * FROM cities`
const selectById = `SELECT * FROM cities WHERE id=$1`
const insert = `INSERT INTO cities(name, state) VALUES ($1, $2) RETURNING id`
const update = `UPDATE cities SET name=$2, state=$3 WHERE id=$1`
const delete = `DELETE FROM cities WHERE id=$1`

type DAO interface {
	List() ([]*model.Cities, error)
	SelectById(id int) (*model.Cities, error)
	Insert(name, state string) (int, error)
	Update(id int, name, state string) error
	Delete(id int) error
}

type DAOImpl struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) DAO {
	return &DAOImpl{db: db}
}

func (d *DAOImpl) List() ([]*model.Cities, error) {
	var res []*model.Cities
	err := d.db.Select(&res, list)
	return res, err
}

func (d *DAOImpl) SelectById(id int) (*model.Cities, error) {
	var res model.Cities
	err := d.db.Get(res, selectById, id)
	return &res, err
}

func (d *DAOImpl) Insert(name, state string) (int, error) {
	var res int
	err := d.db.Get(&res, insert, name, state)
	return res, err
}

func (d *DAOImpl) Update(id int, name, state string) error {
	_, err := d.db.Exec(update, name, state)
	return err
}

func (d *DAOImpl) Delete(id int) error {
	_, err := d.db.Exec(delete, id)
	return err
}
