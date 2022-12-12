package sync_pools

import (
	"sync"
	"testing"
)

// 一般情况下，如果要考虑缓存资源，比如说创建好的对象，那么可以使用 sync.Pool
//  sync.Pool 会先查看自己是否有资源， 有则直接返回， 没有则创建一个新的
//  sync.Pool 会在GC的时候释放缓存的资源
// 一般使用 sync.Pool 都是为了复用内存
// 它减少了内存分配，也减轻了GC压力（最主要）
func TestPool(t *testing.T) {
	p := sync.Pool{
		New: func() interface{} {
			// 创建函数， sync.Pool会回调
			return nil
		},
	}
	obj := p.Get()
	// 在这里使用取出来的对象
	// 用完再放回去
	p.Put(obj)
}
