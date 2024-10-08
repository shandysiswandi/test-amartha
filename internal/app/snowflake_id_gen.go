package app

import (
	"errors"

	"github.com/shandysiswandi/test-amartha/internal/pkg/pkguid"
)

func (app *App) initSnowflakeGen() {
	snowflakeGen, err := pkguid.NewSnowflake()
	if err != nil {
		app.err = errors.Join(app.err, err)
	}

	app.snowflakeGen = snowflakeGen
}
