package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
		
	
		// got := wallet.Balance()
	
		// fmt.Printf("address of balance in test is %v \n", &wallet.balance)
	
		// want := Bitcoin(10)
	
		// if got != want {
		// 	t.Errorf("got %s want %s", got, want)
		// }
	})
	
	t.Run("With", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
		

		// got := wallet.Balance()

		// want := Bitcoin(10)

		// if got != want {
		// 	t.Errorf("got %s want %s", got, want)
		// }
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(1000))

		assertBalance(t, wallet, startingBalance)
		assetError(t, err, ErrInsufficientFunds)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't one")
	}
}

func assetError(t testing.TB, got error,  want error) {
	t.Helper()

	if got == nil {
		// 발생시 테스트 중지
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}