package wallet_test

import (
	"errors"
	"testing"

	"github.com/teamcubation/neocamp-meli/testing/wallet"
)

func TestWalletWithdrawTDT(t *testing.T) {
	tests := []struct {
		name        string
		wallet      wallet.Wallet
		amount      int
		want        int
		wantedError error
	}{
		{
			name: "Balance OK: Withdraw",
			wallet: wallet.Wallet{
				Balance: 100,
			},
			amount:      25,
			want:        75,
			wantedError: nil,
		},
		{
			name: "Balance Error: Withdraw",
			wallet: wallet.Wallet{
				Balance: 20,
			},
			amount:      25,
			wantedError: errors.New("not enough money"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.wallet.Withdraw(tt.amount)
			if tt.wantedError != nil {
				if err == nil {
					t.Error("wanted an error but didn't get one")
				}

				if err.Error() != tt.wantedError.Error() {
					t.Errorf("result '%s' and expected '%s'", err, tt.wantedError)
				}
				return
			}

			if err != nil {
				t.Fatal("unexpected error")
			}

			result := tt.wallet.Balance

			if tt.want != result {
				t.Errorf("result %d and expected %d", result, tt.want)
			}
		})
	}
}
