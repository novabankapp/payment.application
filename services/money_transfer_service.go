package services

import (
	"context"
	common_error "github.com/novabankapp/common.data/error_handling"
	"github.com/novabankapp/payment.application/dtos"
	"github.com/novabankapp/payment.application/services/integrations"
	"github.com/novabankapp/payment.data/error_handling"
	"github.com/novabankapp/wallet.application/commands"
	"github.com/novabankapp/wallet.application/services"
	"github.com/shopspring/decimal"
)

type MoneyTransferService interface {
}

type moneyTransferService struct {
	walletService services.WalletService
	moneyService  integrations.MoneyService
}

func (m *moneyTransferService) ReceiveFromService(
	ctx context.Context,
	fromNovaSuspenseWalletAggregateId,
	toWalletAggregateId,
	description string,
	amount decimal.Decimal,
) (result *string, error error) {
	transId, err := m.moneyService.ReceiveFrom(ctx, toWalletAggregateId, amount)
	if err != nil {
		return nil, error_handling.PaymentServiceError{
			Err: common_error.WrapError(err, err.Error()),
		}
	}

	//debit from NovaSuspense
	err = m.walletService.Commands.DebitWalletCommand.Handle(ctx, commands.NewDebitWalletCommand(
		fromNovaSuspenseWalletAggregateId,
		toWalletAggregateId,
		amount,
		description,
	))
	//credit to Wallet
	err = m.walletService.Commands.CreditWalletCommand.Handle(ctx, commands.NewCreditWalletCommand(
		toWalletAggregateId,
		fromNovaSuspenseWalletAggregateId,
		amount,
		description,
	))

	return transId, nil
}
func (m *moneyTransferService) TransferToService(
	ctx context.Context,
	from dtos.AccountDto,
	to dtos.AccountDto,
	amount decimal.Decimal,
	toNovaSuspenseWalletAggregateId,
	description string,
) (result *string, error error) {

	transId, err := m.moneyService.TransferTo(ctx, to, amount)
	if err != nil {
		return nil, error_handling.PaymentServiceError{
			Err: common_error.WrapError(err, err.Error()),
		}

	}
	err = m.walletService.Commands.DebitWalletCommand.Handle(ctx, commands.NewDebitWalletCommand(
		from.Identifier,
		toNovaSuspenseWalletAggregateId,
		amount,
		description,
	))
	if err != nil {
		return nil, error_handling.PaymentServiceError{
			Err: common_error.WrapError(err, err.Error()),
		}
		//reverse debit
		m.moneyService.ReverseTransfer(ctx, *transId, amount)
	}
	return transId, err
}
func (m *moneyTransferService) WalletToWallet(
	ctx context.Context,
	creditWalletAggregateId string,
	amount decimal.Decimal,
	debitWalletAggregateId,
	description string) (result bool, error error) {
	err := m.walletService.Commands.CreditWalletCommand.Handle(ctx, commands.NewCreditWalletCommand(
		creditWalletAggregateId,
		debitWalletAggregateId,
		amount,
		description,
	))
	if err != nil {
		return false, error_handling.PaymentServiceError{
			Err: common_error.WrapError(err, err.Error()),
		}
	}
	err = m.walletService.Commands.DebitWalletCommand.Handle(ctx, commands.NewDebitWalletCommand(
		debitWalletAggregateId,
		creditWalletAggregateId,
		amount,
		description,
	))
	if err != nil {
		return false, error_handling.PaymentServiceError{
			Err: common_error.WrapError(err, err.Error()),
		}
		//reverse credit
	}
	return true, nil
}
