package microui

/*
#include "microui.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

type MouseButton int32

const (
	MouseLeft MouseButton = 1 << iota
	MouseRight
	MouseMiddle
)

type Key int32

const (
	KeyShift Key = 1 << iota
	KeyCtrl
	KeyAlt
	KeyBackspace
	KeyReturn
)

func (ctx *Context) InputMouseMove(x, y int32) {
	C.mu_input_mousemove(ctx.parent, C.int(x), C.int(y))
}

func (ctx *Context) InputMouseDown(x, y int32, b MouseButton) {
	C.mu_input_mousedown(ctx.parent, C.int(x), C.int(y), C.int(b))
}

func (ctx *Context) InputMouseUp(x, y int32, b MouseButton) {
	C.mu_input_mouseup(ctx.parent, C.int(x), C.int(y), C.int(b))
}

func (ctx *Context) InputScroll(x, y int32) {
	C.mu_input_scroll(ctx.parent, C.int(x), C.int(y))
}

func (ctx *Context) InputKeyDown(k Key) {
	C.mu_input_keydown(ctx.parent, C.int(k))
}

func (ctx *Context) InputKeyUp(k Key) {
	C.mu_input_keyup(ctx.parent, C.int(k))
}

func (ctx *Context) InputText(text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	C.mu_input_text(ctx.parent, ctext)
}
