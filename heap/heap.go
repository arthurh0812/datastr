package heap

import "github.com/arthurh0812/datastr/types"

type Heap struct {
	arr []types.Value
	max bool // default=min
}

func (h *Heap) getChildren(idx int) (left, right types.Value) {
	if h == nil {
		return nil, nil
	}
	return h.arr[idx*2+1], h.arr[idx*2+2]
}

func (h *Heap) getParent(idx int) (parent types.Value) {
	if h == nil {
		return nil
	}
	return h.arr[(idx-1)/2]
}

func (h *Heap) Poll() (val types.Value) {
	return nil
}

func (h *Heap) Add(val types.Value) {

}

func (h *Heap) Peek() (val types.Value) {
	return nil
}

// MakeMin makes this heap a minimum priority queue; this is the default.
func (h *Heap) MakeMin() {
	h.max = false
}

// MakeMax makes this heap a maximum priority queue.
func (h *Heap) MakeMax() {
	h.max = true
}