package recall

import (
	"learn.go/testNe/rac_sys/common"
	"sort"
)

var allProducts = []*common.Product{
	{Id: 1, Sale: 38, Name: "p1"},
	{Id: 2, Sale: 12338, Name: "p2"},
	{Id: 3, Sale: 381, Name: "p3"},
	{Id: 4, Sale: 348, Name: "p4"},
	{Id: 5, Sale: 34448, Name: "p5"},
	{Id: 6, Sale: 35538, Name: "p6"},
	{Id: 7, Sale: 3148, Name: "p7"},
	{Id: 8, Sale: 3418, Name: "p8"},
}

type HotRecall struct {
	Tag string
}

func (r HotRecall) Name() string {
	return r.Tag
}
func (r HotRecall) Recall(n int) []*common.Product {
	sort.Slice(allProducts, func(i, j int) bool {
		return allProducts[i].Sale > allProducts[j].Sale // 销量从大到小排序
	})
	rect := make([]*common.Product, 0, n)
	for _, product := range allProducts {
		rect = append(rect, product)
		if len(rect) >= n {
			//return
			break
		}
	}
	return rect
}
