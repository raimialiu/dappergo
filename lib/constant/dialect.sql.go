package constant

type SqlDialect int

const (
	MYSQL     SqlDialect = 1433
	POSTGRESS SqlDialect = 1453
	SQLSERVER SqlDialect = 1450
	SQLLITE   SqlDialect = 1475
)
