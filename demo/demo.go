package demo

import (
	"fmt"
	"strings"
	"sync"

	//lint:ignore ST1001 this is a demo
	. "github.com/gabstv/microui-go/microui"
)

var (
	demoWindowCheckboxes = [3]int32{1, 0, 1}
	demoWindowBackground = [3]float32{90, 95, 100}
	logbuf               = new(strings.Builder)
	logbufUpdated        bool

	logSubmitBuf = NewBuf(128)
)

func writeLog(text string) {
	if logbuf.Len() > 0 {
		logbuf.WriteRune('\n')
	}
	logbuf.WriteString(text)
	logbufUpdated = true
}

func BackgroundColor() Color {
	return Color{R: uint8(demoWindowBackground[0]), G: uint8(demoWindowBackground[1]), B: uint8(demoWindowBackground[2]), A: 255}
}

func DemoWindow(ctx *Context) {
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
			ctx.Label("Position:")
			postxt := fmt.Sprintf("%d, %d", r.X, r.Y)
			ctx.Label(postxt)
			ctx.Label("Size:")
			postxt = fmt.Sprintf("%d, %d", r.W, r.H)
			ctx.Label(postxt)
		}

		// labels + buttons
		if ctx.HeaderEx("Test Buttons", OptExpanded) {
			ctx.LayoutRow(3, []int32{86, -110, -1}, 0)
			ctx.Label("Test buttons 1:")
			if ctx.Button("Button 1") {
				writeLog("Pressed button 1")
			}
			if ctx.Button("Button 2") {
				writeLog("Pressed button 2")
			}
			ctx.Label("Test buttons 2:")
			if ctx.Button("Button 3") {
				writeLog("Pressed button 3")
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
						writeLog("Pressed button 1")
					}
					if ctx.Button("Button 2") {
						writeLog("Pressed button 2")
					}
					ctx.EndTreenode()
				}
				ctx.EndTreenode()
			}
			if ctx.BeginTreenode("Test 2") {
				ctx.LayoutRow(2, []int32{54, 54}, 0)
				if ctx.Button("Button 3") {
					writeLog("Pressed button 3")
				}
				if ctx.Button("Button 4") {
					writeLog("Pressed button 4")
				}
				if ctx.Button("Button 5") {
					writeLog("Pressed button 5")
				}
				if ctx.Button("Button 6") {
					writeLog("Pressed button 6")
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
			ctx.DrawRect(rect, NewColor(uint8(demoWindowBackground[0]), uint8(demoWindowBackground[1]), uint8(demoWindowBackground[2]), 255))
			txt := fmt.Sprintf("#%02X%02X%02X", uint8(demoWindowBackground[0]), uint8(demoWindowBackground[1]), uint8(demoWindowBackground[2]))
			ctx.DrawControlText(txt, rect, ColorText, OptAlignCenter)
		}
	}
}

func LogWindow(ctx *Context) {
	if ctx.BeginWindow("Log Window", NewRect(350, 40, 300, 200)) {
		defer ctx.EndWindow()
		// output text panel
		ctx.LayoutRow(1, []int32{-1}, -25)
		ctx.BeginPanel("Log Output")
		panel := ctx.GetCurrentContainer()
		ctx.LayoutRow(1, []int32{-1}, -1)
		ctx.Text(logbuf.String())
		ctx.EndPanel()
		if logbufUpdated {
			s := panel.Scroll()
			s.Y = panel.ContentSize().Y
			panel.SetScroll(s)
			logbufUpdated = false
		}

		// input textbox + submit button
		submitted := false
		ctx.LayoutRow(2, []int32{-70, -1}, 0)
		if ctx.Textbox(logSubmitBuf)&ResSubmit != 0 {
			ctx.SetFocus(ctx.LastID())
			submitted = true
		}
		if ctx.Button("Submit") {
			submitted = true
		}
		if submitted {
			writeLog(logSubmitBuf.String())
			logSubmitBuf.Clear()
		}
	}
}

func uint8Slider(ctx *Context, fvalue *float32, value *uint8, low, high float32) ResultFlags {
	*fvalue = float32(*value)
	res := ctx.SliderEx(fvalue, low, high, 0, "%.0f", OptAlignCenter)
	*value = uint8(*fvalue)
	return res
}

var colors = [...]struct {
	label   string
	colorid ColorID
}{
	{"text:", ColorText},
	{"border:", ColorBorder},
	{"windowbg:", ColorWindowbg},
	{"titlebg:", ColorTitlebg},
	{"titletext:", ColorTitletext},
	{"panelbg:", ColorPanelbg},
	{"button:", ColorButton},
	{"buttonhover:", ColorButtonhover},
	{"buttonfocus:", ColorButtonfocus},
	{"base:", ColorBase},
	{"basehover:", ColorBasehover},
	{"basefocus:", ColorBasefocus},
	{"scrollbase:", ColorScrollbase},
	{"scrollthumb:", ColorScrollthumb},
	{},
}

var pcolors [14]Color
var fcolors [14]struct {
	r, g, b, a float32
}

var styleWindowOnce sync.Once

func styleWindowOnceFn(ctx *Context) {
	s := ctx.Style()
	for i := range pcolors {
		pcolors[i] = s.Color(colors[i].colorid)
	}
}

func StyleWindow(ctx *Context) {
	styleWindowOnce.Do(func() { styleWindowOnceFn(ctx) })
	if ctx.BeginWindow("Style Window", NewRect(350, 250, 300, 240)) {
		defer ctx.EndWindow()
		sw := int32(float32(ctx.GetCurrentContainer().Body().W) * 0.14)
		ctx.LayoutRow(6, []int32{80, sw, sw, sw, sw, -1}, 0)
		for i := 0; colors[i].label != ""; i++ {
			prevc := pcolors[i]
			cptr := &pcolors[i]
			fptr := &fcolors[i]
			ctx.Label(colors[i].label)
			uint8Slider(ctx, &fptr.r, &cptr.R, 0, 255)
			uint8Slider(ctx, &fptr.g, &cptr.G, 0, 255)
			uint8Slider(ctx, &fptr.b, &cptr.B, 0, 255)
			uint8Slider(ctx, &fptr.a, &cptr.A, 0, 255)
			ctx.DrawRect(ctx.LayoutNext(), pcolors[i])
			if !prevc.Equals(pcolors[i]) {
				ctx.Style().SetColor(colors[i].colorid, pcolors[i])
			}
		}
	}
}

// 	  int sw = mu_get_current_container(ctx)->body.w * 0.14;
// 	  mu_layout_row(ctx, 6, (int[]) { 80, sw, sw, sw, sw, -1 }, 0);
// 	  for (int i = 0; colors[i].label; i++) {
// 		mu_label(ctx, colors[i].label);
// 		uint8_slider(ctx, &ctx->style->colors[i].r, 0, 255);
// 		uint8_slider(ctx, &ctx->style->colors[i].g, 0, 255);
// 		uint8_slider(ctx, &ctx->style->colors[i].b, 0, 255);
// 		uint8_slider(ctx, &ctx->style->colors[i].a, 0, 255);
// 		mu_draw_rect(ctx, mu_layout_next(ctx), ctx->style->colors[i]);
// 	  }
