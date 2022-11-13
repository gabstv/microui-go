package microui

/*
#include "microui.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

type ResultFlags int32

const (
	ResActive ResultFlags = 1 << iota
	ResSubmit
	ResChange
)

const (
	FormatReal   = "%.3g"
	FormatSlider = "%.2f"
)

type ColorID int32

const (
	ColorText ColorID = iota
	ColorBorder
	ColorWindowbg
	ColorTitlebg
	ColorTitletext
	ColorPanelbg
	ColorButton
	ColorButtonhover
	ColorButtonfocus
	ColorBase
	ColorBasehover
	ColorBasefocus
	ColorScrollbase
	ColorScrollthumb
)

func (ctx *Context) DrawRect(rect Rect, color Color) {
	C.mu_draw_rect(ctx.parent, rect.cval(), color.cval())
}

// TODO: mu_draw_box
// TODO: mu_draw_text

func (ctx *Context) DrawIcon(id int32, rect Rect, color Color) {
	C.mu_draw_icon(ctx.parent, C.int(id), rect.cval(), color.cval())
}

// ...

func (ctx *Context) LayoutRow(items int32, widths []int32, height int32) {
	var cwidths *C.int
	if len(widths) > 0 {
		cwidths = (*C.int)(unsafe.Pointer(&widths[0]))
	}
	C.mu_layout_row(ctx.parent, C.int(items), cwidths, C.int(height))
}

func (ctx *Context) LayoutWidth(width int32) {
	C.mu_layout_width(ctx.parent, C.int(width))
}

func (ctx *Context) LayoutHeight(height int32) {
	C.mu_layout_height(ctx.parent, C.int(height))
}

func (ctx *Context) LayoutBeginColumn() {
	C.mu_layout_begin_column(ctx.parent)
}

func (ctx *Context) LayoutEndColumn() {
	C.mu_layout_end_column(ctx.parent)
}

func (ctx *Context) LayoutSetNext(rect Rect, relative bool) {
	C.mu_layout_set_next(ctx.parent, rect.cval(), cbool(relative))
}

func (ctx *Context) LayoutNext() Rect {
	crect := C.mu_layout_next(ctx.parent)
	return *(*Rect)(unsafe.Pointer(&crect))
}

//TODO: mu_draw_control_frame

func (ctx *Context) DrawControlText(str string, rect Rect, colorid ColorID, flags OptFlags) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	C.mu_draw_control_text(ctx.parent, cstr, rect.cval(), C.int(colorid), C.int(flags))
}

//TODO: mu_draw_control_frame...mu_begin_panel

func (ctx *Context) Text(str string) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	C.mu_text(ctx.parent, cstr)
}

func (ctx *Context) Label(str string) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	C.mu_label(ctx.parent, cstr)
}

func (ctx *Context) Button(label string) bool {
	return ctx.ButtonEx(label, 0, 0)
}

func (ctx *Context) ButtonEx(label string, icon int32, flags OptFlags) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	v := C.mu_button_ex(ctx.parent, clabel, C.int(icon), C.int(flags))
	return v != 0
}

func (ctx *Context) Checkbox(label string, state *int32) ResultFlags {
	cstate := (*C.int)(unsafe.Pointer(state))
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	v := C.mu_checkbox(ctx.parent, clabel, cstate)
	return ResultFlags(v)
}

func (ctx *Context) Textbox(buf *Buf) ResultFlags {
	return ctx.TextboxEx(buf, 0)
}

// mu_textbox_raw would go here

func (ctx *Context) TextboxEx(buf *Buf, flags OptFlags) ResultFlags {
	cbuf := (*C.char)(unsafe.Pointer(&buf.data[0]))
	bufsz := C.int(buf.Size())
	v := C.mu_textbox_ex(ctx.parent, cbuf, bufsz, C.int(flags))
	return ResultFlags(v)
}

// ...

func (ctx *Context) Slider(value *float32, low, high float32) ResultFlags {
	return ctx.SliderEx(value, low, high, 0, FormatSlider, OptAlignCenter)
}

func (ctx *Context) SliderEx(value *float32, low, high, step float32, format string, flags OptFlags) ResultFlags {
	cvalue := (*C.float)(unsafe.Pointer(value))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	v := C.mu_slider_ex(ctx.parent, cvalue, C.float(low), C.float(high), C.float(step), cformat, C.int(flags))
	return ResultFlags(v)
}

// ...

func (ctx *Context) Header(label string) bool {
	return ctx.HeaderEx(label, 0)
}

func (ctx *Context) HeaderEx(label string, flags OptFlags) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel)) // TODO: check if this doesnt break
	v := C.mu_header_ex(ctx.parent, clabel, C.int(flags))
	return v != 0
}

func (ctx *Context) BeginTreenode(label string) bool {
	return ctx.BeginTreenodeEx(label, 0)
}

func (ctx *Context) BeginTreenodeEx(label string, flags OptFlags) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	v := C.mu_begin_treenode_ex(ctx.parent, clabel, C.int(flags))
	return v != 0
}

func (ctx *Context) EndTreenode() {
	C.mu_end_treenode(ctx.parent)
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

func (ctx *Context) BeginPanel(name string) {
	ctx.BeginPanelEx(name, 0)
}

func (ctx *Context) BeginPanelEx(name string, flags OptFlags) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.mu_begin_panel_ex(ctx.parent, cname, C.int(flags))
}

func (ctx *Context) EndPanel() {
	C.mu_end_panel(ctx.parent)
}

// ...

func (ctx *Context) OpenPopup(name string) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.mu_open_popup(ctx.parent, cname)
}

func (ctx *Context) BeginPopup(name string) bool {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	v := C.mu_begin_popup(ctx.parent, cname)
	return v != 0
}

func (ctx *Context) EndPopup() {
	C.mu_end_popup(ctx.parent)
}
