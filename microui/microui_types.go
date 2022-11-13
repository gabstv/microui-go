package microui

/*
#include "microui.h"
#include "microui_init.h"
#include <stdlib.h>
*/
import "C"
import (
	"strings"
	"unsafe"

	"golang.org/x/exp/constraints"
)

type OptFlags int32

const (
	OptAlignCenter OptFlags = 1 << iota
	OptAlignRight
	OptNoInteract
	OptNoFrame
	OptNoResize
	OptNoScroll
	OptNoClose
	OptNoTitle
	OptHoldFocus
	OptAutosize
	OptPopup
	OptClosed
	OptExpanded
)

// typedef struct { int x, y, w, h; } mu_Rect;
type Rect struct {
	X int32
	Y int32
	W int32
	H int32
}

func (r *Rect) cval() C.mu_Rect {
	return C.mu_Rect{C.int(r.X), C.int(r.Y), C.int(r.W), C.int(r.H)}
}

func NewRect(x, y, w, h int32) Rect {
	return Rect{X: x, Y: y, W: w, H: h}
}

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func (r *Color) cval() C.mu_Color {
	return C.mu_Color{C.uchar(r.R), C.uchar(r.G), C.uchar(r.B), C.uchar(r.A)}
}

func NewColor(r, g, b, a uint8) Color {
	return Color{R: r, G: g, B: b, A: a}
}

type Vec2 struct {
	X int32
	Y int32
}

func (r *Vec2) cval() C.mu_Vec2 {
	return C.mu_Vec2{C.int(r.X), C.int(r.Y)}
}

type Container struct {
	parent *C.mu_Container
}

func (c *Container) Rect() Rect {
	v := *(*Rect)(unsafe.Pointer(&c.parent.rect))
	return v
}

func (c *Container) SetRect(r Rect) {
	c.parent.rect = r.cval()
}

func (c *Container) Body() Rect {
	v := *(*Rect)(unsafe.Pointer(&c.parent.body))
	return v
}

func (c *Container) SetBody(r Rect) {
	c.parent.body = r.cval()
}

func (c *Container) ContentSize() Vec2 {
	v := *(*Vec2)(unsafe.Pointer(&c.parent.content_size))
	return v
}

func (c *Container) Scroll() Vec2 {
	v := *(*Vec2)(unsafe.Pointer(&c.parent.scroll))
	return v
}

func (c *Container) SetScroll(v Vec2) {
	c.parent.scroll = v.cval()
}

func (c *Container) ZIndex() int32 {
	v := *(*int32)(unsafe.Pointer(&c.parent.zindex))
	return v
}

func (c *Container) Open() bool {
	v := *(*int32)(unsafe.Pointer(&c.parent.open))
	return v == 1
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

type Font unsafe.Pointer

func cbool(v bool) C.int {
	if v {
		return 1
	}
	return 0
}

type Buf struct {
	data []byte
}

func NewBuf(size int) *Buf {
	return &Buf{
		data: make([]byte, size+1),
	}
}

func (sb *Buf) Size() int {
	return len(sb.data)
}

func (sb *Buf) String() string {
	nb := strings.IndexByte(string(sb.data), 0)
	if nb == -1 || nb == 0 {
		return ""
	}
	return string(sb.data[:nb])
}

// Clear sets the first byte of the buffer to 0.
func (sb *Buf) Clear() {
	sb.data[0] = 0
}

func (sb *Buf) SetString(s string) error {
	slen := len(s)
	if slen > len(sb.data)-1 {
		return ErrBufferTooSmall
	}
	// erase the buffer
	for i := range sb.data {
		sb.data[i] = 0
	}
	// copy the string
	copy(sb.data, []byte(s))
	// ensure null termination
	// sb.data[slen] = 0
	return nil
}

type ID uint32
