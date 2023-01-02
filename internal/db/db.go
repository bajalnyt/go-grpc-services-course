package db

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/bajalnyt/go-grpc-services-course/internal/rocket"

	"github.com/jmoiron/sqlx"
)

type Store struct {
	db *sqlx.DB
}

func New() (Store, error) {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")

	connectString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbTable, dbPassword)

	db, err := sqlx.Connect("postgres", connectString)
	if err != nil {
		return Store{}, err
	}
	return Store{
		db: db,
	}, nil
}

func (s Store) GetRocketById(id string) (rocket.Rocket, error) {
	var rkt rocket.Rocket
	integer_id, err := strconv.Atoi(id)
	if err != nil {
		log.Print(err.Error())
		return rocket.Rocket{}, err
	}
	row := s.db.QueryRow(
		`SELECT ID from rockets where id=$1`,
		integer_id,
	)
	err = row.Scan(&rkt.ID)
	if err != nil {
		log.Print(err.Error())
		return rocket.Rocket{}, err
	}

	return rkt, nil
}

// InsertRocket - inserts a rocket into the rockets table
func (s Store) InsertRocket(rkt rocket.Rocket) (rocket.Rocket, error) {
	_, err := s.db.NamedQuery(
		`INSERT INTO rockets
		(id, name, type)
		VALUES (:id, :name, :type)`,
		rkt,
	)
	if err != nil {
		return rocket.Rocket{}, errors.New("failed to insert into database")
	}
	return rocket.Rocket{
		ID:   rkt.ID,
		Type: rkt.Type,
		Name: rkt.Name,
	}, nil
}

func (s Store) DeleteRocketById(id string) error {
	return nil
}
