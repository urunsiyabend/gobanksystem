package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"github.com/urunsiyabend/simple_bank/util"
	"testing"
	"time"
)

func createRandomTransfer(t *testing.T, fromAccount, toAccount Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestQueries_CreateTransfer(t *testing.T) {
	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)

	createRandomTransfer(t, fromAccount, toAccount)
}

func TestQueries_GetTransfer(t *testing.T) {
	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)

	transfer := createRandomTransfer(t, fromAccount, toAccount)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer.ID, transfer2.ID)
	require.Equal(t, transfer.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer.Amount, transfer2.Amount)
	require.Equal(t, transfer.CreatedAt, transfer2.CreatedAt)
}

func TestQueries_UpdateTransfer(t *testing.T) {
	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)

	transfer := createRandomTransfer(t, fromAccount, toAccount)

	arg := UpdateTransferParams{
		ID:     transfer.ID,
		Amount: util.RandomMoney(),
	}

	transfer2, err := testQueries.UpdateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer.ID, transfer2.ID)
	require.Equal(t, transfer.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, arg.Amount, transfer2.Amount)
	require.Equal(t, transfer.CreatedAt, transfer2.CreatedAt)
}

func TestQueries_ListTransfers(t *testing.T) {
	for i := 0; i < 10; i++ {
		fromAccount := createRandomAccount(t)
		toAccount := createRandomAccount(t)

		createRandomTransfer(t, fromAccount, toAccount)
	}

	arg := ListTransfersParams{
		Limit:  5,
		Offset: 5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)

		require.NotEmpty(t, transfer.ID)
		require.NotEmpty(t, transfer.FromAccountID)
		require.NotEmpty(t, transfer.ToAccountID)
		require.NotEmpty(t, transfer.Amount)
		require.NotEmpty(t, transfer.CreatedAt)
	}
}

func TestQueries_DeleteTransfer(t *testing.T) {
	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)

	transfer := createRandomTransfer(t, fromAccount, toAccount)

	err := testQueries.DeleteTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, transfer2)
}

func TestQueries_ListTransfersByAccount(t *testing.T) {
	fromAccount := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		toAccount := createRandomAccount(t)

		createRandomTransfer(t, fromAccount, toAccount)
	}

	arg := ListTransfersByAccountParams{
		FromAccountID: fromAccount.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransfersByAccount(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)

		require.NotEmpty(t, transfer.ID)
		require.NotEmpty(t, transfer.FromAccountID)
		require.NotEmpty(t, transfer.ToAccountID)
		require.NotEmpty(t, transfer.Amount)
		require.NotEmpty(t, transfer.CreatedAt)

		require.Equal(t, fromAccount.ID, transfer.FromAccountID)
	}
}

func TestQueries_ListTransfersByAccountAndTime(t *testing.T) {
	fromAccount := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		toAccount := createRandomAccount(t)

		createRandomTransfer(t, fromAccount, toAccount)
	}

	arg := ListTransfersByAccountAndTimeParams{
		FromAccountID: fromAccount.ID,
		Limit:         5,
		Offset:        5,
		CreatedAt:     fromAccount.CreatedAt,
		CreatedAt_2:   time.Now(),
	}

	transfers, err := testQueries.ListTransfersByAccountAndTime(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)

		require.NotEmpty(t, transfer.ID)
		require.NotEmpty(t, transfer.FromAccountID)
		require.NotEmpty(t, transfer.ToAccountID)
		require.NotEmpty(t, transfer.Amount)
		require.NotEmpty(t, transfer.CreatedAt)

		require.Equal(t, fromAccount.ID, transfer.FromAccountID)
		require.True(t, transfer.CreatedAt.After(fromAccount.CreatedAt))
		require.True(t, transfer.CreatedAt.Before(time.Now()))
	}
}

func TestQueries_ListTransfersByAccountAndToAccount(t *testing.T) {
	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomTransfer(t, fromAccount, toAccount)
	}

	arg := ListTransfersByAccountAndToAccountParams{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransfersByAccountAndToAccount(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)

		require.NotEmpty(t, transfer.ID)
		require.NotEmpty(t, transfer.FromAccountID)
		require.NotEmpty(t, transfer.ToAccountID)
		require.NotEmpty(t, transfer.Amount)
		require.NotEmpty(t, transfer.CreatedAt)

		require.Equal(t, fromAccount.ID, transfer.FromAccountID)
		require.Equal(t, toAccount.ID, transfer.ToAccountID)
	}
}

func TestQueries_ListTransfersByAccountAndToAccountAndTime(t *testing.T) {
	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)

	for {
		if fromAccount.ID == toAccount.ID {
			toAccount = createRandomAccount(t)
		} else {
			break
		}
	}

	for i := 0; i < 10; i++ {
		createRandomTransfer(t, fromAccount, toAccount)
	}

	arg := ListTransfersByAccountAndToAccountAndTimeParams{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
		Limit:         5,
		Offset:        5,
		CreatedAt:     fromAccount.CreatedAt,
		CreatedAt_2:   time.Now(),
	}

	transfers, err := testQueries.ListTransfersByAccountAndToAccountAndTime(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)

		require.NotEmpty(t, transfer.ID)
		require.NotEmpty(t, transfer.FromAccountID)
		require.NotEmpty(t, transfer.ToAccountID)
		require.NotEmpty(t, transfer.Amount)
		require.NotEmpty(t, transfer.CreatedAt)

		require.Equal(t, fromAccount.ID, transfer.FromAccountID)
		require.Equal(t, toAccount.ID, transfer.ToAccountID)
		require.True(t, transfer.CreatedAt.After(fromAccount.CreatedAt))
		require.True(t, transfer.CreatedAt.Before(time.Now()))
	}
}
