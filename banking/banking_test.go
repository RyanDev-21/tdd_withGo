package Banking

import (
	"testing"
)


func TestWallet(t *testing.T){

	showBalance := func(t testing.TB,wallet Wallet ,want Bitcoin){
		t.Helper()
		got := wallet.Balance()	
		if got != want{
			t.Errorf("got %s want %s",got,want)
		}
	}

	showError := func(t testing.TB,got error,want string){
		t.Helper()
		if got == nil{
			t.Fatal("wanted an error didn't get one")
		}
		if got.Error() != want{
			t.Errorf("got %q want %q",got,want)
		}
	}


	t.Run("deposit",func(t *testing.T) {

		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		expected := Bitcoin(10)
		
		showBalance(t,wallet,expected)
		

	})

	t.Run("withdraw",func(t *testing.T) {
		
		wallet := Wallet{Bitcoin(100)}
        err :=wallet.Withdraw(50)

		expected := Bitcoin(50)
		
		showBalance(t,wallet,expected)
		showError(t,err,"cannot withdraw, insufficient funds")

	})

	


	t.Run("testing with insufficient funds ", func(t *testing.T) {
		startingBalance := Bitcoin(20)	
		wallet := Wallet{startingBalance}
		err :=wallet.Withdraw(100)

		showBalance(t,wallet,startingBalance)
		showError(t,err,"cannot withdraw, insufficient funds")

	}	)

}
