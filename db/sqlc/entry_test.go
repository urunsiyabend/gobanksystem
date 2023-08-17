package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"github.com/urunsiyabend/simple_bank/util"
	"testing"
	"time"
)

func createRandomEntry(t *testing.T, acc Account) Entry {
	arg := CreateEntryParams{
		AccountID: acc.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestQueries_CreateEntry(t *testing.T) {
	createRandomEntry(t, createRandomAccount(t))
}

func TestQueries_GetEntry(t *testing.T) {
	entry1 := createRandomEntry(t, createRandomAccount(t))

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.Equal(t, entry1.CreatedAt, entry2.CreatedAt)
}

func TestQueries_UpdateEntry(t *testing.T) {
	entry := createRandomEntry(t, createRandomAccount(t))

	arg := UpdateEntryParams{
		ID:     entry.ID,
		Amount: util.RandomMoney(),
	}

	entry2, err := testQueries.UpdateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry.ID, entry2.ID)
	require.Equal(t, entry.AccountID, entry2.AccountID)
	require.Equal(t, arg.Amount, entry2.Amount)
	require.Equal(t, entry.CreatedAt, entry2.CreatedAt)
}

func TestQueries_DeleteEntry(t *testing.T) {
	entry := createRandomEntry(t, createRandomAccount(t))

	err := testQueries.DeleteEntry(context.Background(), entry.ID)

	require.NoError(t, err)

	entry2, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.Error(t, err)
	require.Empty(t, entry2)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}

func TestQueries_ListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntry(t, createRandomAccount(t))
	}

	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}

func TestQueries_ListEntriesByAccount(t *testing.T) {
	acc := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomEntry(t, acc)
	}

	arg := ListEntriesByAccountParams{
		AccountID: acc.ID,
		Limit:     5,
		Offset:    5,
	}

	entries, err := testQueries.ListEntriesByAccount(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, entry.AccountID, acc.ID)
	}
}

func TestQueries_ListEntriesByAccountAndTime(t *testing.T) {
	acc := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomEntry(t, acc)
	}

	arg := ListEntriesByAccountAndTimeParams{
		AccountID:   acc.ID,
		Limit:       5,
		Offset:      5,
		CreatedAt:   time.Now().Add(-1 * time.Minute),
		CreatedAt_2: time.Now(),
	}

	entries, err := testQueries.ListEntriesByAccountAndTime(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, entry.AccountID, acc.ID)
		require.True(t, entry.CreatedAt.After(acc.CreatedAt))
		require.True(t, entry.CreatedAt.Before(time.Now()))
	}
}
