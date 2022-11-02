package sort

import (
	"learn.go/testNe/rac_sys/common"
	"sort"
)

type SizeSorter struct {
	Tag string
}

func (s SizeSorter) Name() string {
	return s.Tag
}
func (s SizeSorter) Sort(products []*common.Product) []*common.Product {
	sort.Slice(products, func(i, j int) bool {
		return products[i].Size > products[j].Size
	})
	return products
}
