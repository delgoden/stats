package stats

import "github.com/delgoden/bank/pkg/bank/types"

// Avg calculates the average payment amount
func Avg(payments []types.Payment) types.Money {
	avgPayment := types.Money(0)
	for _, payment := range payments {
		avgPayment += payment.Amount
	}
	return avgPayment / types.Money(len(payments))
}

// TotalInCategory calculates the amount of purchases in a category
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	paymentsInCategory := types.Money(0)
	for _, payment := range payments {
		if payment.Category == category {
			paymentsInCategory += payment.Amount
		}
	}
	return paymentsInCategory
}