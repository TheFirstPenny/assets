package money 

import (
	"fmt"
    "math"

	"github.com/TheFirstPenny/assets/pkg/currency"
)

const NEGATIVE_VALUE_MESSAGE = "The amount cannot be negative"
const DIFFERENT_CURRENCIES_MESSAGE = "The currencies are different"
const OVERFLOW_VALUE_MESSAGE = "The amount too big"

type Money struct {
	amount   int64
	currency *currency.Currency
    name string
    description string
}

func NewMoney(name string, description string, amount int64, currency *currency.Currency) (Money, error) {
	if amount < 0 {
		return Money{}, fmt.Errorf(NEGATIVE_VALUE_MESSAGE)
	}

    return Money{name: name, description: description, amount: amount, currency: currency}, nil
}

func (m *Money) Income(am *Money) (Money, error) {
	if !m.isCurrencyEqual(am) {
		return Money{}, fmt.Errorf(DIFFERENT_CURRENCIES_MESSAGE)
	}
    if (math.MaxInt64 - am.amount) < m.amount {
        return Money{}, fmt.Errorf(OVERFLOW_VALUE_MESSAGE)
    }
 
    return Money{amount: m.amount + am.amount, currency: m.currency}, nil
}

func (m *Money) Expense(am *Money) (Money, error) {
	if !m.isCurrencyEqual(am) {
		return Money{}, fmt.Errorf(DIFFERENT_CURRENCIES_MESSAGE)
	}
	if m.isLessThan(am) {
		return Money{}, fmt.Errorf(NEGATIVE_VALUE_MESSAGE)
	}

    return Money{amount: m.amount - am.amount, currency: m.currency}, nil
}

func (m *Money) isCurrencyEqual(am *Money) bool {
    
	return m.currency.IsEqual(am.currency)
}

func (m *Money) isLessThan(am *Money) bool {

	return m.amount < am.amount
}
