package microui

/*
#include "microui.h"
#include "microui_init.h"
#include <stdlib.h>
*/
import "C"
import (
	"unsafe"
)

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

func (ctx *Context) BeginWindow(title string, rect Rect) bool {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	iv := C.mu_begin_window_ex(ctx.parent, cTitle, rect.cval(), C.int(0))
	return iv != 0
}

func (ctx *Context) EndWindow() {
	C.mu_end_window(ctx.parent)
}

func (ctx *Context) Begin() {
	C.mu_begin(ctx.parent)
}

func (ctx *Context) End() {
	C.mu_end(ctx.parent)
}

func (ctx *Context) GetCurrentContainer() *Container {
	return &Container{
		parent: C.mu_get_current_container(ctx.parent),
	}
}

func (ctx *Context) Header(label string) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel)) // TODO: check if this doesnt break
	v := C.mu_header_ex(ctx.parent, clabel, 0)
	return v != 0
}
