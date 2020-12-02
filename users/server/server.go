package server

import (
	"context"

	"github.com/ben-toogood/kite/common/validations"
	oz "github.com/go-ozzo/ozzo-validation/v4"
)

type Users struct{}

func grpcError(ctx context.Context, err error) error {
	ozerr, ok := err.(oz.Errors)
	if ok && len(ozerr) > 0 {
		return validations.NewError(ctx, err)
	}

	// handle database errors etc
	// add errors to tracing

	return err
}
