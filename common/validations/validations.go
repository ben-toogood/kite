package validations

import (
	"context"

	"github.com/cloudflare/cfssl/log"
	validations "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// NewError may be used to create a GRPC error with validation errors attached to it
func NewError(ctx context.Context, err error) error {
	st := status.New(codes.InvalidArgument, "Some of the fields you provided in your request are not valid")
	errors := validationErrToPBErrors(err)
	st, _ = st.WithDetails(&Errors{Errors: errors})
	return st.Err()
}

// NewErrorFromValidations create a custom error from you own validations
func NewErrorFromValidations(ctx context.Context, errors []*Error) error {
	st := status.New(codes.InvalidArgument, "Some of the fields you provided in your request are not valid")
	st, _ = st.WithDetails(&Errors{Errors: errors})
	return st.Err()
}

// ExtractErrors may be used to extract the validation errors from a GRPC error
func ExtractErrors(e error) []*Error {
	st := status.Convert(e)
	for _, detail := range st.Details() {
		switch v := detail.(type) {
		case *Errors:
			return v.Errors
		case error:
			return []*Error{
				&Error{
					Field:   "__unknown__",
					Message: v.Error(),
				},
			}
		default:
			log.Errorf("Unexpected error details %+v", v)
		}
	}
	return nil
}

// validationErrToPBErrors transforms errors into `Error`
func validationErrToPBErrors(err error) []*Error {
	verrs := err.(validations.Errors)
	var errs []*Error
	for field, err := range verrs {
		v := &Error{
			Field:   field,
			Message: err.Error(),
		}
		errs = append(errs, v)
	}
	return errs
}
