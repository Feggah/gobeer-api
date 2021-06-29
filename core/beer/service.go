package beer

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

const (
	errInvalidID = "Invalid ID"
)

type UseCase interface {
	GetAll() ([]*Beer, error)
	Get(ID int64) (*Beer, error)
	Store(beer *Beer) error
	Update(beer *Beer) error
	Remove(ID int64) error
}

type Service struct {
	DB *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (svc *Service) GetAll() ([]*Beer, error) {
	var result []*Beer

	rows, err := svc.DB.Query("SELECT id, name, type, style from beer")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var beer Beer
		err = rows.Scan(&beer.ID, &beer.Name, &beer.Type, &beer.Style)
		if err != nil {
			return nil, err
		}
		result = append(result, &beer)
	}
	return result, nil
}

func (svc *Service) Get(ID int64) (*Beer, error) {
	var beer Beer

	stmt, err := svc.DB.Prepare("select id, name, type, style from beer where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(ID).Scan(&beer.ID, &beer.Name, &beer.Type, &beer.Style)
	if err != nil {
		return nil, err
	}
	return &beer, nil
}

func (svc *Service) Store(beer *Beer) error {
	tx, err := svc.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("insert into beer(id, name, type, style) values (?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(beer.ID, beer.Name, beer.Type, beer.Style)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (svc *Service) Update(beer *Beer) error {
	if beer.ID == 0 {
		return errors.New(errInvalidID)
	}

	tx, err := svc.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("update beer set name=?, type=?, style=? where id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(beer.Name, beer.Type, beer.Style, beer.ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (svc *Service) Remove(ID int64) error {
	if ID == 0 {
		return errors.New(errInvalidID)
	}

	tx, err := svc.DB.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("delete from beer where id=?", ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
