package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	// Utilizei "postgres" como host, que corresponde ao nome do servico no docker-compose, pois a aplicacao esta sendo executada em outro conteiner na mesma rede Docker
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database" + dbname)

	return db, nil
}
