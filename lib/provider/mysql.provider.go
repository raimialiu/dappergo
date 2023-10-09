package provider

import "database/sql"

type MysqlProvider struct {
}

func New() {
	sql.Open()
}
