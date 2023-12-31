package server

import (
	"database/sql"
	"errors"
	"log"
	"time"

	api "github.com/retr0h/sma/activity-log/api/v1"
	"google.golang.org/protobuf/types/known/timestamppb"

	_ "github.com/mattn/go-sqlite3"
)

const create string = `
		CREATE TABLE IF NOT EXISTS activities (
		id INTEGER NOT NULL PRIMARY KEY,
		time DATETIME NOT NULL,
		description TEXT
		);
`
const file string = "activities.db"

type Activities struct {
	db *sql.DB
}

func NewActivities() (*Activities, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}
	if _, err := db.Exec(create); err != nil {
		return nil, err
	}
	return &Activities{
		db: db,
	}, nil
}

func (c *Activities) Insert(activity *api.Activity) (int, error) {
	res, err := c.db.Exec("INSERT INTO activities VALUES(NULL,?,?);",
		activity.Time.AsTime(),
		activity.Description,
	)
	if err != nil {
		return 0, err
	}

	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return 0, err
	}
	log.Printf("Added %v as %d", activity, id)
	return int(id), nil
}

var ErrIDNotFound = errors.New("Id not found")

func (c *Activities) Retrieve(id int) (*api.Activity, error) {
	log.Printf("Getting %d", id)

	// Query DB row based on ID
	row := c.db.QueryRow("SELECT id, time, description FROM activities WHERE id=?", id)

	// Parse row into Interval struct
	activity := api.Activity{}
	var err error
	var time time.Time
	if err = row.Scan(&activity.Id, &time, &activity.Description); err == sql.ErrNoRows {
		log.Printf("Id not found")
		return &api.Activity{}, ErrIDNotFound
	}
	activity.Time = timestamppb.New(time)
	return &activity, err
}

func (c *Activities) List(offset int) ([]*api.Activity, error) {
	log.Printf("Getting list from offset %d\n", offset)

	// Query DB row based on ID
	rows, err := c.db.Query("SELECT * FROM activities WHERE ID > ? ORDER BY id DESC LIMIT 100", offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := []*api.Activity{}
	for rows.Next() {
		i := api.Activity{}
		var time time.Time
		err = rows.Scan(&i.Id, &time, &i.Description)
		if err != nil {
			return nil, err
		}
		i.Time = timestamppb.New(time)
		data = append(data, &i)
	}
	return data, nil
}
