package microui

/*
#include "microui.h"
#include "microui_init.h"
#include <stdlib.h>
*/
import "C"
import (
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

type Vec2 struct {
	X int32
	Y int32
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
