package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
)

var db *sql.DB

func main() {
	os.Remove("./foo.db")

	var err error
	db, err = sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table foo (id integer not null primary key, name text);
	delete from foo;
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	var input string
	for {
		fmt.Scanln(&input)
		input = strings.ToLower(input)
		if input == "e" || input == "exit" {
			break
		}
		if input == "i" {
			doInsert()
		} else if input == "d" {
			doDelete()
		} else if input == "da" {
			doDeleteAll()
		} else if input == "u" {
			doUpdate()
		} else if input == "s" {
			doSelect()
		} else if input == "sa" {
			doSelectAll()
		}
	}
}
func pResult(str string, result sql.Result) {
	resultStr := "result => "
	LastInsertId, err := result.LastInsertId()
	if err != nil {
		resultStr += "LastInsertId err:"
		resultStr += err.Error()
	} else {
		resultStr += "LastInsertId:"
		resultStr += strconv.FormatInt(LastInsertId, 10)
	}
	resultStr += " "
	RowsAffected, err := result.RowsAffected()
	if err != nil {
		resultStr += "RowsAffected err:"
		resultStr += err.Error()
	} else {
		resultStr += "RowsAffected:"
		resultStr += strconv.FormatInt(RowsAffected, 10)
	}
	log.Println(str, " result:", resultStr)
}
func doInsert() {
	tx, err := db.Begin()
	if err != nil {
		log.Println("doInsert err:", err)
	}
	stmt, err := tx.Prepare("insert into foo(id, name) values(?, ?)")
	if err != nil {
		log.Println("doInsert err:", err)
	}
	defer stmt.Close()

	for i := 0; i < 3; i++ {
		result, err := stmt.Exec(i, fmt.Sprintf("こんにちわ世界%03d", i))
		if err != nil {
			log.Println("doInsert stmt.Exec err:", err)
		}
		pResult("doInsert", result)
	}
	err = tx.Commit()
	if err != nil {
		log.Println("tx.Commit err:", err)
	}
}
func doDelete() {
	result, err := db.Exec("delete from foo where id = 1")
	if err != nil {
		log.Println("doDelete err:", err)
	}
	pResult("doDelete", result)
}
func doDeleteAll() {
	result, err := db.Exec("delete from foo")
	if err != nil {
		log.Println("doDelete err:", err)
	}
	pResult("doDelete", result)
}
func doUpdate() {
	tx, err := db.Begin()
	if err != nil {
		log.Println("doUpdate err:", err)
	}
	stmt, err := tx.Prepare("update foo set name = ? where id = ?")
	if err != nil {
		log.Println("doUpdate err:", err)
	}
	defer stmt.Close()
	for i := 0; i < 3; i++ {
		_, err = stmt.Exec(fmt.Sprintf("こんにちわ世界%03d", i+10), i)
		if err != nil {
			log.Println("doUpdate err:", err)
		}
	}
	tx.Commit()
}
func doSelect() {
	stmt, err := db.Prepare("select name from foo where id = ?")
	if err != nil {
		log.Println("doSelect err:", err)
	}
	defer stmt.Close()
	var name string
	err = stmt.QueryRow("0").Scan(&name)
	if err != nil {
		log.Println("doSelect err:", err)
	}
	fmt.Println(name)
}
func doSelectAll() {
	stmt, err := db.Prepare("select id, name from foo")
	if err != nil {
		log.Println("doSelect err:", err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Error("doSelectAll stmt.Query err:", err)
		return
	}

	var id int64
	var name string
	for rows.Next() {
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Error("doSelectAll rows.Scan err:", err)
			continue
		}
		log.Println("doSelectAll rows.Scan id:", id, " name:", name)
	}
}
