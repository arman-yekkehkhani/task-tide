package chore

import (
	"database/sql"
	"fmt"
	chore2 "github.com/arman-yekkehkhani/task-tide/internal/model/chore"
	_ "github.com/glebarez/go-sqlite"
)

type Repository interface {
	Create(chore *chore2.Chore) (chore2.ID, error)
	GetByID(id chore2.ID) *chore2.Chore
	Update(c *chore2.Chore) (*chore2.Chore, error)
}

type SqliteRepository struct {
	db *sql.DB
}

func NewSqliteRepository(source string) *SqliteRepository {
	r := &SqliteRepository{}
	r.initDb(source)
	return r
}

func (r *SqliteRepository) Create(chore *chore2.Chore) (chore2.ID, error) {
	id := r.getLastId()
	if i, err := r.addChore(chore, id); err != nil {
		return i, err
	}
	return chore2.ID(id), nil
}

func (r *SqliteRepository) addChore(chore *chore2.Chore, id int64) (chore2.ID, error) {
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

func (r *SqliteRepository) initDb(source string) (chore2.ID, error) {
	if r.db == nil {
		r.db, _ = sql.Open("sqlite", source)
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
