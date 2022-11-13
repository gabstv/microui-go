package microui

/*
#include "microui.h"
#include "microui_init.h"
#include <stdlib.h>
*/
import "C"

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

func (ctx *Context) SetFocus(id ID) {
	C.mu_set_focus(ctx.parent, C.uint(id))
}

func (ctx *Context) LastID() ID {
	return ID(ctx.parent.last_id)
}

func (ctx *Context) GetCurrentContainer() *Container {
	return &Container{
		parent: C.mu_get_current_container(ctx.parent),
	}
}
