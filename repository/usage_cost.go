package repository

import (
	"context"
	"joi-energy-golang/domain"
	"sort"
	"time"
)

// Define struct that will have meter reading field
type CostParams struct {
	Readings *MeterReadings
	Plans    *PricePlans
	Accounts *Accounts
}

func (res CostParams) GetUsageCost(ctx context.Context, meterID string, timePeriod int) (*domain.UsageCostResponse, error) {
	// TODO:: use model.UsageCost struct
	// BL here...
	pricePlanId := res.Accounts.PricePlanIdForSmartMeterId(meterID)
	plan := getPricePlan(res.Plans.pricePlans, pricePlanId)
	specificReadings := getReadingsOfNDays(res.Readings.GetReadings(meterID), timePeriod)
	calculateCost(specificReadings, plan)
	// return domain wala
	return &domain.UsageCostResponse{}, nil

}
func getPricePlan(allPlans []domain.PricePlan, planId string) domain.PricePlan {
	for _, plan := range allPlans {
		if plan.PlanName == planId {
			return plan
		}
	}
	return domain.PricePlan{}
}

func getReadingsOfNDays(allReadings []domain.ElectricityReading, timePeriod int) []domain.ElectricityReading {
	// Get the current date and time
	now := time.Now()

	// Calculate the date N days ago from the current date
	NdaysAgo := now.AddDate(0, 0, -timePeriod)

	// Filter the readings that fall within the last N days
	var lastNDaysReadings []domain.ElectricityReading
	for _, reading := range allReadings {
		if reading.Time.After(NdaysAgo) && reading.Time.Before(now) {
			lastNDaysReadings = append(lastNDaysReadings, reading)
		}
	}

	return lastNDaysReadings
}

func getTotalDuration(lastNDaysReadings []domain.ElectricityReading) time.Duration {
	if len(lastNDaysReadings) == 0 {
		return 0
	}

	// Sort the readings by time in ascending order
	sort.Slice(lastNDaysReadings, func(i, j int) bool {
		return lastNDaysReadings[i].Time.Before(lastNDaysReadings[j].Time)
	})

	// Find the first and last readings in the slice
	firstReading := lastNDaysReadings[0]
	lastReading := lastNDaysReadings[len(lastNDaysReadings)-1]

	// Calculate the total duration between the first and last readings
	totalDuration := lastReading.Time.Sub(firstReading.Time)

	return totalDuration
}
