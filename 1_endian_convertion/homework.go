package endian_convertion

import (
	"unsafe"
)

type Num interface {
	uint16 | uint32 | uint64
}

func ToLittleEndian[T Num](number T) T {
	if number == 0 {
		return number
	}
	size := int(unsafe.Sizeof(number))
	const byteSize = 8

	var result T
	for i := 0; i < size; i++ {
		byt := *(*uint8)(unsafe.Add(unsafe.Pointer(&number), i))
		result = (T(byt) << ((size - i - 1) * byteSize)) | result
	}

	return result
}
