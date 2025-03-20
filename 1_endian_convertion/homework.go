package endian_convertion

import (
	"unsafe"
)

func ToLittleEndian(number uint32) uint32 {
	if number == 0 || number == 0xffffffff {
		return number
	}
	size := int(unsafe.Sizeof(number))

	var result uint32
	for i := 0; i < size; i++ {
		byt := *(*uint8)(unsafe.Add(unsafe.Pointer(&number), i))
		result = (uint32(byt) << ((size - i - 1) * 8)) | result
	}

	return result
}
