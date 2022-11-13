package microui

/*
#include "microui.h"
#include "microui_init.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

type Context struct {
	parent     *C.mu_Context
	cmdbuf     []byte
	lastcmdbuf []byte

	beginRender   func()
	renderCommand func(cmd *Command)
	endRender     func()
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

func (ctx *Context) Begin() {
	C.mu_begin(ctx.parent)
}

func (ctx *Context) End() {
	C.mu_end(ctx.parent)
}

// Focus returns the id of the currently focused item.
func (ctx *Context) Focus() ID {
	return ID(ctx.parent.focus)
}

// Hover returns the id of the currently hovered item.
func (ctx *Context) Hover() ID {
	return ID(ctx.parent.hover)
}

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

func PushID[T any](ctx *Context, data *T) {
	size := unsafe.Sizeof(*data)
	C.mu_push_id(ctx.parent, unsafe.Pointer(data), C.int(size))
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
