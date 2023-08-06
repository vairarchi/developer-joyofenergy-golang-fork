package usagecost

import (
	"context"

	"github.com/sirupsen/logrus"

	"joi-energy-golang/domain"
	"joi-energy-golang/repository"
)

type Service interface {
	GetUsageCost(ctx context.Context, smartMeterId string, timePeriod int) (*domain.UsageCostResponse, error)
}

type service struct {
	logger    *logrus.Entry
	usageCost *repository.CostParams
	accounts  *repository.Accounts
}

func NewService(
	logger *logrus.Entry,
	usageCost *repository.CostParams,
	accounts *repository.Accounts,
) Service {
	return &service{
		logger:    logger,
		usageCost: usageCost,
		accounts:  accounts,
	}
}

func (s *service) GetUsageCost(ctx context.Context, smartMeterId string, timePeriod int) (*domain.UsageCostResponse, error) {
	return s.usageCost.GetUsageCost(ctx, smartMeterId, timePeriod)
}
