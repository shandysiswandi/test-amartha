package loan

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/shandysiswandi/test-amartha/internal/loan/internal/gateway"
	"github.com/shandysiswandi/test-amartha/internal/loan/internal/interactor"
	"github.com/shandysiswandi/test-amartha/internal/pkg/pkgsql"
	"github.com/shandysiswandi/test-amartha/internal/pkg/pkguid"
	"go.uber.org/zap"
)

type Exposed struct {
}

type Dependencies struct {
	DB           *sql.DB
	Logger       *zap.SugaredLogger
	QueryBuilder pkgsql.GoquBuilder
	SnowflakeGen pkguid.Snowflake
	HttpRouter   *httprouter.Router
	Validator    *validator.Validate
}

func New(deps Dependencies) *Exposed {
	loanSQLstore := gateway.NewLoanSQLGateway(deps.DB, deps.Logger, deps.QueryBuilder)

	createProposedLoanUsecase := interactor.NewCreateProposedLoan(
		loanSQLstore,
		deps.Logger,
		deps.SnowflakeGen,
	)

	approveLoanUsecase := interactor.NewApproveLoan(
		loanSQLstore,
		deps.Logger,
	)

	investLoanUsecase := interactor.NewInvestLoan(
		loanSQLstore,
		deps.Logger,
		deps.SnowflakeGen,
	)

	disburseLoanUsecase := interactor.NewDisburseLoan(
		loanSQLstore,
		deps.Logger,
	)

	loanHTTPEndpoint := gateway.NewLoanHTTPEndpoint(
		createProposedLoanUsecase,
		approveLoanUsecase,
		investLoanUsecase,
		disburseLoanUsecase,

		deps.Logger,
		deps.Validator,
	)

	gateway.NewLoanHTTPGateway(deps.HttpRouter, deps.Logger, loanHTTPEndpoint, deps.Validator)

	return &Exposed{}
}
