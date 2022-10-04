package main

import (
	"fmt"
	"learn.go/testNe/rac_sys/common"
	"learn.go/testNe/rac_sys/filter"
	"learn.go/testNe/rac_sys/recall"
	"learn.go/testNe/rac_sys/sort"
	"log"
	"time"
)

type Recommender struct {
	// 建议接口这边不要放指针
	ReCallers []recall.Recaller
	Sorter    sort.Sorter
	Filters   []filter.Filter
}

func (rec *Recommender) Rec() []*common.Product {
	RecallMap := make(map[int]*common.Product, 100)
	// 顺序遍历每一路召回
	for _, recalled := range rec.ReCallers {
		begin := time.Now()
		products := recalled.Recall(10)
		log.Printf("召回%s耗时%d,召回了%d商品\n", recalled.Name(), time.Since(begin).Nanoseconds(), len(products))
		for _, product := range products {
			RecallMap[product.Id] = product
		}
	}
	log.Printf("一共召回了%d商品\n", len(RecallMap))
	// 把map转换为切片
	RecallSlice := make([]*common.Product, 0, len(RecallMap))
	for _, product := range RecallMap {
		RecallSlice = append(RecallSlice, product)
	}
	// 对召回结果进行排序 time.Since
	begin := time.Now()
	SortedResult := rec.Sorter.Sort(RecallSlice)
	log.Printf("排序花费%dns\n", time.Since(begin).Nanoseconds())
	//顺利执行多个过滤规则
	FilterResult := SortedResult
	prevCount := len(FilterResult)
	for _, filters := range rec.Filters {
		begin := time.Now()
		FilterResult = filters.Filter(FilterResult)
		log.Printf("过滤规则%s耗时%dns 过滤%d个商品\n", filters.Name(),
			time.Since(begin).Nanoseconds(), prevCount-len(FilterResult))
		prevCount = len(FilterResult)
	}
	return FilterResult
}

func main() {
	recer := Recommender{
		ReCallers: []recall.Recaller{
			recall.HotRecall{
				Tag: "hot",
			},
			recall.SizeRecall{Tag: "size"},
		},
		Sorter: sort.RatioSorter{Tag: "Ratio"},
		Filters: []filter.Filter{
			filter.RatioFilter{
				Tag: "Ratio",
			},
		},
	}
	result := recer.Rec()
	for i, p := range result {
		fmt.Printf("第%d名：%d %s \n", i, p.Id, p.Name)
	}
}
