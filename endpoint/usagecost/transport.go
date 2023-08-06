package usagecost

import (
	"context"
	"strings"

	"github.com/go-kit/kit/endpoint"
	kitlogrus "github.com/go-kit/kit/log/logrus"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/sirupsen/logrus"

	mhttp "joi-energy-golang/http"
	"joi-energy-golang/http/middleware"
	"joi-energy-golang/http/serveroption"
	"net/http"
)

// MakeGetUsageCostHandler returns a handler for the Readings service.
func MakeGetUsageCostHandler(
	s Service,
	logger *logrus.Entry,
) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerBefore(serveroption.ExtractAcceptHeaderIntoContext),
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(kitlogrus.NewLogrusLogger(logger))),
		kithttp.ServerErrorEncoder(middleware.MakeEncodeErrorFunc(logger)),
	}

	mw := endpoint.Chain(
		middleware.MakeAcceptHeaderValidationMiddleware(),
	)

	endpointHandler := kithttp.NewServer(
		mw(makeGetUsageCostEndpoint(s)),
		decodeSmartMeterIdAndTimeFromRequest,
		mhttp.EncodeResponse,
		opts...,
	)

	return endpointHandler
}

func decodeSmartMeterIdAndTimeFromRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := strings.Split(r.URL.Path, "/")

	// URL must contains /usage-cost/calculate/<smart-meter-id>/<days>
	// TODO:: Need optimization here, get the time value as query params
	return map[string]string{
		"smartMeterId": params[3],
		"numberOfDays": params[4],
	}, nil
}
