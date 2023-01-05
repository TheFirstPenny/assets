package money 

import (
	"math"
	"testing"

	"github.com/TheFirstPenny/assets/pkg/currency"
)

func TestNewMoneySuccess(t *testing.T) {
	_, err := NewMoney("Money", "Money description", 10, &currency.RussianRuble)
	if err != nil {
		t.Fatalf("NewMoney should not return err value")
	}
}

func TestNewMoneyNegativeAmount(t *testing.T) {
	var amount int64 = -1
	_, err := NewMoney("Money", "Money description", amount, &currency.RussianRuble)
    if (err == nil) {
        t.Fatalf(`NewMoney function: passed amount=%d, expected err != nil`, amount)
    } else if err.Error() != NEGATIVE_VALUE_MESSAGE {
		t.Fatalf(`NewMoney function: passed amount=%d, expected error message = %s`, amount, err.Error())
	}
}

func TestIncomeSuccess(t *testing.T) {
	var amountBase, amountToAdd, expectedAmount int64 = 10, 5, 15
	moneyBase, err := NewMoney("Money", "Money description", amountBase, &currency.RussianRuble)
	moneyToAdd, err := NewMoney("Money", "Money description", amountToAdd, &currency.RussianRuble)

	newMoney, err := moneyBase.Income(&moneyToAdd)

	if err != nil {
		t.Fatalf(`Money.Add method: passed base Money struct: %+v, Money struct to add: %+v, got unexpected error: %s`, moneyBase, moneyToAdd, err.Error())
	}
	if newMoney.amount != expectedAmount {
		t.Fatalf(`Money.Add method: passed base Money struct: %+v, Money struct to add: %+v, got amount=%d, expected %d`, moneyBase, moneyToAdd, newMoney.amount, expectedAmount)
	}
}

func TestIncomeWithDifferentCurrency(t *testing.T) {
    var amount int64 = 1
    moneyBase, _ := NewMoney("Money", "Money description", amount, &currency.RussianRuble)
    moneyToAdd, _ := NewMoney("Money", "Money description", amount, &currency.BelarusianRuble)

    _, err := moneyBase.Income(&moneyToAdd)

    if err == nil {
        t.Fatalf(`Money.Add() method: passed base Money struct: %+v, Money struct to add: %+v, expected error, got nil`, moneyBase, moneyToAdd)
    } else if err.Error() != DIFFERENT_CURRENCIES_MESSAGE {
        t.Fatalf(`Money.Add() method: passed base Money struct: %+v, Money struct to add: %+v, expected %s error message`, moneyBase, moneyToAdd, DIFFERENT_CURRENCIES_MESSAGE)
    }
}

func TestIncomeWithOverflow(t *testing.T) {
    var amount int64 = math.MaxInt64
    moneyBase, _ := NewMoney("Money", "Money description", amount, &currency.RussianRuble)
    moneyToAdd, _ := NewMoney("Money", "Money description", 1, &currency.RussianRuble)

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
   moneyBase, err := NewMoney("Money", "Money description", amountBase, &currency.RussianRuble)
   moneyToSubtract, err := NewMoney("Money", "Money description", amountToSubtract, &currency.RussianRuble)

   newMoney, err := moneyBase.Expense(&moneyToSubtract)

   if err != nil {
       t.Fatalf(`Expected no error, actual got error`)
   }
   if newMoney.amount != expectedAmount {
       t.Fatalf(`Expected value: %d, actual value: %d`, expectedAmount, newMoney.amount)
   }

   amountToSubtract = 15
   expectedAmount = 0
   moneyBase, _ = NewMoney("Money", "Money description", amountBase, &currency.RussianRuble)
   moneyToSubtract, err1 := NewMoney("Money", "Money description", amountToSubtract, &currency.RussianRuble)

   newMoneyZero, err1 := moneyBase.Expense(&moneyToSubtract)

   if err1 != nil {
       t.Fatalf(`Exptected no error, actual got errror`)
   }
   if newMoneyZero.amount != expectedAmount {
       t.Fatalf(`Expected value: %d, actual value: %d`, expectedAmount, newMoney.amount)
   }
}

func TestExpenseWithDifferentCurrency(t *testing.T) {
    var amountBase, amountToSubtract int64 = 2, 1
    moneyBase, _ := NewMoney("Money", "Money description", amountBase, &currency.RussianRuble)
    moneyToSubtract, _ := NewMoney("Money", "Money description", amountToSubtract, &currency.BelarusianRuble)

    _, err := moneyBase.Expense(&moneyToSubtract)

    if err == nil {
        t.Fatalf(`Expected: error message, actual: no error message`)
    } else if err.Error() != DIFFERENT_CURRENCIES_MESSAGE {
        t.Fatalf(`Expected error message: %s, actual error message: %s`, DIFFERENT_CURRENCIES_MESSAGE, err.Error())
    }
}

func TestExpenseWithNegativeResult(t *testing.T) {
    var amountBase, amountToSubtract int64 = 5, 10; 
    moneyBase, _ := NewMoney("Money", "Money description", amountBase, &currency.RussianRuble)
    moneyToSubtract, _ := NewMoney("Money", "Money description", amountToSubtract, &currency.RussianRuble)

    _, err1 := moneyBase.Expense(&moneyToSubtract)

    if err1 == nil {
        t.Fatalf(`Expected: error message, actual: no error message`)
    } else if err1.Error() != NEGATIVE_VALUE_MESSAGE {
        t.Fatalf(`Expected error message: %s, actual error message: %s`, NEGATIVE_VALUE_MESSAGE, err1.Error())
    }

}
