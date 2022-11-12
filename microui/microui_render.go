package microui

/*
#include "microui.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

func ensureBuf(slc *[]byte, size int32) {
	for len(*slc) < int(size) {
		*slc = append(*slc, make([]byte, 1024)...)
	}
}

func (ctx *Context) SetBeginRender(fn func()) {
	ctx.beginRender = fn
}

func (ctx *Context) SetRenderCommand(fn func(cmd *Command)) {
	ctx.renderCommand = fn
}

func (ctx *Context) SetEndRender(fn func()) {
	ctx.endRender = fn
}

func (ctx *Context) Render() {
	var cmd *C.mu_Command
	gocmd := &Command{
		data: ctx.lastcmdbuf,
	}
	ctx.beginRender()
	for C.mu_next_command(ctx.parent, &cmd) == 1 {
		gocmd.ctype = CommandType(*(*int32)(unsafe.Pointer(cmd)))
		gocmd.size = *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(cmd)) + 4))
		ensureBuf(&ctx.cmdbuf, gocmd.size)
		copy(ctx.cmdbuf, (*[1 << 30]byte)(unsafe.Pointer(cmd))[:gocmd.size:gocmd.size])
		ensureBuf(&gocmd.data, gocmd.size-8)
		copy(gocmd.data, ctx.cmdbuf[8:])
		ctx.renderCommand(gocmd)
	}
	ctx.endRender()
}
