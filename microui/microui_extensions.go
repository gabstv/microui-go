package microui

/*
#include "microui.h"
#include "microui_extensions.h"
#include <stdlib.h>
*/
import "C"

type IntStyles int32

const (
	StylePadding IntStyles = iota
	StyleSpacing
	StyleIndent
	StyleTitleHeight
	StyleFooterHeight
	StyleScrollbarSize
	StyleThumbSize
)

func (c *Context) PushStyleColor(id ColorID, color Color) {
	C.mux_push_style_color(c.parent, C.int(id), color.cval())
}

func (c *Context) PushStyleFont(font Font) {
	C.mux_push_style_font(c.parent, cvalOfFont(font))
}

func (c *Context) PushStyleSize(size Vec2) {
	C.mux_push_style_size(c.parent, size.cval())
}

func (c *Context) PushStyleInt(id IntStyles, value int32) {
	C.mux_push_style_int(c.parent, C.int(id), C.int(value))
}

func (c *Context) PopStyle() {
	C.mux_pop_style(c.parent)
}
