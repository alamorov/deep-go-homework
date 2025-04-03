package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestСonversion(t *testing.T) {
	tests := map[string]struct {
		number uint32
		result uint32
	}{
		"test case #1": {
			number: 0x00000000,
			result: 0x00000000,
		},
		"test case #2": {
			number: 0xFFFFFFFF,
			result: 0xFFFFFFFF,
		},
		"test case #3": {
			number: 0x00FF00FF,
			result: 0xFF00FF00,
		},
		"test case #4": {
			number: 0x0000FFFF,
			result: 0xFFFF0000,
		},
		"test case #5": {
			number: 0x01020304,
			result: 0x04030201,
		},
	}

	tests16 := map[string]struct {
		number uint16
		result uint16
	}{
		"test case #1": {
			number: 0x0102,
			result: 0x0201,
		},
		"test case #2": {
			number: 0x01FF,
			result: 0xFF01,
		},
	}

	tests64 := map[string]struct {
		number uint64
		result uint64
	}{
		"test case #1": {
			number: 0x0102030405060708,
			result: 0x0807060504030201,
		},
		"test case #2": {
			number: 0x0000FFFF0000FFFF,
			result: 0xFFFF0000FFFF0000,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndian(test.number)
			assert.Equal(t, test.result, result)
		})
	}

	for name, test := range tests16 {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndian(test.number)
			assert.Equal(t, test.result, result)
		})
	}

	for name, test := range tests64 {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndian(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}
