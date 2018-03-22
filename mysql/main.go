package main

import (
	"database/sql"
	"flag"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func main() {
	host := flag.String("host", "127.0.0.1", "host server")
	port := flag.Int("port", 3306, "port")
	user := flag.String("user", "", "username")
	pass := flag.String("pass", "", "users password")
	db := flag.String("db", "", "database to connect to")
	flag.Parse()

	config := newConfig(*host, *port, *user, *pass, *db)
	conn := connect(config)

	currencies := queryCurrencies(conn)

	for _, currency := range currencies {
		fmt.Println(currency.Name)
	}
}

type currency struct {
	Name string
}

func newConfig(host string, port int, user string, pass string, db string) mysql.Config {
	return mysql.Config{
		Net:    "tcp",
		Addr:   fmt.Sprintf("%s:%d", host, port),
		User:   user,
		Passwd: pass,
		DBName: db,
	}
}

func connect(config mysql.Config) *sql.DB {
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("connected to %s\n", config.Addr)
	return db
}

func queryCurrencies(db *sql.DB) []*currency {
	rows, err := db.Query(`
		select c.currency_name
		from currency c
		limit 10;
	`)

	if err != nil {
		panic(err)
	}

	defer rows.Close()
	return readCurrencyRows(rows)
}

func readCurrencyRows(rows *sql.Rows) []*currency {
	results := make([]*currency, 0, 0)
	for rows.Next() {
		result := currency{}
		err := rows.Scan(
			&result.Name,
		)
		if err != nil {
			panic(err)
		}
		results = append(results, &result)
	}
	return results
}
