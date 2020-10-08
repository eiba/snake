package main

import "container/heap"

type PriorityNode struct {
	position position
	fScore   int
	index    int
}

type PriorityQueue []*PriorityNode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].fScore < pq[j].fScore
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq PriorityQueue) Exist(value position) (*PriorityNode, bool) {
	for _, priorityNode := range pq {
		if priorityNode.position == value {
			return priorityNode, true
		}
	}
	return nil, false
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*PriorityNode)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *PriorityNode, value position, priority int) {
	item.position = value
	item.fScore = priority
	heap.Fix(pq, item.index)
}