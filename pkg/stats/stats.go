package stats

import "github.com/delgoden/bank/v2/pkg/types"

// Avg calculates the average payment amount
func Avg(payments []types.Payment) types.Money {
	avgPayment := types.Money(0)
	countNotFail := types.Money(0)
	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			avgPayment += payment.Amount
			countNotFail ++
		}
	}
	return avgPayment / countNotFail
}

// TotalInCategory calculates the amount of purchases in a category
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	paymentsInCategory := types.Money(0)
	for _, payment := range payments {
		if payment.Category == category && payment.Status != types.StatusFail {
			paymentsInCategory += payment.Amount
		}
	}
	return paymentsInCategory
}