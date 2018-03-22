package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "127.0.0.1", "host server")
	port := flag.Int("port", 3306, "port")
	user := flag.String("user", "", "username")
	pass := flag.String("pass", "", "users password")
	db := flag.String("db", "", "database to connect to")
	flag.Parse()

	connStr := build(host, port, user, pass, db)

	fmt.Println(connStr)
}

func build(host *string, port *int, user *string, pass *string, db *string) string {
	connStr := ""
	if *user != "" && *pass != "" {
		connStr += *user + ":" + *pass + "@"
	}
	connStr += "tcp(" + *host + ":" + fmt.Sprintf("%d", *port) + ")"

	if *db != "" {
		connStr += "/" + *db
	}

	return connStr
}
