package dtos

import "github.com/shopspring/decimal"

type TransactionDto struct {
	FromAccount AccountDto      `json:"from_account"`
	ToAccount   AccountDto      `json:"to_account"`
	Amount      decimal.Decimal `json:"amount"`
	Description string          `json:"description"`
}
