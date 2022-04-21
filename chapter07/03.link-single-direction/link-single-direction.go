package main

import "fmt"

type LinkNode struct {
	data int
	next *LinkNode
}

func main() {
	n1 := &LinkNode{
		data: 1,
		next: nil,
	}
	n2 := &LinkNode{
		data: 2,
		next: nil,
	}
	n3 := &LinkNode{
		data: 3,
		next: nil,
	}
	n4 := &LinkNode{
		data: 4,
		next: nil,
	}
	n6 := &LinkNode{
		data: 6,
		next: nil,
	}
	n1.next = n2
	n2.next = n3
	n3.next = n4
	n4.next = n6

	tmpNode := n1
	for {
		if tmpNode.next != nil {
			fmt.Println(tmpNode.data)
		} else {
			break
		}
		tmpNode = tmpNode.next
	}

	{
		fmt.Println("插入5")
		n5 := &LinkNode{
			data: 5,
			next: nil,
		}
		tmpNode := n1
		for {
			if tmpNode.next != nil {
				fmt.Println(tmpNode.data)
				if n5.data > tmpNode.data {
					if tmpNode.next == nil {
						// 已经到结尾，直接追加
						tmpNode.next = n5
					} else {
						if tmpNode.next.data >= n5.data {
							// 找到合适位置，准备插入数据
							n5.next = tmpNode.next
							tmpNode.next = n5
							break
						}
					}
				}
			} else {
				break
			}
			tmpNode = tmpNode.next
		}
	}

}
