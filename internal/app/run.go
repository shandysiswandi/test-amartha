package app

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (app *App) run() {
	app.spinUp()

	terminateChan := app.createTerminateSignal()

	if err := app.Start(); app.err != nil || err != nil {
		app.logger.Sugar().Panicw("failed to start application", "error", err)
	}

	<-terminateChan
}

func (app *App) createTerminateSignal() <-chan struct{} {
	var terminateChan = make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		sigVal := <-sigint

		if err := app.Stop(ctx); err != nil {
			app.err = errors.Join(app.err, err)
		}

		switch sigVal {
		case syscall.SIGHUP:
			app.logger.Sugar().Info("application reloading...")

			app.run()
		default:
			terminateChan <- struct{}{}
			close(terminateChan)

			app.logger.Sugar().Info("application terminated")
		}

	}()

	return terminateChan
}
