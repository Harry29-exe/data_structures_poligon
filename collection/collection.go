package collection

type Iterator[T any] interface {
	HasNext() bool
	Next() T
}

type BiIterator[K any, T any] interface {
	HasNext() bool
	Next() (K, T)
}

type Collection[T any] interface {
	Size() uint
	Insert(T)
	InsertAll(Iterator[T])
	Contains(T) bool
	Delete(T)
	DeleteAll(Iterator[T])
	Iterator() Iterator[T]
}

type List[T any] interface {
	Collection[T]
	GetAt(uint) T
	InsertAt(uint, T)
	DeleteAt(uint) T
}

type Set[T any] interface {
	Insert(T)
	Contains(T) bool
	Iterator() Iterator[T]
	Delete(T)
}

type Map[K any, T any] interface {
	Insert(K, T)
	Contains(K) bool
	Get(K) T
	Delete(K)
	Iterator() BiIterator[K, T]
}
