package storage

import (
	"github.com/jmoiron/sqlx"
	"studentgit.kata.academy/eldar/review_4_2/model"
)

const list = `SELECT * FROM cities`
const selectById = `SELECT * FROM cities WHERE id=?`
const insert = `INSERT INTO cities(name, state) VALUES (?, ?)`
const update = `UPDATE cities SET name=?, state=? WHERE id=?`
const delete = `DELETE FROM cities WHERE id=?`

type DAO interface {
	List() ([]*model.City, error)
	SelectById(id int) (*model.City, error)
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

func (d *DAOImpl) List() ([]*model.City, error) {
	var res []*model.City
	err := d.db.Select(&res, list)
	return res, err
}

func (d *DAOImpl) SelectById(id int) (*model.City, error) {
	var res model.City
	err := d.db.Get(res, selectById, id)
	return &res, err
}

func (d *DAOImpl) Insert(name, state string) (int, error) {
	res, err := d.db.Exec(insert, name, state)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return int(id), err
}

func (d *DAOImpl) Update(id int, name, state string) error {
	_, err := d.db.Exec(update, name, state, id)
	return err
}

func (d *DAOImpl) Delete(id int) error {
	_, err := d.db.Exec(delete, id)
	return err
}
