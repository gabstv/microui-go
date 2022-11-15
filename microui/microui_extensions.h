#include "microui.h"
#include <stdlib.h>

#ifndef MICROUI_EXTENSIONS_H
#define MICROUI_EXTENSIONS_H

#if defined(__cplusplus)
extern "C" {            // Prevents name mangling of functions
#endif

enum {
    MUX_STYLE_INT_PADDING,
    MUX_STYLE_INT_SPACING,
    MUX_STYLE_INT_INDENT,
    MUX_STYLE_INT_TITLE_HEIGHT,
    MUX_STYLE_INT_FOOTER_HEIGHT,
    MUX_STYLE_INT_SCROLLBAR_SIZE,
    MUX_STYLE_INT_THUMB_SIZE,
    MUX_STYLE_INT_MAX
};

void mux_ensure_style_stack(mu_Context *ctx);
void mux_init_style_stack(mux_StyleStack *stack);
void mux_push_style_color(mu_Context *ctx, int colorid, mu_Color color);
void mux_push_style_font(mu_Context *ctx, mu_Font font);
void mux_push_style_size(mu_Context *ctx, mu_Vec2 size);
void mux_push_style_int(mu_Context *ctx, int styleid, int value);
void mux_pop_style(mu_Context *ctx);

#if defined(__cplusplus)
}
#endif

#endif