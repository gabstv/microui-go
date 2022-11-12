package microui

import "unsafe"

type CommandType int32

const (
	CommandJump CommandType = iota + 1
	CommandClip
	CommandRect
	CommandText
	CommandIcon
	CommandMax
)

type Command struct {
	ctype CommandType
	size  int32
	data  []byte
}

func (c *Command) Type() CommandType {
	return c.ctype
	// return *(*CommandType)(unsafe.Pointer(&c.data[0]))
}

func (c *Command) Size() int32 {
	return c.size
}

// type BaseCommand struct {
// 	Type CommandType
// 	Size int32
// }

// type JumpCommand struct {
// 	Base BaseCommand
// 	Dst unsafe.Pointer
// }

// type ClipCommand struct {
// 	Base BaseCommand
// 	Rect Rect
// }

// type RectCommand struct {
// 	Base BaseCommand
// 	Rect Rect
// 	Color Color
// }

// type TextCommand struct {
// 	Base BaseCommand
// 	Font Font
// 	Pos Vec2
// 	Color Color
// 	Text string
// }

// type BaseCommand struct {
// 	*Command
// }

// func (c *BaseCommand) Type() CommandType {
// 	return c.ctype
// }

// func (c *BaseCommand) Size() int32 {
// 	return c.size
// }

type JumpCommand struct {
	*Command
}

func (c *JumpCommand) Dst() uintptr {
	_ = c.data[7] // assert size
	return *(*uintptr)(unsafe.Pointer(&c.data[0]))
}

type ClipCommand struct {
	*Command
}

func (c *ClipCommand) Rect() Rect {
	_ = c.data[15] // assert size
	return *(*Rect)(unsafe.Pointer(&c.data[0]))
}

type RectCommand struct {
	*Command
}

func (c *RectCommand) Rect() Rect {
	_ = c.data[15] // assert size
	return *(*Rect)(unsafe.Pointer(&c.data[0]))
}

func (c *RectCommand) Color() Color {
	_ = c.data[19] // assert size
	return *(*Color)(unsafe.Pointer(&c.data[16]))
}

type TextCommand struct {
	*Command
}

func (c *TextCommand) Font() Font {
	_ = c.data[7] // assert size
	return *(*Font)(unsafe.Pointer(&c.data[0]))
}

func (c *TextCommand) Pos() Vec2 {
	_ = c.data[15] // assert size
	return *(*Vec2)(unsafe.Pointer(&c.data[8]))
}

func (c *TextCommand) Color() Color {
	_ = c.data[19] // assert size
	return *(*Color)(unsafe.Pointer(&c.data[16]))
}

func (c *TextCommand) Text() string {
	_ = c.data[20] // assert size
	return string(c.data[20:])
}

type IconCommand struct {
	*Command
}

func (c *IconCommand) Rect() Rect {
	_ = c.data[15] // assert size
	return *(*Rect)(unsafe.Pointer(&c.data[0]))
}

func (c *IconCommand) ID() int32 {
	_ = c.data[19] // assert size
	return *(*int32)(unsafe.Pointer(&c.data[16]))
}

func (c *IconCommand) Color() Color {
	_ = c.data[23] // assert size
	return *(*Color)(unsafe.Pointer(&c.data[20]))
}

func (c *Command) Jump() JumpCommand {
	return JumpCommand{c}
}

func (c *Command) Clip() ClipCommand {
	return ClipCommand{c}
}

func (c *Command) Rect() RectCommand {
	return RectCommand{c}
}

func (c *Command) Text() TextCommand {
	return TextCommand{c}
}

func (c *Command) Icon() IconCommand {
	return IconCommand{c}
}
