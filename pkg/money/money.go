package money

import (
	"fmt"

	"github.com/TheFirstPenny/assets/pkg/currency"
)

const NEGATIVE_VALUE_MESSAGE = "The amount cannot be negative"
const DIFFERENT_CURRENCIES_MESSAGE = "The currencies are different"

type Money struct {
	amount   int64
	currency *currency.Currency
}

func NewMoney(amount int64, currency *currency.Currency) (Money, error) {
	if amount < 0 {
		return Money{}, fmt.Errorf(NEGATIVE_VALUE_MESSAGE)
	}
	return Money{amount, currency}, nil
}

func (m *Money) add(am *Money) (Money, error) {
	if !m.isCurrencyEqual(am) {
		return Money{}, fmt.Errorf(DIFFERENT_CURRENCIES_MESSAGE)
	}
	return Money{m.amount + am.amount, m.currency}, nil
}

func (m *Money) subtract(am *Money) (Money, error) {
	if !m.isCurrencyEqual(am) {
		return Money{}, fmt.Errorf(DIFFERENT_CURRENCIES_MESSAGE)
	}
	if m.isLessThan(am) {
		return Money{}, fmt.Errorf(NEGATIVE_VALUE_MESSAGE)
	}
	return Money{m.amount - am.amount, m.currency}, nil
}

func (m *Money) isCurrencyEqual(am *Money) bool {
	return m.currency.IsEqual(am.currency)
}

func (m *Money) isLessThan(am *Money) bool {
	return m.amount < am.amount
}
