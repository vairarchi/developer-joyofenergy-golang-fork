package usagecost

import (
	"joi-energy-golang/domain"

	validation "github.com/go-ozzo/ozzo-validation"
)

func validateRequestParams(request domain.UsageCostRequst) error {
	if err := validation.Validate(request.SmartMeterId, validation.Required); err != nil {
		return err
	}
	return validation.Validate(request.NumberOfDays, validation.Required)
}
