package cash

import (
	"math"
	"testing"

	"github.com/TheFirstPenny/assets/pkg/currency"
)

func TestNewCashSuccess(t *testing.T) {
	_, err := NewCash(10, &currency.RussianRuble)
	if err != nil {
		t.Fatalf("NewCash should not return err value")
	}
}

func TestNewCashNegativeAmount(t *testing.T) {
	var amount int64 = -1
	_, err := NewCash(amount, &currency.RussianRuble)
    if (err == nil) {
        t.Fatalf(`NewCash function: passed amount=%d, expected err != nil`, amount)
    } else if err.Error() != NEGATIVE_VALUE_MESSAGE {
		t.Fatalf(`NewCash function: passed amount=%d, expected error message = %s`, amount, err.Error())
	}
}

func TestIncomeSuccess(t *testing.T) {
	var amountBase, amountToAdd, expectedAmount int64 = 10, 5, 15
	moneyBase, err := NewCash(amountBase, &currency.RussianRuble)
	moneyToAdd, err := NewCash(amountToAdd, &currency.RussianRuble)

	newCash, err := moneyBase.Income(&moneyToAdd)

	if err != nil {
		t.Fatalf(`Cash.Add method: passed base Cash struct: %+v, Cash struct to add: %+v, got unexpected error: %s`, moneyBase, moneyToAdd, err.Error())
	}
	if newCash.amount != expectedAmount {
		t.Fatalf(`Cash.Add method: passed base Cash struct: %+v, Cash struct to add: %+v, got amount=%d, expected %d`, moneyBase, moneyToAdd, newCash.amount, expectedAmount)
	}
}

func TestIncomeWithDifferentCurrency(t *testing.T) {
    var amount int64 = 1
    moneyBase, _ := NewCash(amount, &currency.RussianRuble)
    moneyToAdd, _ := NewCash(amount, &currency.BelarusianRuble)

    _, err := moneyBase.Income(&moneyToAdd)

    if err == nil {
        t.Fatalf(`Cash.Add() method: passed base Cash struct: %+v, Cash struct to add: %+v, expected error, got nil`, moneyBase, moneyToAdd)
    } else if err.Error() != DIFFERENT_CURRENCIES_MESSAGE {
        t.Fatalf(`Cash.Add() method: passed base Cash struct: %+v, Cash struct to add: %+v, expected %s error message`, moneyBase, moneyToAdd, DIFFERENT_CURRENCIES_MESSAGE)
    }
}

func TestIncomeWithOverflow(t *testing.T) {
    var amount int64 = math.MaxInt64
    moneyBase, _ := NewCash(amount, &currency.RussianRuble)
    moneyToAdd, _ := NewCash(1, &currency.RussianRuble)

    _, err := moneyBase.Income(&moneyToAdd)

    if err == nil {
        t.Fatalf(`Expected: error message, actual: no error message`)
    } else if err.Error() != OVERFLOW_VALUE_MESSAGE {
        t.Fatalf(`Expected: %s, actual %s`, OVERFLOW_VALUE_MESSAGE, err.Error())
    }

    _, err1 := moneyToAdd.Income(&moneyBase)

    if err1 == nil {
        t.Fatalf(`Expected: error message, actual: no error message`)
    } else if err1.Error() != OVERFLOW_VALUE_MESSAGE {
        t.Fatalf(`Expected: %s, actual %s`, OVERFLOW_VALUE_MESSAGE, err1.Error())
    }
 }

func TestExpenseSuccess(t *testing.T) {
   var amountBase, amountToSubtract, expectedAmount int64 = 15, 10, 5
   moneyBase, err := NewCash(amountBase, &currency.RussianRuble)
   moneyToSubtract, err := NewCash(amountToSubtract, &currency.RussianRuble)

   newCash, err := moneyBase.Expense(&moneyToSubtract)

   if err != nil {
       t.Fatalf(`Expected no error, actual got error`)
   }
   if newCash.amount != expectedAmount {
       t.Fatalf(`Expected value: %d, actual value: %d`, expectedAmount, newCash.amount)
   }

   amountToSubtract = 15
   expectedAmount = 0
   moneyBase, _ = NewCash(amountBase, &currency.RussianRuble)
   moneyToSubtract, err1 := NewCash(amountToSubtract, &currency.RussianRuble)

   newCashZero, err1 := moneyBase.Expense(&moneyToSubtract)

   if err1 != nil {
       t.Fatalf(`Exptected no error, actual got errror`)
   }
   if newCashZero.amount != expectedAmount {
       t.Fatalf(`Expected value: %d, actual value: %d`, expectedAmount, newCash.amount)
   }
}

func TestExpenseWithDifferentCurrency(t *testing.T) {
    var amountBase, amountToSubtract int64 = 2, 1
    moneyBase, _ := NewCash(amountBase, &currency.RussianRuble)
    moneyToSubtract, _ := NewCash(amountToSubtract, &currency.BelarusianRuble)

    _, err := moneyBase.Expense(&moneyToSubtract)

    if err == nil {
        t.Fatalf(`Expected: error message, actual: no error message`)
    } else if err.Error() != DIFFERENT_CURRENCIES_MESSAGE {
        t.Fatalf(`Expected error message: %s, actual error message: %s`, DIFFERENT_CURRENCIES_MESSAGE, err.Error())
    }
}

func TestExpenseWithNegativeResult(t *testing.T) {
    var amountBase, amountToSubtract int64 = 5, 10; 
    moneyBase, _ := NewCash(amountBase, &currency.RussianRuble)
    moneyToSubtract, _ := NewCash(amountToSubtract, &currency.RussianRuble)

    _, err1 := moneyBase.Expense(&moneyToSubtract)

    if err1 == nil {
        t.Fatalf(`Expected: error message, actual: no error message`)
    } else if err1.Error() != NEGATIVE_VALUE_MESSAGE {
        t.Fatalf(`Expected error message: %s, actual error message: %s`, NEGATIVE_VALUE_MESSAGE, err1.Error())
    }

}
