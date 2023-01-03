package cash

import (
	"fmt"
    "math"

	"github.com/TheFirstPenny/assets/pkg/currency"
)

const NEGATIVE_VALUE_MESSAGE = "The amount cannot be negative"
const DIFFERENT_CURRENCIES_MESSAGE = "The currencies are different"
const OVERFLOW_VALUE_MESSAGE = "The amount too big"

type Cash struct {
	amount   int64
	currency *currency.Currency
    name string
    description string
}

func NewCash(name string, description string, amount int64, currency *currency.Currency) (Cash, error) {
	if amount < 0 {
		return Cash{}, fmt.Errorf(NEGATIVE_VALUE_MESSAGE)
	}

    return Cash{name: name, description: description, amount: amount, currency: currency}, nil
}

func (m *Cash) Income(am *Cash) (Cash, error) {
	if !m.isCurrencyEqual(am) {
		return Cash{}, fmt.Errorf(DIFFERENT_CURRENCIES_MESSAGE)
	}
    if (math.MaxInt64 - am.amount) < m.amount {
        return Cash{}, fmt.Errorf(OVERFLOW_VALUE_MESSAGE)
    }
 
    return Cash{amount: m.amount + am.amount, currency: m.currency}, nil
}

func (m *Cash) Expense(am *Cash) (Cash, error) {
	if !m.isCurrencyEqual(am) {
		return Cash{}, fmt.Errorf(DIFFERENT_CURRENCIES_MESSAGE)
	}
	if m.isLessThan(am) {
		return Cash{}, fmt.Errorf(NEGATIVE_VALUE_MESSAGE)
	}

    return Cash{amount: m.amount - am.amount, currency: m.currency}, nil
}

func (m *Cash) isCurrencyEqual(am *Cash) bool {
    
	return m.currency.IsEqual(am.currency)
}

func (m *Cash) isLessThan(am *Cash) bool {

	return m.amount < am.amount
}
