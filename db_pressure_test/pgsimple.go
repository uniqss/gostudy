package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
	"runtime"
)

const (
	DbHost = "192.168.87.210"
	DbPort = 5432
	//DbUser  = "postgres"
	//DbPaswd = "postgres"
	DbUser  = "benchmarkdbuser"
	DbPaswd = "benchmarkdbpass"
	DbName  = "hello_world"

	WorldSelectSQL      = "SELECT id, random_number FROM world WHERE id = $1"
	SelectAllTablesSQL  = "select * from pg_tables"
	WorldSelectCacheSQL = "SELECT id, random_number FROM world LIMIT $1"
	WorldUpdateSQL      = "UPDATE world SET random_number = $1 WHERE id = $2"
	FortuneSelectSQL    = "SELECT id, message FROM fortune"
)

var DbSimple *pgxpool.Pool

func InitDBSimple(maxConn int) error {
	pgx, err := newPGXSimple(maxConn)
	if err != nil {
		return err
	}

	DbSimple = pgx

	return nil
}

func newPGXSimple(maxConn int) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s pool_max_conns=%d",
		DbHost, DbPort, DbUser, DbPaswd, DbName, maxConn,
	)

	return pgxpool.Connect(context.Background(), dsn)
}

func numCPUSimple() int {
	n := runtime.NumCPU()
	if n == 0 {
		n = 8
	}

	return n
}

func doInsert(ctx context.Context, sql string, userId int) {
	infoJson := "{\"uinfo\":{\"level\":123, \"name\":\"肉山大王\", \"score\":100}}"
	_, err := DbSimple.Exec(ctx, sql, userId, infoJson)
	if err != nil {
		log.Error("DbSimple.Exec err:", err, " sql:", sql)
	}
}
func doDelete(ctx context.Context, sql string, userId int) {
	_, err := DbSimple.Exec(ctx, sql, userId)
	if err != nil {
		log.Error("DbSimple.Exec err:", err, " sql:", sql)
	}
}
func doUpdate(ctx context.Context, sql string, userId int) {
	_, err := DbSimple.Exec(ctx, sql, 456, userId)
	if err != nil {
		log.Error("DbSimple.Exec err:", err, " sql:", sql)
	}
}
func doSelect(ctx context.Context, sql string, userId int) {
	_, err := DbSimple.Exec(ctx, sql, userId)
	if err != nil {
		log.Error("DbSimple.Exec err:", err, " sql:", sql)
	}
}

func main() {
	maxConn := numCPUSimple() * 4
	fmt.Println("maxConn:", maxConn)

	err := InitDBSimple(maxConn)
	if err != nil {
		log.Error("InitDBSimple err:", err)
	}

	ctx := context.Background()

	userId := 50000000
	var sql string
	sql = "insert into pressure_test(uid, info_json) values($1, $2)"
	doInsert(ctx, sql, userId)

	sql = "delete from pressure_test where uid = $1"
	doDelete(ctx, sql, userId)

	sql = "insert into pressure_test(uid, info_json) values($1, $2)"
	doInsert(ctx, sql, userId)

	sql = "update pressure_test set info_json.level = $1 where uid = $2"
	doUpdate(ctx, sql, userId)

	sql = "select * from pressure_test where uid = $1"
	doSelect(ctx, sql, userId)
}
