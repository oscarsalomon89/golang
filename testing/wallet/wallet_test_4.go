package wallet_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
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
				assert.NotNil(t, err, "wanted an error but didn't get one")
				assert.Equal(t, err, tt.wantedError, "they should be equal")
				return
			}

			assert.Nil(t, err)
			assert.Equal(t, tt.wallet.Balance, tt.want, "they should be equal")
		})
	}
}
