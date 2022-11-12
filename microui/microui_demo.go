package microui

func DrawDemoWindow(ctx *Context) {
	if ctx.BeginWindow("Demo Window", NewRect(40, 40, 300, 450)) {
		defer ctx.EndWindow()
		win := ctx.GetCurrentContainer()
		// win->rect.w = mu_max(win->rect.w, 240);
		// win->rect.h = mu_max(win->rect.h, 300);
		r := win.Rect()
		r.W = Max(r.W, 240)
		r.H = Max(r.H, 300)
		win.SetRect(r)

		// window info
		if ctx.Header("Window Info") {
			// mu_Container *win = mu_get_current_container(ctx);
			win := ctx.GetCurrentContainer()
			_ = win
			// char buf[64];
			// mu_layout_row(ctx, 2, (int[]) { 54, -1 }, 0);
			// mu_label(ctx,"Position:");
			// sprintf(buf, "%d, %d", win->rect.x, win->rect.y); mu_label(ctx, buf);
			// mu_label(ctx, "Size:");
			// sprintf(buf, "%d, %d", win->rect.w, win->rect.h); mu_label(ctx, buf);
		}
	}
}
