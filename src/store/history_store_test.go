package store

import (
	"context"
	"testing"

	"github.com/sisu-network/pairswap-be/src/model"
	"github.com/stretchr/testify/require"
)

func TestHistoryStore(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	t.Run("create_history", func(t *testing.T) {
		t.Parallel()

		db, err := NewDBStores(dbConfig)
		require.NoError(t, err)
		history := &model.History{
			Address:     "0x24188b8f4431a982c3e192c609d77bdb398a8c3d",
			Recipient:   "addr_test1qq03mhme54m5sz7ywj7jhfw3vmexgngc0rd2v6uz58q64qu62stt3lxlavkqwa8gujx6k6yqreavfpks0cduqpnxrkasas9duc",
			TokenSymbol: "TIGER",
			SrcChain:    "binance-testnet",
			DestChain:   "fantom-testnet",
			Amount:      "1",
			SrcHash:     "0x9737ca634e8ca3d73a42104f7ff0b9255262c9bc9726beed2f40e7ffc0026aa1",
			SrcLink:     "https://testnet.bscscan.com/tx/0x9737ca634e8ca3d73a42104f7ff0b9255262c9bc9726beed2f40e7ffc0026aa1",
		}

		require.NoError(t, db.HistoryStore.Create(ctx, history))
	})
}
