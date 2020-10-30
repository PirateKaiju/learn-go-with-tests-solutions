package errors

import "testing"

func TestWallet(t *testing.T) {

	t.Run("testing deposit", func(t *testing.T) {

		wallet := Wallet{}
		wallet.deposit(10)

		assertBalance(t, wallet, 10)

	})

	t.Run("testing withdraw", func(t *testing.T) {
		wallet := Wallet{20}

		err := wallet.withdraw(10)

		assertBalance(t, wallet, 10)
		assertNoError(t, err)

	})

	t.Run("testing withdraw with insufficient funds", func(t *testing.T) {
		wallet := Wallet{10}

		err := wallet.withdraw(20)

		assertBalance(t, wallet, 10)
		assertError(t, err, ErrInsufficientFunds)

	})

}

func assertBalance(t *testing.T, wallet Wallet, want int) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func assertNoError(t *testing.T, err error) {

	t.Helper()

	if err != nil {
		t.Errorf("error occurred")
	}

}

func assertError(t *testing.T, got error, want error) {

	t.Helper()

	if got == nil {
		t.Fatal("wanted error")
	}

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

}
