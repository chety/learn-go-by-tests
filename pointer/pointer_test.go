package pointer

import (
	"testing"
)

func TestPointer(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(100))
		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(90))
	})

	t.Run("withdraw insufficient balance", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(50)}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInSufficientBalance)
		assertBalance(t, wallet, Bitcoin(50))
	})
}

func assertBalance(t testing.TB, w Wallet, expected Bitcoin) {
	t.Helper()
	got := w.Balance()
	if expected != got {
		t.Errorf("expected %s got %s", expected, got)
	}
}

func assertError(t testing.TB, actual, expected error) {
	t.Helper()
	if actual == nil {
		t.Fatal("expected error got nothing")
	}

	if actual != expected {
		t.Errorf("expected %s got %s", expected, actual)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("Got an error. Was not expecting one")
	}
}
