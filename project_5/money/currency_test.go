package money_test

import (
	"reflect"
	"testing"

	"github.com/sidpatel93/pocket-sized-go-projects/project_5/money"
)

func TestConvert(t *testing.T) {
	tc := map[string]struct {
		amount   money.Amount
		to       money.Currency
		validate func(t *testing.T, amount money.Amount, err error)
	}{
		"convert USD to EUR": {
			amount: money.Amount{},
			to:     money.Currency{},
			validate: func(t *testing.T, got money.Amount, err error) {
				if err != nil {
					t.Fatalf("expected no error, got %v", err)
				}
				expected := money.Amount{}
				if !reflect.DeepEqual(expected, got) {
					t.Fatalf("expected %v, got %v", expected, got)
				}
			},
		},
	}
	for name, tc := range tc {
		t.Run(name, func(t *testing.T) {
			got, err := money.Convert(tc.amount, tc.to)
			tc.validate(t, got, err)
		})
	}
}
