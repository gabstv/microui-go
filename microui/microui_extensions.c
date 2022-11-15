#include "microui.h"
#include "microui_extensions.h"
#include <stdio.h>

#define expect(x) do {                                               \
    if (!(x)) {                                                      \
      fprintf(stderr, "Fatal error: %s:%d: assertion '%s' failed\n", \
        __FILE__, __LINE__, #x);                                     \
      abort();                                                       \
    }                                                                \
  } while (0)

static void muxi_push_style_stack(mu_Context *ctx) {
    mux_ensure_style_stack(ctx);
    mux_StyleStack *ss = ctx->style_stack;
    if (ss->top == -1) {
        ss->top = 0;
        ss->stack[0] = *ctx->style;
        ss->original = ctx->style;
        ctx->style = &ss->stack[0];
        return;
    }
    expect(ss->top+1 < MUX_MAX_STYLE_STACK);
    ss->top++;
    ss->stack[ss->top] = ss->stack[ss->top-1];
    ctx->style = &ss->stack[ss->top];
}

static void muxi_pop_style_stack(mu_Context *ctx) {
    expect(ctx->style_stack != NULL);
    mux_StyleStack *ss = ctx->style_stack;
    expect(ss->top >= 0);
    ss->top--;
    if (ss->top == -1) {
        ctx->style = ss->original;
        return;
    }
    ctx->style = &ss->stack[ss->top];
}

void mux_ensure_style_stack(mu_Context *ctx) {
    if (ctx->style_stack == NULL) {
        ctx->style_stack = malloc(sizeof(mux_StyleStack));
        mux_init_style_stack(ctx->style_stack);
    }
}

void mux_init_style_stack(mux_StyleStack *stack) {
    stack->top = -1;
}

void mux_push_style_color(mu_Context *ctx, int colorid, mu_Color color) {
    muxi_push_style_stack(ctx);
    ctx->style->colors[colorid] = color;
}

void mux_push_style_font(mu_Context *ctx, mu_Font font) {
    muxi_push_style_stack(ctx);
    ctx->style->font = font;
}

void mux_push_style_size(mu_Context *ctx, mu_Vec2 size) {
    muxi_push_style_stack(ctx);
    ctx->style->size = size;
}

void mux_push_style_int(mu_Context *ctx, int styleid, int value) {
    expect(styleid >= 0 && styleid < MUX_STYLE_INT_MAX);
    muxi_push_style_stack(ctx);
    switch (styleid)
    {
    case MUX_STYLE_INT_PADDING:
        ctx->style->padding = value;
        break;
    case MUX_STYLE_INT_SPACING:
        ctx->style->spacing = value;
        break;
    case MUX_STYLE_INT_INDENT:
        ctx->style->indent = value;
        break;
    case MUX_STYLE_INT_TITLE_HEIGHT:
        ctx->style->title_height = value;
        break;
    case MUX_STYLE_INT_FOOTER_HEIGHT:
        ctx->style->footer_height = value;
        break;
    case MUX_STYLE_INT_SCROLLBAR_SIZE:
        ctx->style->scrollbar_size = value;
        break;
    case MUX_STYLE_INT_THUMB_SIZE:
        ctx->style->thumb_size = value;
        break;
    default:
        break;
    }
}

void mux_pop_style(mu_Context *ctx) {
    muxi_pop_style_stack(ctx);
}