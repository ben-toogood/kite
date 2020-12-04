package resolvers

import (
	"encoding/json"
	"errors"
)

type GraphQLUpload struct {
	Filename string `json:"filename"`
	MIMEType string `json:"mimetype"`
	Filepath string `json:"filepath"`
}

func (u GraphQLUpload) ImplementsGraphQLType(name string) bool {
	return name == "Upload"
}
func (u *GraphQLUpload) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case map[string]interface{}:
		data, err := json.Marshal(input)
		if err != nil {
			u = &GraphQLUpload{}
		} else {
			json.Unmarshal(data, u)
		}

		return nil
	default:
		return errors.New("Cannot unmarshal received type as a GraphQLUpload type")
	}
}
