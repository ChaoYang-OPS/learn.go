package sort

import (
	"learn.go/testNe/rac_sys/common"
	"sort"
)

type RatioSorter struct {
	Tag string
}

func (s RatioSorter) Name() string {
	return s.Tag
}
func (s RatioSorter) Sort(products []*common.Product) []*common.Product {
	sort.Slice(products, func(i, j int) bool {
		return products[i].PositiveRatio > products[j].PositiveRatio
	})
	return products
}
