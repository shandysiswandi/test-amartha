package app

import (
	"errors"

	"go.uber.org/zap"
)

func (app *App) initLogger() {
	logger, err := zap.NewDevelopment(zap.AddStacktrace(zap.FatalLevel))
	if err != nil {
		app.err = errors.Join(app.err, err)
	}

	app.logger = logger
}
