package stats

import (
	"fmt"
	"github.com/delgoden/bank/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID: 1,
			Amount: 10_000_00,
			Category: "auto",
		},
		{
			ID: 2,
			Amount: 15_000_00,
			Category: "auto",
		},
		{
			ID: 3,
			Amount: 13_000_00,
			Category: "auto",
		},
		{
			ID: 4,
			Amount: 16_000_00,
			Category: "auto",
		},
		{
			ID: 5,
			Amount: 15_000_00,
			Category: "auto",
		},
		{
			ID: 6,
			Amount: 14_000_00,
			Category: "auto",
		},
	}

	result := Avg(payments)
	fmt.Println(result)

	// Output: 1383333
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID: 1,
			Amount: 10_000_00,
			Category: "auto",
		},
		{
			ID: 2,
			Amount: 15_000_00,
			Category: "auto",
		},
		{
			ID: 3,
			Amount: 13_000_00,
			Category: "mobile",
		},
		{
			ID: 4,
			Amount: 16_000_00,
			Category: "mobile",
		},
		{
			ID: 5,
			Amount: 15_000_00,
			Category: "auto",
		},
		{
			ID: 6,
			Amount: 14_000_00,
			Category: "auto",
		},
	}
	result := TotalInCategory(payments, "auto")
	fmt.Println(result)

	// Output: 5400000
}