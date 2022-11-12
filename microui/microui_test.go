package microui_test

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/gabstv/microui-go/microui"
)

func TestInitialization(t *testing.T) {
	ctx := microui.NewContext()
	fmt.Println(ctx)
	ctx.Begin()
	ok := ctx.BeginWindow("test", microui.Rect{X: 0, Y: 0, W: 100, H: 100})
	if ok {
		ctx.EndWindow()
	}
	ctx.End()
	if microui.SizeOfCommand() != 32 {
		t.Fatal("SizeOfCommand() != 32")
	}
}

func init() {
	runtime.LockOSThread()
}
