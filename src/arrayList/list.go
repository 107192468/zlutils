package arrayList

import ()
const(
)
type ArrayList struct {
	capacity, size int
	data           []interface{}
}

func NewList() *ArrayList {
	return &ArrayList{
		capacity: 10,
		size:     0,
		data:     make([]interface{}, 12),
	}
}
func (list *ArrayList) Add(data interface{}) {
	
	
	if list.size < list.capacity {
		list.data[list.size] = data
		list.size++
	}
}
