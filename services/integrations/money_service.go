package integrations

import (
	"context"
	"github.com/novabankapp/payment.application/dtos"
	"github.com/shopspring/decimal"
)

type MoneyService interface {
	VerifyUser(ctx context.Context, userIdentifier, secretCode string) (user *string, error error)
	TransferTo(ctx context.Context, toAccount dtos.AccountDto, amount decimal.Decimal) (transId *string, error error)
	ReceiveFrom(ctx context.Context, toWalletId string, amount decimal.Decimal) (transId *string, error error)
	TransferFrom(ctx context.Context, fromAccount dtos.AccountDto, amount decimal.Decimal) (transId *string, error error)
	ReverseTransfer(ctx context.Context, transactionId string, amount decimal.Decimal) (error error)
}

type moneyService struct {
}

func (m moneyService) VerifyUser(ctx context.Context, userIdentifier, secretCode string) (user *string, error error) {
	//TODO implement me
	panic("implement me")
}

func (m moneyService) TransferTo(ctx context.Context, toAccount dtos.AccountDto, amount decimal.Decimal) (transId *string, error error) {
	//TODO implement me
	panic("implement me")
}

func (m moneyService) ReceiveFrom(ctx context.Context, toWalletId string, amount decimal.Decimal) (transId *string, error error) {
	//TODO implement me
	panic("implement me")
}

func (m moneyService) TransferFrom(ctx context.Context, fromAccount dtos.AccountDto, amount decimal.Decimal) (transId *string, error error) {
	//TODO implement me
	panic("implement me")
}

func (m moneyService) ReverseTransfer(ctx context.Context, transactionId string, amount decimal.Decimal) (error error) {
	//TODO implement me
	panic("implement me")
}

func NewMoneyService() MoneyService {
	return &moneyService{}
}
