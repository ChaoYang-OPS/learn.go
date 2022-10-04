package recall

import "learn.go/testNe/rac_sys/common"

type Recaller interface {
	Recall(n int) []*common.Product
	Name() string
}
