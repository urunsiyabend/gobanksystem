// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package db

import (
	"context"
)

type Querier interface {
	AddAccountBalance(ctx context.Context, arg AddAccountBalanceParams) (Account, error)
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error)
	CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteAccount(ctx context.Context, id int64) error
	DeleteEntry(ctx context.Context, id int64) error
	DeleteTransfer(ctx context.Context, id int64) error
	GetAccount(ctx context.Context, id int64) (Account, error)
	GetAccountForUpdate(ctx context.Context, id int64) (Account, error)
	GetEntry(ctx context.Context, id int64) (Entry, error)
	GetTransfer(ctx context.Context, id int64) (Transfer, error)
	GetUser(ctx context.Context, username string) (User, error)
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error)
	ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error)
	ListEntriesByAccount(ctx context.Context, arg ListEntriesByAccountParams) ([]Entry, error)
	ListEntriesByAccountAndTime(ctx context.Context, arg ListEntriesByAccountAndTimeParams) ([]Entry, error)
	ListTransfers(ctx context.Context, arg ListTransfersParams) ([]Transfer, error)
	ListTransfersByAccount(ctx context.Context, arg ListTransfersByAccountParams) ([]Transfer, error)
	ListTransfersByAccountAndTime(ctx context.Context, arg ListTransfersByAccountAndTimeParams) ([]Transfer, error)
	ListTransfersByAccountAndToAccount(ctx context.Context, arg ListTransfersByAccountAndToAccountParams) ([]Transfer, error)
	ListTransfersByAccountAndToAccountAndTime(ctx context.Context, arg ListTransfersByAccountAndToAccountAndTimeParams) ([]Transfer, error)
	ListTransfersToAccount(ctx context.Context, arg ListTransfersToAccountParams) ([]Transfer, error)
	ListTransfersToAccountAndTime(ctx context.Context, arg ListTransfersToAccountAndTimeParams) ([]Transfer, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error)
	UpdateEntry(ctx context.Context, arg UpdateEntryParams) (Entry, error)
	UpdateTransfer(ctx context.Context, arg UpdateTransferParams) (Transfer, error)
}

var _ Querier = (*Queries)(nil)
