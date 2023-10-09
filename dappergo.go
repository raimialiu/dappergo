package dappergo

import "dappergo/lib/types"

type DapperGo struct {
	config *types.DapperGoConfig
}

func NewInstance(config types.DapperGoConfig) *DapperGo {
	return &DapperGo{
		config: &config,
	}
}

func (d *DapperGo) Open() {

}
