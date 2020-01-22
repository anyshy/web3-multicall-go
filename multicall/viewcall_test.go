package multicall

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestViewCall(t *testing.T) {
	vc := viewCall{
		id:        "key",
		target:    "0x0",
		method:    "balanceOf(address, uint64)(int256)",
		arguments: []interface{}{"0x1234", uint64(12)},
	}
	expectedArgTypes := []string{"address", "uint64"}
	expectedCallData := []byte{
		0x29, 0x5e, 0xaa, 0xdf, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x12, 0x34, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0xc}
	assert.Equal(t, expectedArgTypes, vc.argumentTypes())
	callData, err := vc.callData()
	assert.Nil(t, err)
	assert.Equal(t, expectedCallData, callData)
}
