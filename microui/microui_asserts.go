package microui

/*
#include "microui.h"
#include "microui_asserts.h"
#include <stdlib.h>
*/
import "C"

func SizeOfCommand() int32 {
	sz := C.SizeOfCommand()
	return int32(sz)
}
