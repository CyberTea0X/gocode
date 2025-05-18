package main

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		window := new(app.Window)
		err := run(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(window *app.Window) error {
	// Создаем темную тему
	theme := material.NewTheme()
	// Onedark
	theme.Palette = material.Palette{
		Bg:         color.NRGBA{R: 0x28, G: 0x2c, B: 0x34, A: 0xff},
		Fg:         color.NRGBA{R: 0xab, G: 0xb2, B: 0xbf, A: 0xff},
		ContrastBg: color.NRGBA{R: 0x61, G: 0xaf, B: 0xef, A: 0xff},
		ContrastFg: color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff},
	}
	var ops op.Ops
	var editor widget.Editor

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			ops.Reset()
			gtx := app.NewContext(&ops, e)

			// Очистка экрана цветом фона
			paint.Fill(gtx.Ops, theme.Palette.Bg)

			// Основной макет с топбаром и заголовком
			layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					// Топбар
					return layout.Inset{Top: 10, Bottom: 10, Left: 10, Right: 10}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
							layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
								// Поле ввода с подсказкой
								editor := material.Editor(theme, &editor, "Enter search text...")
								editor.Color = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}     // Белый цвет текста
								editor.HintColor = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x80} // Белый полупрозрачный для подсказки
								editor.Editor.Alignment = text.Middle
								return editor.Layout(gtx)
							}),
						)
					})
				}),
			)

			e.Frame(gtx.Ops)
		}
	}
}
