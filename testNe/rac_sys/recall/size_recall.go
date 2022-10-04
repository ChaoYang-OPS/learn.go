package recall

import "learn.go/testNe/rac_sys/common"

type SizeRecall struct {
	Tag string
}

func (s SizeRecall) Name() string {
	return s.Tag
}
func (s SizeRecall) Recall(n int) []*common.Product {
	rect := make([]*common.Product, 0, n)
	for _, item := range allProducts {
		if item.Size < 200 { // 只召回size小于200的商品
			rect = append(rect, item)
			if len(rect) >= n {
				//return
				break
			}
		}
	}
	return rect
}
