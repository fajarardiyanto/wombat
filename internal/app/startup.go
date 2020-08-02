// Copyright 2020 Rogchap. All Rights Reserved.

package app

import (
	"os"
	"path/filepath"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"

	"rogchap.com/wombat/internal/debug"
)

// The following variables are set via LDFlags at build time
var (
	appname = "Wombat"
	semver  = "0.1.0-alpha"
	isDebug = true
)

// Startup is the main startup of the application
func Startup() int {
	core.QCoreApplication_SetApplicationName(appname)
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	app := gui.NewQGuiApplication(len(os.Args), os.Args)
	app.SetWindowIcon(gui.NewQIcon5(":/qml/img/icon_128x128@2x.png"))

	engine := qml.NewQQmlApplicationEngine(nil)

	entry := "qrc:/qml/main.qml"
	if isDebug {
		entry = filepath.Join(".", "qml", "main.qml")
		debug.HotReloader(engine)
		app.SetQuitOnLastWindowClosed(false)
	}

	mc := NewMainController(nil)

	engine.RootContext().SetContextProperty("mc", mc)
	engine.Load(core.NewQUrl3(entry, 0))

	return app.Exec()
}
