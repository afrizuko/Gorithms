package gorithms

type Queue struct {
	values []interface{}
}

func (q *Queue) IsEmpty() bool {
	return len(q.values) == 0
}

func (q *Queue) Size() int {
	return len(q.values)
}

func (q *Queue) Enqueue(item interface{}) {
	q.values = append(q.values, item)
}

func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	v := q.values[0]
	q.values = q.values[1:]
	return v
}
