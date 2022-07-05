package integrations

import "context"

type LinkToMoneyService interface {
	VerifyUser(ctx context.Context) (bool, error)
	SaveUser(ctx context.Context) (bool, error)
	RemoveLink(ctx context.Context) (bool, error)
}
