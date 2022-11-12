#include "microui.h"
#include "microui_init.h"

static int defaultGetTextWidthFromGo(mu_Font font, const char *text, int len0) {
    return goDefaultGetTextWidthFunc(font, text, len0);
}

static int defaultGetTextHeightFromGo(mu_Font font) {
    return goDefaultGetTextHeightFunc(font);
}

mu_Context* NewContext(void) {
    mu_Context *ctx = malloc(sizeof(mu_Context));
    mu_init(ctx);
    ctx->text_width = defaultGetTextWidthFromGo;
    ctx->text_height = defaultGetTextHeightFromGo;
    return ctx;
}

void SetTextWidthFunc(mu_Context* ctx, text_width_fn fn) {
    ctx->text_width = fn;
}