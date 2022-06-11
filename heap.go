package heap

func createHeap[T any](cap int, isMaxHeap bool, less func(a, b T) bool) *heap[T] {
	return &heap[T]{
		[]T{},
		isMaxHeap,
		less,
		cap,
	}
}

//if a<b expect less(a,b)=true
func MaxHeap[T any](less func(a, b T) bool, cap ...int) *heap[T] {
	if len(cap) == 0 {
		return createHeap(0, true, less)
	}
	return createHeap(cap[0], true, less)
}

//if a<b expect less(a,b)=true
func MinHeap[T any](less func(a, b T) bool, cap ...int) *heap[T] {
	if len(cap) == 0 {
		return createHeap(0, false, less)
	}
	return createHeap(cap[0], false, less)
}

type heap[T any] struct {
	store   []T
	maxHeap bool
	less    func(T, T) bool
	cap     int
}

func (h *heap[T]) Push(t T) {
	task := func() {
		defer h.shiftUp(len(h.store))
		h.store = append(h.store, t)
	}
	if h.cap == 0 || len(h.store) < h.cap {
		task()
	} else {
		less := h.less(t, h.Peek())
		if (h.maxHeap && less) || (!h.maxHeap && !less) {
			defer h.Pop()
			task()
		}

	}
}
func (h *heap[T]) Pop() (item T) {
	return h.Remove(0)
}
func (h *heap[T]) Peek() (item T) {
	if len(h.store) == 0 {
		return
	}
	return h.store[0]
}
func (h *heap[T]) shouldUp(son, p T) bool {
	less := h.less(son, p)
	return (h.maxHeap && !less) || (!h.maxHeap && less)
}
func (h *heap[T]) shiftUp(pos int) {
	if pos <= 0 {
		return
	}
	p := parent(pos)
	if h.shouldUp(h.store[pos], h.store[p]) {
		h.swap(pos, p)
		h.shiftUp(p)
	}
}
func (h *heap[T]) swap(a, b int) {
	h.store[a], h.store[b] = h.store[b], h.store[a]
}
func (h *heap[T]) shiftDown(pos int) {
	l, r := lSon(pos), rSon(pos)
	if pos < len(h.store) {
		if l < len(h.store) && h.shouldUp(h.store[l], h.store[pos]) {
			h.swap(l, pos)
			h.shiftDown(l)
		}
		if r < len(h.store) && h.shouldUp(h.store[r], h.store[pos]) {
			h.swap(r, pos)
			h.shiftDown(r)
		}
	}
}
func (h *heap[T]) Init(ts []T) {
	for _, item := range ts {
		h.Push(item)
	}
}
func (h *heap[T]) Remove(pos int) (t T) {
	if len(h.store) <= pos {
		return
	}
	t = h.store[pos]
	h.store[pos] = h.store[len(h.store)-1]
	h.store = h.store[:len(h.store)-1]
	h.shiftDown(pos)
	return
}
func (h *heap[T]) Fix(pos int) {
	h.Push(h.Remove(pos))
}
func parent(pos int) int {
	return (pos - 1) >> 1
}
func lSon(pos int) int {
	return pos*2 + 1
}
func rSon(pos int) int {
	return pos*2 + 2
}
