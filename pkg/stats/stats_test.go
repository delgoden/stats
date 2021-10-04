package stats

import (
	"github.com/delgoden/bank/v2/pkg/types"
	"reflect"
	"testing"
)

func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount: 20000},
		{ID: 1, Category: "mobile", Amount: 1000},
		{ID: 1, Category: "auto", Amount: 2000},
		{ID: 1, Category: "food", Amount: 1000},
		{ID: 1, Category: "mobile", Amount: 1000},
		{ID: 1, Category: "fun", Amount: 10000},
		{ID: 1, Category: "auto", Amount: 30000},
		{ID: 1, Category: "food", Amount: 3000},
	}
	want := map[types.Category]types.Money{
		"auto":   17333,
		"mobile": 1000,
		"food":   2000,
		"fun":    10000,
	}
	
	got := CategoriesAvg(payments)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("invalid result,\n got:\n%v\n want:\n%v\n", got, want)
	}
}
