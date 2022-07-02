package integrations

import (
	"context"
	"github.com/novabankapp/payment.application/dtos"
	"github.com/shopspring/decimal"
)

type MoneyService interface {
	VerifyUser(ctx context.Context, userIdentifier string) (user *string, error error)
	TransferTo(ctx context.Context, toAccount dtos.AccountDto, amount decimal.Decimal) (transId *string, error error)
	TransferFrom(ctx context.Context, fromAccount dtos.AccountDto, amount decimal.Decimal) (transId *string, error error)
	ReverseTransfer(ctx context.Context, transactionId string, amount decimal.Decimal) (error error)
}
