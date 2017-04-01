package queue

// import (
// 	"fmt"
// )

type Queue struct {
	f func(string)
}

func NewQueue(f func(string)) Queue {
	return Queue{f}
}

func (q Queue) Enqueue(i interface{}) error {
	q.f(i.(string))
	return nil
}
