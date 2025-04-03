package main

import (
	"unsafe"
)

type COWBuffer struct {
	data []byte
	refs *int
}

func NewCOWBuffer(data []byte) COWBuffer {
	i := 1
	return COWBuffer{
		data: data,
		refs: &i,
	}
}

func (b *COWBuffer) Clone() COWBuffer {
	*b.refs++
	return COWBuffer{
		data: b.data,
		refs: b.refs,
	}
}

func (b *COWBuffer) Close() {
	*b.refs -= 1
	if *b.refs <= 0 {
		*&b.data = nil
	}
}

func (b *COWBuffer) Update(index int, value byte) bool {
	if index < 0 || index >= len(b.data) {
		return false
	}

	if b.data == nil {
		return false
	}

	if *b.refs > 1 {
		*b.refs--
		i := 1
		b.refs = &i
		d := make([]byte, len(b.data))
		copy(d, b.data)
		b.data = d
	}

	(b.data)[index] = value

	return true
}

func (b *COWBuffer) String() string {
	return unsafe.String(unsafe.SliceData(b.data), len(b.data))
}
