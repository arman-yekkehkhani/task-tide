package chore

import (
	"database/sql"
	"fmt"
	_ "github.com/glebarez/go-sqlite"
)

type Repository interface {
	Create(chore *Chore) (ID, error)
	//Update(chore chore.Chore) error
	//Delete(chore chore.Chore) error
	//GetById(id int32) (*chore.Chore, error)
}

type SqliteRepository struct {
	db *sql.DB
}

func (r *SqliteRepository) Create(chore *Chore) (ID, error) {
	if i, err := r.initDb(); err != nil {
		return i, err
	}

	id := r.getLastId()

	if i, err := r.addChore(chore, id); err != nil {
		return i, err
	}
	return ID(id), nil
}

func (r *SqliteRepository) addChore(chore *Chore, id int64) (ID, error) {
	res, err := r.db.Exec("INSERT INTO chores (id, title, description) values (?, ?, ?)", id+1, chore.Title, chore.Description)
	fmt.Printf("res: %q \n", res)
	if err != nil {
		fmt.Printf("err: %q \n", err)
		return 0, err
	}
	return 0, nil
}

func (r *SqliteRepository) getLastId() int64 {
	var id int64
	err := r.db.QueryRow("SELECT id FROM chores ORDER BY id DESC LIMIT 1").Scan(&id)
	if err != nil {
		fmt.Printf("error id : %q \n", err)
		id = 0
	}
	return id
}

func (r *SqliteRepository) initDb() (ID, error) {
	if r.db == nil {
		r.db, _ = sql.Open("sqlite", "db")
		_, err := r.db.Exec(`
			CREATE TABLE IF NOT EXISTS chores (
    			id BIGINT PRIMARY KEY,
    			title VARCHAR(255) NOT NULL,
    			description TEXT
			); 
		`)
		if err != nil {
			return 0, err
		}
	}
	return 0, nil
}

type InMemoryRepository struct {
	db []Chore
}

func (r *InMemoryRepository) Create(chore *Chore) (ID, error) {
	if r.db == nil {
		r.db = make([]Chore, 0)
	}
	chore.ID = ID(len(r.db) + 1)
	r.db = append(r.db, *chore)
	return chore.ID, nil
}
