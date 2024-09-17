package pointerserrors

import (
	"testing"
)

func assertBalance(t testing.TB, wallet Wallet, expected Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	want := Bitcoin(expected)

	if got != want {
		t.Errorf("Expected %s got %s", want, got)
	}
}

func assertError(t testing.TB, err error, expected error) {
	t.Helper()

	if err == nil {
		t.Error("wanted an error but didn't get one")
	}

	if err != expected {
		t.Errorf("got %q, want %q", err, expected)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Error("wanted an error but didn't get one")
	}
}

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, 10)

	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 20}

		err := wallet.Withdraw(Bitcoin(10))

		assertBalance(t, wallet, 10)
		assertNoError(t, err)

	})

	t.Run("Withdraw with insufficent funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}
