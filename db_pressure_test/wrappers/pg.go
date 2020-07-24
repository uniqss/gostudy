package wrappers

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	DbHost  = "192.168.87.210"
	DbPort  = 5432
	//DbUser  = "postgres"
	//DbPaswd = "postgres"
	DbUser  = "benchmarkdbuser"
	DbPaswd = "benchmarkdbpass"
	DbName  = "hello_world"

	WorldSelectSQL      = "SELECT id, random_number FROM world WHERE id = $1"
	SelectAllTablesSQL = "select * from pg_tables"
	WorldSelectCacheSQL = "SELECT id, random_number FROM world LIMIT $1"
	WorldUpdateSQL      = "UPDATE world SET random_number = $1 WHERE id = $2"
	FortuneSelectSQL    = "SELECT id, message FROM fortune"
)

var Db *pgxpool.Pool

func InitDB(maxConn int) error {
	pgx, err := newPGX(maxConn)
	if err != nil {
		return err
	}

	Db = pgx

	return nil
}

func CloseDB() {
	Db.Close()
}

func newPGX(maxConn int) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s pool_max_conns=%d",
		DbHost, DbPort, DbUser, DbPaswd, DbName, maxConn,
	)

	return pgxpool.Connect(context.Background(), dsn)
}
