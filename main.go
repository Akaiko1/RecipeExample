package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"io/ioutil"
	"log"
	"os"
)

type SettingsWindow struct {
	window  *widgets.QWidget
	buttons []*widgets.QPushButton
	sliders []*widgets.QSlider
}

type MainWindow struct {
	window  *widgets.QMainWindow
	overlay *widgets.QLabel
	transparency float64
}

func (qw *SettingsWindow) Init() {
	qw.window = widgets.NewQWidget(nil, 0)
	qw.window.SetWindowTitle("Settings")
	qw.window.SetMinimumSize2(300, 100)
	qw.window.SetLayout(widgets.NewQVBoxLayout())
	qw.setStyleSheet()

	qw.window.Move2(700, 200)
}

func (mw *MainWindow) Init() {
	mw.window = widgets.NewQMainWindow(nil, 0)
	mw.window.SetWindowTitle("Editor")
	mw.window.SetMinimumSize2(500, 500)
	mw.window.Move2(200, 200)
	mw.overlay = widgets.NewQLabel(mw.window, 0)
	mw.overlay.SetFixedSize(mw.window.Size())
	mw.transparency = 0.25

	mw.window.ConnectPaintEvent(mw.PaintSettings)
	mw.window.ConnectCloseEvent(mw.closeMain)
}

func (qw *SettingsWindow) appendControls(buttons []*widgets.QPushButton, sliders []*widgets.QSlider) {
	qw.buttons = append(qw.buttons, buttons...)
	qw.sliders = append(qw.sliders, sliders...)

	buttonsLayout := widgets.NewQHBoxLayout()
	for _, button := range qw.buttons {
		buttonsLayout.AddWidget(button, 0, core.Qt__AlignCenter)
	}
	qw.window.Layout().AddItem(buttonsLayout)

	for _, slider := range qw.sliders {
		qw.window.Layout().AddWidget(slider)
	}
}

func (qw *SettingsWindow) setStyleSheet() {
	content, err := ioutil.ReadFile("settings.qss")
	if err != nil {
		log.Fatal(err)
	}
	qw.window.SetStyleSheet(string(content))
}

func (mw *MainWindow) PaintSettings(event *gui.QPaintEvent)  {
	painter := gui.NewQPainter2(mw.window)
	painter.DrawPixmap10(mw.window.Rect(), gui.NewQPixmap3("tray.png", "png", 0))
	painter.SetOpacity(mw.transparency)
	painter.DrawPixmap10(mw.overlay.Rect(), gui.NewQPixmap3("circles.bmp", "bmp", 0))
	painter.DestroyQPainter()
	mw.overlay.SetFixedSize(mw.window.Size())
}

func (mw *MainWindow) closeMain(event *gui.QCloseEvent)  {
	settings.window.Close()
}

func transparencyChanged(value int) {
	editor.transparency = float64(value)/100.
	editor.window.Repaint()
}

var (
	editor MainWindow;
	settings SettingsWindow
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	editor.Init()
	settings.Init()

	sliderTransparency := widgets.NewQSlider2(core.Qt__Horizontal, settings.window)
	sliderTransparency.SetValue(25)
	sliderTransparency.ConnectSliderMoved(transparencyChanged)

	settings.appendControls([]*widgets.QPushButton{
		widgets.NewQPushButton2("Press", settings.window),
		widgets.NewQPushButton2("Me", settings.window),
	},
	[]*widgets.QSlider{
		sliderTransparency,
	},
	)

	editor.window.Show()
	settings.window.Show()

	widgets.QApplication_Exec()
}
