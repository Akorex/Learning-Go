package main

type Queue struct {
	vals []int
}

func (q Queue) Push(v int) {
	q.vals = append(q.vals, v)
}

func (q Queue) Top() int {
	return q.vals[0]
}

func (q Queue) Dequeue() int {
	if len(q.vals) == 0 {
		return -1
	}
	ans := q.vals[0]

	q.vals = q.vals[1:]
	return ans
}

func main() {

}

// implementation of a LIFO system using a slice as underlying storage
