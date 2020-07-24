package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/jackc/pgx/v4"
)

var conn *pgx.Conn

func main() {
	var err error
	conn, err = pgx.Connect(context.Background(), "host=192.168.87.210 port=5432 user=benchmarkdbuser password=benchmarkdbpass dbname=todo")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}

	if len(os.Args) == 1 {
		printHelp()
		os.Exit(0)
	}

	switch os.Args[1] {
	case "list":
		err = listTasks()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to list tasks: %v\n", err)
			//os.Exit(1)
			break
		}

	case "add":
		err = addTask(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to add task: %v\n", err)
			//os.Exit(1)
			break
		}

	case "update":
		n, err := strconv.ParseInt(os.Args[2], 10, 32)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable convert task_num into int32: %v\n", err)
			//os.Exit(1)
			break
		}
		err = updateTask(int32(n), os.Args[3])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to update task: %v\n", err)
			//os.Exit(1)
			break
		}

	case "remove":
		n, err := strconv.ParseInt(os.Args[2], 10, 32)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable convert task_num into int32: %v\n", err)
			//os.Exit(1)
			break
		}
		err = removeTask(int32(n))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to remove task: %v\n", err)
			//os.Exit(1)
			break
		}

	default:
		fmt.Fprintln(os.Stderr, "Invalid command")
		printHelp()
		//os.Exit(1)
		break
	}
}

func listTasks() error {
	sql := "select * from tasks"
	fmt.Println("conn.Query sql:", sql)
	rows, err := conn.Query(context.Background(), sql)
	if err != nil {
		fmt.Println("conn.Query, err:", err)
	}

	for rows.Next() {
		var id int32
		var description string
		err := rows.Scan(&id, &description)
		if err != nil {
			return err
		}
		fmt.Printf("%d. %s\n", id, description)
	}

	return rows.Err()
}

func addTask(description string) error {
	sql := "insert into tasks(description) values($1)"
	fmt.Println("addTask conn.Exec sql:", sql, " description:", description)
	cmdTags, err := conn.Exec(context.Background(), sql, description)
	fmt.Println("addTask conn.Exec cmdTags:", cmdTags)
	return err
}

func updateTask(itemNum int32, description string) error {
	sql := "update tasks set description=$1 where id=$2"
	fmt.Println("updateTask conn.Exec sql:", sql, " description:", description, " itemNum:", itemNum)
	cmdTags, err := conn.Exec(context.Background(), sql, description, itemNum)
	fmt.Println("updateTask conn.Exec cmdTags:", cmdTags)
	return err
}

func removeTask(itemNum int32) error {
	sql := "delete from tasks where id=$1"
	fmt.Println("removeTask conn.Exec sql:", sql, " itemNum:", itemNum)
	cmdTags, err := conn.Exec(context.Background(), sql, itemNum)
	fmt.Println("removeTask conn.Exec cmdTags:", cmdTags)
	return err
}

func printHelp() {
	fmt.Print(`Todo pgx demo
Usage:
  todo list
  todo add task
  todo update task_num item
  todo remove task_num
Example:
  todo add 'Learn Go'
  todo list
`)
}
