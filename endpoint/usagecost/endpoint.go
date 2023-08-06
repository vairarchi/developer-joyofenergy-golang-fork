package usagecost

import (
	"context"
	"joi-energy-golang/domain"

	"github.com/go-kit/kit/endpoint"
)

func makeGetUsageCostEndpoint(s Service) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.UsageCostRequst)

		err := validateRequestParams(req)
		if err != nil {
			return nil, err
		}

		// call BL func that will get
		usageCost, err := s.GetUsageCost(ctx, req.SmartMeterId, req.NumberOfDays)
		if err != nil {
			return nil, err
		}
		return usageCost, nil
	}
}
