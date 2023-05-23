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

type Context struct {
	parent     *C.mu_Context
	cmdbuf     []byte
	lastcmdbuf []byte

	beginRender   func()
	renderCommand func(cmd *Command)
	endRender     func()

	beginCallback func()
	endCallback   func()
}

func NewContext() *Context {
	return &Context{
		parent:      C.NewContext(),
		cmdbuf:      make([]byte, 1024*32),
		beginRender: func() {},
		renderCommand: func(cmd *Command) {
			panic("renderCommand not set")
		},
		endRender: func() {},
	}
}

func (ctx *Context) SetBeginCallback(fn func()) {
	ctx.beginCallback = fn
}

func (ctx *Context) SetEndCallback(fn func()) {
	ctx.endCallback = fn
}

func (ctx *Context) Begin() {
	C.mu_begin(ctx.parent)
	if ctx.beginCallback != nil {
		ctx.beginCallback()
	}
}

func (ctx *Context) End() {
	C.mu_end(ctx.parent)
	if ctx.endCallback != nil {
		ctx.endCallback()
	}
}

// Focus returns the id of the currently focused item.
func (ctx *Context) Focus() ID {
	return ID(ctx.parent.focus)
}

// Hover returns the id of the currently hovered item.
func (ctx *Context) Hover() ID {
	return ID(ctx.parent.hover)
}

// HoverRoot returns the container that is currently hovered.
// This is useful for determining if the mouse is hovering over a window.
func (ctx *Context) HoverRoot() *Container {
	if ctx.parent.hover_root == nil {
		return nil
	}
	return &Container{
		parent: ctx.parent.hover_root,
	}
}

func (ctx *Context) SetFocus(id ID) {
	C.mu_set_focus(ctx.parent, C.uint(id))
}

func GetID[T any](ctx *Context, data *T) ID {
	size := unsafe.Sizeof(*data)
	return ID(C.mu_get_id(ctx.parent, unsafe.Pointer(data), C.int(size)))
}

type AddressableByC interface {
	constraints.Integer | constraints.Float
}

func PushID[T AddressableByC](ctx *Context, data *T) {
	size := unsafe.Sizeof(*data)
	C.mu_push_id(ctx.parent, unsafe.Pointer(data), C.int(size))
}

func (ctx *Context) PushIDInt32(id *int32) {
	C.mu_push_id(ctx.parent, unsafe.Pointer(id), C.int(4))
}

func (ctx *Context) PopID() {
	C.mu_pop_id(ctx.parent)
}

func (ctx *Context) LastID() ID {
	return ID(ctx.parent.last_id)
}

func (ctx *Context) GetCurrentContainer() *Container {
	return &Container{
		parent: C.mu_get_current_container(ctx.parent),
	}
}

func (ctx *Context) GetContainer(name string) *Container {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	container := C.mu_get_container(ctx.parent, cname)
	if container == nil {
		return nil
	}
	return &Container{
		parent: container,
	}
}

func (ctx *Context) Style() *Style {
	return &Style{
		parent: ctx.parent.style,
	}
}

type Style struct {
	parent *C.mu_Style
}

func (s *Style) Color(id ColorID) Color {
	return *(*Color)(unsafe.Pointer(&s.parent.colors[int32(id)]))
}

func (s *Style) SetColor(id ColorID, color Color) {
	s.parent.colors[int32(id)] = *(*C.mu_Color)(unsafe.Pointer(&color))
}
