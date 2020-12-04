package resolvers

import (
	"errors"
	"time"
)

type Timestamp struct {
	Time time.Time
}

func (u Timestamp) ImplementsGraphQLType(name string) bool {
	return name == "Timestamp"
}
func (u *Timestamp) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case int64:
		u.Time = time.Unix(input, 0)
		return nil
	case nil:
		return nil
	default:
		return errors.New("Cannot unmarshal received type as a Timestamp type")
	}
}
