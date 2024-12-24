package hitachidisplay

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var getcmdTests = []struct {
	name          string
	cmd           uint8
	expectedValue int8
}{
	{name: "cmd 1", cmd: 0b00000001, expectedValue: 1},
	{name: "cmd 2", cmd: 0b00000010, expectedValue: 2},
	{name: "cmd 2", cmd: 0b00000011, expectedValue: 2},
	{name: "cmd 3", cmd: 0b00000100, expectedValue: 3},
	{name: "cmd 3", cmd: 0b00000101, expectedValue: 3},
	{name: "cmd 4", cmd: 0b00001000, expectedValue: 4},
	{name: "cmd 5", cmd: 0b00010000, expectedValue: 5},
	{name: "cmd 6", cmd: 0b00100000, expectedValue: 6},
	{name: "cmd 7", cmd: 0b01000000, expectedValue: 7},
	{name: "cmd 8", cmd: 0b10000000, expectedValue: 8},
}

func Test_getcmd(t *testing.T) {
	for _, tt := range getcmdTests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			assert.Equal(tt.expectedValue, getcmdindex(tt.cmd))
		})
	}
}
