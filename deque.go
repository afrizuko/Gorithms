package gorithms

// combines power of both a queue and stack..
// one can stack and pop items from both ends

type Deque struct {
	values []interface{}
}

func (d *Deque) IsEmpty() bool {
	return d.Size() == 0
}

func (d *Deque) Size() int {
	return len(d.values)
}

func (d *Deque) AddFront(item interface{}) {
	// insert item to last position of slice
	d.values = append(d.values, item)
}

func (d *Deque) AddRear(item interface{}) {
	temp := []interface{}{item}
	d.values = append(temp, d.values...)
}

func (d *Deque) PopFront() interface{} {
	last := d.Size() - 1
	defer func() {
		d.values = d.values[:last]
	}()
	return d.values[last]
}

func (d *Deque) PopRear() interface{} {
	defer func() {
		d.values = d.values[1:]
	}()
	return d.values[0]
}
