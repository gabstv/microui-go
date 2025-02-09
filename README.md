# ![microui-go](https://user-images.githubusercontent.com/3920290/75171571-be83c500-5723-11ea-8a50-504cc2ae1109.png)
A Go binding of [microui](https://github.com/rxi/microui), a *tiny*, portable, immediate-mode UI library written in ANSI C

## Features
* Tiny: around `1100 sloc` of ANSI C
* Works within a fixed-sized memory region: no additional memory is allocated
* Built-in controls: window, scrollable panel, button, slider, textbox, label,
  checkbox, wordwrapped text
* Works with any rendering system that can draw rectangles and text
* Designed to allow the user to easily add custom controls
* Simple layout system

## Example
![example](docs/example.png)

This example is using a [raylib-go driver](https://github.com/gabstv/microui-go-raylib). The source code is available at:  
https://github.com/gabstv/microui-go-raylib/blob/main/example/demo/main.go

### Ebitengine
There is also a [microui-go-ebitengine](https://github.com/gabstv/microui-go-ebitengine) driver, which uses the Ebitengine library to render.

## Usage
A Go version of the usage docs is not ready yet. In the meantime, you can check the
[original usage docs](https://github.com/rxi/microui/blob/master/doc/usage.md) at the
rxi repository (and use the equivalent functions in Go).

## Notes
The library expects the user to provide input and handle the resultant drawing
commands, it does not do any drawing itself.

## Contributing
The library is designed to be lightweight, providing a foundation to which you
can easily add custom controls and UI elements; pull requests adding additional
features will likely not be merged. Bug reports are welcome.

## License
This library is free software; you can redistribute it and/or modify it under
the terms of the MIT license. See [LICENSE](LICENSE.txt) for details.