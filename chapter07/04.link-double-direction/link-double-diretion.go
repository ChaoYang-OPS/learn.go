package main

import "fmt"

type LinkNode struct {
	data int
	next *LinkNode
	prev *LinkNode
}

// todo 后期补全算法与数据结构相关的代码
func buildDLink() *LinkNode {
	n1 := &LinkNode{
		data: 1,
	}
	n2 := &LinkNode{
		data: 5,
	}
	n3 := &LinkNode{
		data: 10,
	}
	n1.next = n2
	n2.prev = n1
	n2.next = n3
	n3.prev = n2
	return n1
}

func insertNode(root *LinkNode, newNode *LinkNode) *LinkNode {
	tmpNode := root
	// 整个链表是空的情况，新增
	if root == nil {
		return newNode
	}
	// 在链表的头添加节点
	if root.data >= newNode.data {
		newNode.next = tmpNode
		tmpNode.prev = newNode
		return newNode
	}
	for {
		if tmpNode.next == nil {
			// 已经到头，追加节点即可
			tmpNode.next = newNode
			newNode.prev = tmpNode
		} else {
			if tmpNode.next.data >= newNode.data {
				// 找到位置，在此插入新节点
				newNode.prev = tmpNode
				newNode.next = tmpNode.next

				tmpNode.next = newNode
				newNode.next.prev = newNode
				return root
			}
		}
		tmpNode = tmpNode.next
	}
}
func rangeLink(root *LinkNode) {
	fmt.Println("从头到尾")
	tmpNode := root
	for {
		fmt.Println(tmpNode.data)
		if tmpNode.next == nil {
			break
		}
		tmpNode = tmpNode.next
	}
	fmt.Println("从尾到头")
	for {
		fmt.Println(tmpNode.data)
		if tmpNode.prev == nil {
			break
		}
		tmpNode = tmpNode.prev
	}
}
func main() {
	root := buildDLink()
	insertNode(root, &LinkNode{data: 3})
	insertNode(root, &LinkNode{data: 7})
	insertNode(root, &LinkNode{data: 11})
	insertNode(root, &LinkNode{data: 0})
	rangeLink(root)
}
