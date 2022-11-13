package microui

import "fmt"

var (
	demoWindowCheckboxes = [3]int32{1, 0, 1}
	demoWindowBackground = [3]float32{90, 95, 100}
)

func DemoWindowBackgroundColor() Color {
	return Color{R: uint8(demoWindowBackground[0]), G: uint8(demoWindowBackground[1]), B: uint8(demoWindowBackground[2]), A: 255}
}

func DrawDemoWindow(ctx *Context) {
	if ctx.BeginWindow("Demo Window", NewRect(40, 40, 300, 450)) {
		defer ctx.EndWindow()
		win := ctx.GetCurrentContainer()
		r := win.Rect()
		r.W = Max(r.W, 240)
		r.H = Max(r.H, 300)
		win.SetRect(r)

		// window info
		if ctx.Header("Window Info") {
			win := ctx.GetCurrentContainer()
			r := win.Rect()
			ctx.LayoutRow(2, []int32{54, -1}, 0)
			// mu_label(ctx,"Position:");
			ctx.Label("Position:")
			postxt := fmt.Sprintf("%d, %d", r.X, r.Y)
			// sprintf(buf, "%d, %d", win->rect.x, win->rect.y); mu_label(ctx, buf);
			ctx.Label(postxt)
			// mu_label(ctx, "Size:");
			ctx.Label("Size:")
			postxt = fmt.Sprintf("%d, %d", r.W, r.H)
			ctx.Label(postxt)
		}

		// labels + buttons
		if ctx.HeaderEx("Test Buttons", OptExpanded) {
			ctx.LayoutRow(3, []int32{86, -110, -1}, 0)
			ctx.Label("Test buttons 1:")
			if ctx.Button("Button 1") {
				// TODO: write_log("Pressed button 1")
			}
			if ctx.Button("Button 2") {
				// TODO: write_log("Pressed button 2")
			}
			ctx.Label("Test buttons 2:")
			if ctx.Button("Button 3") {
				// TODO: write_log("Pressed button 3")
			}
			if ctx.Button("Popup") {
				ctx.OpenPopup("Test Popup")
			}
			if ctx.BeginPopup("Test Popup") {
				ctx.Button("Hello")
				ctx.Button("World")
				ctx.EndPopup()
			}
		}

		// tree
		if ctx.HeaderEx("Tree and Text", OptExpanded) {
			ctx.LayoutRow(2, []int32{140, -1}, 0)
			ctx.LayoutBeginColumn()
			if ctx.BeginTreenode("Test 1") {
				if ctx.BeginTreenode("Test 1a") {
					ctx.Label("Hello")
					ctx.Label("World")
					ctx.EndTreenode()
				}
				if ctx.BeginTreenode("Test 1b") {
					if ctx.Button("Button 1") {
						//TODO: write_log("Pressed button 1")
					}
					if ctx.Button("Button 2") {
						//TODO: write_log("Pressed button 2")
					}
					ctx.EndTreenode()
				}
				ctx.EndTreenode()
			}
			if ctx.BeginTreenode("Test 2") {
				ctx.LayoutRow(2, []int32{54, 54}, 0)
				if ctx.Button("Button 3") {
					//TODO: write_log("Pressed button 3")
				}
				if ctx.Button("Button 4") {
					//TODO: write_log("Pressed button 3")
				}
				if ctx.Button("Button 5") {
					//TODO: write_log("Pressed button 3")
				}
				if ctx.Button("Button 6") {
					//TODO: write_log("Pressed button 3")
				}
				ctx.EndTreenode()
			}
			if ctx.BeginTreenode("Test 3") {
				ctx.Checkbox("Checkbox 1", &demoWindowCheckboxes[0])
				ctx.Checkbox("Checkbox 2", &demoWindowCheckboxes[1])
				ctx.Checkbox("Checkbox 3", &demoWindowCheckboxes[2])
				ctx.EndTreenode()
			}
			ctx.LayoutEndColumn()

			ctx.LayoutBeginColumn()
			ctx.LayoutRow(1, []int32{-1}, 0)
			ctx.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas lacinia, sem eu lacinia molestie, mi risus faucibus ipsum, eu varius magna felis a nulla.")
			ctx.LayoutEndColumn()
		}

		// background color sliders
		if ctx.HeaderEx("Background Color", OptExpanded) {
			ctx.LayoutRow(2, []int32{-78, -1}, 74)
			//  sliders
			ctx.LayoutBeginColumn()
			ctx.LayoutRow(2, []int32{46, -1}, 0)
			ctx.Label("Red:")
			ctx.Slider(&demoWindowBackground[0], 0, 255)
			ctx.Label("Green:")
			ctx.Slider(&demoWindowBackground[1], 0, 255)
			ctx.Label("Blue:")
			ctx.Slider(&demoWindowBackground[2], 0, 255)
			ctx.LayoutEndColumn()
			// color preview
			rect := ctx.LayoutNext()
			ctx.DrawRect(rect, Color{uint8(demoWindowBackground[0]), uint8(demoWindowBackground[1]), uint8(demoWindowBackground[2]), 255})
			txt := fmt.Sprintf("#%02X%02X%02X", uint8(demoWindowBackground[0]), uint8(demoWindowBackground[1]), uint8(demoWindowBackground[2]))
			ctx.DrawControlText(txt, rect, ColorText, OptAlignCenter)
		}
	}
}
