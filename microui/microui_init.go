package microui

/*
#include "microui.h"
#include "microui_init.h"
#include <stdlib.h>
*/
import "C"
import (
	"unsafe"
)

type GetTextWidthFunc func(font Font, text string) int32
type GetTextHeightFunc func(font Font) int32

var (
	DefaultGetTextWidth GetTextWidthFunc = func(font Font, text string) int32 {
		println("DefaultGetTextWidth", font, text)
		//TODO!
		return 1
	}
	DefaultGetTextHeight GetTextHeightFunc = func(font Font) int32 {
		println("DefaultGetTextHeight", font)
		//TODO!
		return 1
	}
)

//export goDefaultGetTextWidthFunc
func goDefaultGetTextWidthFunc(font unsafe.Pointer, text unsafe.Pointer, size C.int) C.int {
	// fmt.Println("goDefaultGetTextWidthFunc", font, text, size)
	var txt string
	if size == -1 {
		txt = C.GoString((*C.char)(text))
	} else {
		txt = string(C.GoBytes(text, size))
	}
	// fmt.Println("goDefaultGetTextWidthFunc :)", txt)
	v := DefaultGetTextWidth((Font)(font), txt)
	// fmt.Println(v)
	return C.int(v)
}

//export goDefaultGetTextHeightFunc
func goDefaultGetTextHeightFunc(font unsafe.Pointer) C.int {
	v := DefaultGetTextHeight((Font)(font))
	return C.int(v)
}
