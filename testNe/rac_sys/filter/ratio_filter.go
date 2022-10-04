package filter

import (
	"learn.go/testNe/rac_sys/common"
)

type RatioFilter struct {
	Tag string
}

func (s RatioFilter) Name() string {
	return s.Tag
}
func (s RatioFilter) Filter(products []*common.Product) []*common.Product {
	rect := make([]*common.Product, 0, len(products))
	for _, product := range products {
		if product.RatioCount >= 0 && product.PositiveRatio >= 0 {
			rect = append(rect, product)
		}
	}
	return rect
}
