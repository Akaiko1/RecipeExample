package main

import (
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"os"
)

type SettingsWindow struct {
	window  *widgets.QWidget
	buttons []*widgets.QPushButton
}

type MainWindow struct {
	window  *widgets.QMainWindow
	overlay *widgets.QLabel
}

func (qw *SettingsWindow) Init() {
	qw.window = widgets.NewQWidget(nil, 0)
	qw.window.SetWindowTitle("Settings")
	qw.window.SetMinimumSize2(200, 100)
	qw.window.Move2(700, 200)
}

func (qw *SettingsWindow) appendButtons(buttons []*widgets.QPushButton) {
	qw.buttons = append(qw.buttons, buttons...)

	for _, button := range qw.buttons {
		qw.window.Layout().AddWidget(button)
	}
}

func (mw *MainWindow) Init() {
	mw.window = widgets.NewQMainWindow(nil, 0)
	mw.window.SetWindowTitle("Editor")
	mw.window.SetMinimumSize2(500, 500)
	mw.window.Move2(200, 200)

	mw.overlay = widgets.NewQLabel(mw.window, 0)
	mw.overlay.SetFixedSize(mw.window.Size())

	mw.window.ConnectPaintEvent(mw.PaintSettings)
}

func (mw *MainWindow) PaintSettings(event *gui.QPaintEvent)  {
	painter := gui.NewQPainter2(mw.window)
	painter.DrawPixmap10(mw.window.Rect(), gui.NewQPixmap3("tray.png", "png", 0))
	painter.SetOpacity(0.25)
	painter.DrawPixmap10(mw.overlay.Rect(), gui.NewQPixmap3("circles.bmp", "bmp", 0))
	painter.DestroyQPainter()
	mw.overlay.SetFixedSize(mw.window.Size())
}

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	var editor MainWindow
	var settings SettingsWindow

	editor.Init()
	settings.Init()
	settings.appendButtons([]*widgets.QPushButton{widgets.NewQPushButton(settings.window)})

	editor.window.Show()
	settings.window.Show()

	widgets.QApplication_Exec()
}
