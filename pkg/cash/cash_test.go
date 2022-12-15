package cash

import (
	"testing"
)

func TestMoneyCreation(t *testing.T) {
	money := NewMoney(1, &RussianRuble)
}
