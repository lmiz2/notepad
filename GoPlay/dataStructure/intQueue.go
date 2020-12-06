package intqueue

import (
	"strconv"
)

type myQueue struct {
	size      int
	firstNode *node
	lastNode  *node
}

type node struct {
	value int
	next  *node
}

func newMyQueue() myQueue {
	obj := myQueue{0, nil, nil}
	return obj
}
func (q *myQueue) first() int {
	if q.size > 0 {
		return q.firstNode.value
	}
	return -1
}
func (q *myQueue) add(num int) {
	obj := node{num, nil}
	if q.size <= 0 {
		q.firstNode = &obj
		q.lastNode = &obj
	} else {
		q.lastNode.next = &obj
		q.lastNode = &obj
	}
	q.size++
}
func (q *myQueue) poll() int {
	if q.size > 0 {
		tmpPtr := q.firstNode
		q.firstNode = q.firstNode.next
		q.size--
		return tmpPtr.value
	}
	return -1
}
func (q *myQueue) isEmpty() bool {
	return q.size <= 0
}
func (q *myQueue) String() (rs string) {
	if q.size > 0 {
		tnode := q.firstNode
		for tnode != nil {
			rs += tnode.String() + " "
			tnode = tnode.next
		}
	}
	return
}
func (n node) String() string {
	return "[" + strconv.Itoa(n.value) + "]"
}
func test() {

	// que := newMyQueue()
	// que.add(1)
	// que.add(3)
	// que.add(2)
	// que.add(5)
	// fmt.Println(que.String())
	// fmt.Println(que.poll())
	// fmt.Println(que.poll())
	// fmt.Println(que.poll())
	// fmt.Println(que.poll())
	// fmt.Println(que.poll())
	// fmt.Println(que.poll())
	// que.add(3)
	// que.add(2)
	// fmt.Println(que.poll())
	// fmt.Println(que.poll())
	// fmt.Println(que.poll())
}
