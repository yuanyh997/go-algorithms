package main

type Heap []int

func (h *Heap) Insert(val int) {
	h = append(h,val)
	for i:=len(h);i>=0;{

		i = i/2
	}
}