package filter

import "learn.go/testNe/rac_sys/common"

type Filter interface {
	Filter([]*common.Product) []*common.Product
	Name() string
}
f