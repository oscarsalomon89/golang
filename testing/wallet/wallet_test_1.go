package wallet_test

import (
	"testing"

	"github.com/teamcubation/neocamp-meli/testing/wallet"
)

func TestWalletDeposit(t *testing.T) {
	wallet := wallet.Wallet{
		Balance: 10,
	}

	wallet.Deposit(15)

	result := wallet.Balance
	expected := 25

	if expected != result {
		t.Errorf("result %d and expected %d", result, expected)
	}
}

func TestWalletWithdraw(t *testing.T) {
	wallet := wallet.Wallet{
		Balance: 100,
	}

	err := wallet.Withdraw(30)
	if err != nil {
		t.Errorf("error not expected %s", err.Error())
	}

	result := wallet.Balance
	expected := 70

	if expected != result {
		t.Errorf("result %d and expected %d", result, expected)
	}
}

func TestWalletWithdrawError(t *testing.T) {
	wallet := wallet.Wallet{
		Balance: 20,
	}

	err := wallet.Withdraw(30)
	if err == nil {
		t.Error("wanted an error but didn't get one")
	}

	expected := "not enough money"

	if err.Error() != expected {
		t.Errorf("result '%s' and expected '%s'", err, expected)
	}
}
