package sort

import "learn.go/testNe/rac_sys/common"

type Sorter interface {
	Sort([]*common.Product) []*common.Product
	Name() string
}
