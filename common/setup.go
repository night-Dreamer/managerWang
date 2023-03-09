package common

import "database/sql"

var Db *sql.DB

type App struct {
	Id   string
	Key  string
	Role string
}

type User_p struct {
	Id      string
	Name    string
	Purview []string
}
