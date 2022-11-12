package microui

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func tobytes[T any](data T) []byte {
	if str, ok := any(data).(string); ok {
		return []byte(str)
	}
	if b, ok := any(data).([]byte); ok {
		return b
	}
	sz := unsafe.Sizeof(data)
	bb := make([]byte, sz)
	for i := 0; i < int(sz); i++ {
		bb[i] = *(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&data)) + uintptr(i)))
	}
	return bb
}

func newslicecp(src []byte, sz int) []byte {
	dst := make([]byte, sz)
	copy(dst, src[:sz])
	return dst
}

func TestCommandConversions(t *testing.T) {
	var cmd Command
	buf := make([]byte, 1024)

	// test jump command
	cmd.ctype = CommandJump
	cmd.size = 16
	copy(buf, tobytes(uintptr(0x12345678)))
	cmd.data = newslicecp(buf, 16) // should be 8, but were testing that a bigger buf should not matter
	j := cmd.Jump()
	assert.Equal(t, int(CommandJump), int(j.Type()))
	assert.Equal(t, int32(16), j.Size())
	assert.Equal(t, uintptr(0x12345678), j.Dst())

	// test clip command
	cmd.ctype = CommandClip
	cmd.size = 24
	copy(buf, tobytes(Rect{1, 2, 3, 4}))
	cmd.data = newslicecp(buf, 16)
	c := cmd.Clip()
	assert.Equal(t, Rect{1, 2, 3, 4}, c.Rect())

	// test rect command
	cmd.ctype = CommandRect
	copy(buf, tobytes(Rect{10, 20, 30, 40}))
	copy(buf[16:], tobytes(Color{1, 2, 3, 4}))
	cmd.data = newslicecp(buf, 24)
	cmd.size = int32(unsafe.Sizeof(Rect{})+unsafe.Sizeof(Color{})) + 8
	r := cmd.Rect()
	assert.Equal(t, Rect{10, 20, 30, 40}, r.Rect())
	assert.Equal(t, Color{1, 2, 3, 4}, r.Color())

	// test text command
	cmd.ctype = CommandText
	copy(buf, tobytes(Font(0x12345678)))
	copy(buf[8:], tobytes(Vec2{1, 2}))
	copy(buf[16:], tobytes(Color{3, 4, 5, 6}))
	copy(buf[20:], tobytes("hello world"))
	cmd.data = newslicecp(buf, 31)
	cmd.size = 39
	tt := cmd.Text()
	assert.Equal(t, Font(0x12345678), tt.Font())
	assert.Equal(t, Vec2{1, 2}, tt.Pos())
	assert.Equal(t, Color{3, 4, 5, 6}, tt.Color())
	assert.Equal(t, "hello world", tt.Text())

	// test icon command
	cmd.ctype = CommandIcon
	copy(buf, tobytes(Rect{11, 22, 33, 44}))
	copy(buf[16:], tobytes(int32(12)))
	copy(buf[20:], tobytes(Color{0xff, 0x00, 0xaf, 0xFF}))
	cmd.data = newslicecp(buf, 24)
	cmd.size = 32
	i := cmd.Icon()
	assert.Equal(t, Rect{11, 22, 33, 44}, i.Rect())
	assert.Equal(t, int32(12), i.ID())
	assert.Equal(t, Color{0xff, 0x00, 0xaf, 0xFF}, i.Color())
}
