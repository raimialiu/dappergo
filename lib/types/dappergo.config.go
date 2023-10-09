package types

import "dappergo/lib/constant"

type DapperGoConfig struct {
	connectionString string
	dialect          constant.SqlDialect
	autoConnect      *bool
}
