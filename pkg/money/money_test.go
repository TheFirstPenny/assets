package money

import (
	"testing"

	"github.com/TheFirstPenny/assets/pkg/currency"
)

func TestNewMoneySuccess(t *testing.T) {
	_, err := NewMoney(10, &currency.RussianRuble)
	if err != nil {
		t.Fatalf("NewMoney should not return err value")
	}
}

func TestNewMoneyNegativeAmount(t *testing.T) {
	var amount int64 = -1
	_, err := NewMoney(amount, &currency.RussianRuble)
	if err == nil {
		t.Fatalf(`NewMoney function: pass amount=%d, expect err != nil, get err=%v`, amount, err)
	}
}
