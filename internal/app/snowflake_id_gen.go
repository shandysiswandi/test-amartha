package app

import (
	"github.com/hashicorp/go-multierror"
	"github.com/shandysiswandi/test-amartha/internal/pkg/pkguid"
)

func (app *App) initSnowflakeGen() {
	snowflakeGen, err := pkguid.NewSnowflake()
	if err != nil {
		app.err = multierror.Append(app.err, err)
	}

	app.snowflakeGen = snowflakeGen
}
