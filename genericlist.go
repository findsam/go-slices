package glist

type GenericList[T comparable] struct {
	data []T
}

func New[T comparable]() *GenericList[T] {
	return &GenericList[T]{
		data: []T{},
	}
}

func (l *GenericList[T]) Add(elem T) {
	l.data = append(l.data, elem)
}

func (l *GenericList[T]) Get(i int) T {
	if i > len(l.data)-1 {
		panic("Index out of range")
	}

	for it := 0; it < len(l.data); it++ {
		if i == it {
			return l.data[it]
		}
	}
	// will never trigger but ya good to have
	panic("value not found")
}

func (l *GenericList[T]) Remove(i int) {
	if i > len(l.data)-1 {
		panic("Index out of range")
	}
	for it := 0; it < len(l.data); it++ {
		if it == i {
			// remove from the slice by creaing a new one but preserve the order
			l.data = append(l.data[:it], l.data[it+1:]...)
		}
	}
}

func (l *GenericList[T]) RemoveByValue(value T) {
	for i := 0; i < len(l.data); i++ {
		if l.data[i] == value {
			l.data = append(l.data[:i], l.data[i+1:]...)
		}
	}

}
